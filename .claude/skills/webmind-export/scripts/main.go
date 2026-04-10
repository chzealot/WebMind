package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func main() {
	htmlFile := flag.String("html", "", "HTML file path")
	pdfOut := flag.String("pdf", "", "Output PDF path")
	pngOut := flag.String("png", "", "Output PNG path")
	width := flag.Int("width", 1440, "Viewport width in CSS pixels")
	scale := flag.Float64("scale", 2.0, "Device scale factor for PNG")
	flag.Parse()

	if *htmlFile == "" {
		fmt.Fprintln(os.Stderr, "Usage: webmind-export -html <file> [-pdf out.pdf] [-png out.png]")
		os.Exit(1)
	}
	if *pdfOut == "" && *pngOut == "" {
		fmt.Fprintln(os.Stderr, "Error: at least one of -pdf or -png is required")
		os.Exit(1)
	}

	absHTML, err := filepath.Abs(*htmlFile)
	if err != nil {
		log.Fatalf("resolve HTML path: %v", err)
	}
	if _, err := os.Stat(absHTML); os.IsNotExist(err) {
		log.Fatalf("HTML file not found: %s", absHTML)
	}

	fileURL := "file://" + absHTML

	// Setup Chrome allocator
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-gpu", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Navigate and wait for page ready
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(int64(*width), 900),
		chromedp.Navigate(fileURL),
		chromedp.WaitReady("body"),
	); err != nil {
		log.Fatalf("navigate: %v", err)
	}

	// Wait for web fonts to finish loading
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		_, exceptionDetails, err := runtime.Evaluate(`document.fonts.ready.then(() => "loaded")`).
			WithAwaitPromise(true).
			WithTimeout(runtime.TimeDelta(10000)).
			Do(ctx)
		if err != nil {
			return fmt.Errorf("wait fonts: %w", err)
		}
		if exceptionDetails != nil {
			return fmt.Errorf("wait fonts exception: %s", exceptionDetails.Text)
		}
		return nil
	})); err != nil {
		// Font waiting failed, proceed anyway (local files may not have web fonts)
		fmt.Fprintf(os.Stderr, "Warning: font wait: %v (proceeding anyway)\n", err)
	}

	// Export PDF
	if *pdfOut != "" {
		absPDF, _ := filepath.Abs(*pdfOut)
		var buf []byte
		if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithDisplayHeaderFooter(false).
				Do(ctx)
			return err
		})); err != nil {
			log.Fatalf("PDF export: %v", err)
		}
		if err := os.WriteFile(absPDF, buf, 0644); err != nil {
			log.Fatalf("write PDF: %v", err)
		}
		fmt.Printf("PDF: %s\n", absPDF)
	}

	// Export PNG (full page, retina)
	if *pngOut != "" {
		absPNG, _ := filepath.Abs(*pngOut)
		var buf []byte
		if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
			// Get full page content size
			_, _, _, _, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return fmt.Errorf("get layout metrics: %w", err)
			}

			w := math.Ceil(contentSize.Width)
			h := math.Ceil(contentSize.Height)

			// Override device metrics for retina rendering
			if err := emulation.SetDeviceMetricsOverride(int64(w), int64(h), *scale, false).Do(ctx); err != nil {
				return fmt.Errorf("set device metrics: %w", err)
			}

			// Capture full page screenshot
			buf, err = page.CaptureScreenshot().
				WithFormat(page.CaptureScreenshotFormatPng).
				WithCaptureBeyondViewport(true).
				WithClip(&page.Viewport{
					X:      0,
					Y:      0,
					Width:  w,
					Height: h,
					Scale:  1,
				}).
				Do(ctx)
			if err != nil {
				return fmt.Errorf("capture screenshot: %w", err)
			}

			return nil
		})); err != nil {
			log.Fatalf("PNG export: %v", err)
		}
		if err := os.WriteFile(absPNG, buf, 0644); err != nil {
			log.Fatalf("write PNG: %v", err)
		}
		fmt.Printf("PNG: %s (%dx retina, %dpx viewport)\n", absPNG, int(*scale), *width)
	}
}
