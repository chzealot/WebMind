---
name: webmind-export
description: "Export WebMind HTML knowledge pages to PDF and PNG. Use this skill whenever an HTML file in the WebMind project is created or changed and needs exporting. Trigger words: export, screenshot, PDF generation, PNG long-screenshot, HTML to PDF/PNG. Even if the user simply says 'export this' or 'generate PDF and PNG', use this skill. Note: export is mandatory after every HTML creation or modification — it is a required step in the WebMind workflow."
---

# WebMind HTML Export

Export WebMind knowledge page HTML files to PDF (vector) and PNG long-screenshot (Retina 2x).

## Tool Setup

The binary path is:

```
<skill-dir>/scripts/webmind-export
```

If the binary doesn't exist, obtain it using **one of the following methods** (in order of preference):

### Option 1: Build from source (requires Go 1.22+)

```bash
cd <skill-dir>/scripts && go build -o webmind-export . && cd -
```

### Option 2: Download pre-built binary from GitHub Releases

If the local machine has no Go build environment, download the latest pre-built binary:

```bash
# Detect OS and architecture, then download the matching archive
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "${ARCH}" in x86_64) ARCH="amd64" ;; aarch64|arm64) ARCH="arm64" ;; esac

# Download and extract the latest release
gh release download --repo chzealot/WebMind --pattern "webmind-export-*-${OS}-${ARCH}*" --dir /tmp/webmind-dl
cd /tmp/webmind-dl
tar xzf *.tar.gz 2>/dev/null || unzip *.zip 2>/dev/null
cp */webmind-export* <skill-dir>/scripts/webmind-export
chmod +x <skill-dir>/scripts/webmind-export
rm -rf /tmp/webmind-dl
cd -
```

> **Note**: Replace `<github-owner>` with the actual GitHub username or organization. On Windows, download the `.zip` archive manually from the GitHub Releases page.

## Usage

Export both PDF and PNG in a single command (one Chrome session, efficient):

```bash
<skill-dir>/scripts/webmind-export \
  -html ./<knowledge-dir>/<topic-slug>.html \
  -pdf ./<knowledge-dir>/<dir-name>.pdf \
  -png ./<knowledge-dir>/<dir-name>.png
```

You can also export just one format:

```bash
# PDF only
<skill-dir>/scripts/webmind-export -html ./example.html -pdf ./out.pdf

# PNG only
<skill-dir>/scripts/webmind-export -html ./example.html -png ./out.png
```

### Parameters

| Flag | Default | Description |
|------|---------|-------------|
| `-html` | (required) | Input HTML file path |
| `-pdf` | | Output PDF path |
| `-png` | | Output PNG path |
| `-width` | 1440 | Viewport width (CSS pixels) |
| `-scale` | 2.0 | PNG device scale factor (2.0 = Retina 2x, output width 2880px) |

## Naming Convention

- **PDF**: `<dir-name>/<dir-name>.pdf` — filename matches its parent directory name
- **PNG**: `<dir-name>/<dir-name>.png` — filename matches its parent directory name
- Export files are placed alongside the HTML file in the same directory

## Workflow

After any HTML creation or change, export is mandatory:

1. Ensure the HTML file is saved and complete
2. Run the export command with both `-pdf` and `-png` flags
3. Verify output: PDF should be a multi-page vector document; PNG should be a 2880px-wide Retina long-screenshot

## Technical Details

- **PDF**: Generated via Chrome DevTools Protocol's PrintToPDF, preserves background colors (`printBackground: true`), no headers/footers, supports CSS `@media print` pagination rules
- **PNG**: Full-page screenshot, `deviceScaleFactor: 2` (2880px physical pixels), waits for `document.fonts.ready` to ensure web fonts are loaded
- **Dependencies**: Locally installed Google Chrome. Go 1.22+ is needed only if building from source; otherwise use pre-built binaries from GitHub Releases.
- Uses the `chromedp` library (Go Chrome DevTools Protocol client) — no Node.js or Puppeteer needed
