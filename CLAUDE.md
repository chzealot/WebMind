# CLAUDE.md

Working guidelines for Claude Code in this repository.

> **Core Principle**: **All** user questions should follow the workflow — create `prompt.md`, HTML, PDF, and PNG files — rather than answering directly.

---

## Project Overview

WebMind is a personal knowledge base that uses HTML pages as knowledge carriers, with a Mintlify-inspired design system. Each knowledge entry consists of a directory, a `prompt.md` file, an HTML file, and exported PDF and PNG files.

### Directory Structure

```
WebMind/
├── CLAUDE.md              # Workflow & guidelines (this file)
├── DESIGN.md              # Design system spec (Mintlify-inspired)
├── <knowledge-directory>/  # Named to describe the topic
│   ├── prompt.md          # The user's question / requirement
│   ├── assets/            # Attachments (images, files, etc. — created as needed)
│   ├── <topic-slug>.html  # Knowledge article (HTML)
│   ├── <dir-name>.pdf     # Exported PDF (filename matches directory name)
│   └── <dir-name>.png     # Exported PNG long-screenshot (filename matches directory name)
```

---

## Multilingual Behavior

WebMind automatically adapts to the user's language. **All generated content must match the language of the user's input**:

| User's language | Directory name | HTML filename | HTML body content | PDF/PNG filename |
|---|---|---|---|---|
| Chinese | 中文主题名 | english-slug.html | 中文 | 中文主题名.pdf/.png |
| English | English Topic Name | english-slug.html | English | English Topic Name.pdf/.png |
| Japanese | 日本語トピック名 | english-slug.html | 日本語 | 日本語トピック名.pdf/.png |
| Other | In that language | english-slug.html | In that language | In that language.pdf/.png |

**Rules:**
- **Directory names**: Use the user's language to describe the topic.
- **HTML filenames**: Always lowercase English words joined by hyphens (`-`), 2–5 words (e.g., `css-grid-layout.html`). This ensures cross-platform path compatibility.
- **PDF / PNG filenames**: Match the directory name exactly (in the user's language).
- **HTML body content**: Written entirely in the user's language.
- **Attachment filenames**: Always lowercase English + hyphens, semantically named (e.g., `floor-plan-overview.jpg`).

---

## Workflow

### Handling a New Question

1. **Search existing knowledge**: Look through existing directories for a matching entry.
2. **If a match exists**: Propose updating it — **you must get the user's approval** before modifying.
3. **If no match**: Create a new directory and files.
4. **Write `prompt.md` first**: Record the user's question/requirement (do NOT copy content from `DESIGN.md` into `prompt.md`).
5. **Generate HTML**: Combine `prompt.md`, `CLAUDE.md`, and `DESIGN.md` to produce the HTML file.
6. **Export PDF and PNG**: After any HTML creation or change, export is mandatory (see "Export Guidelines").
7. **Check directory structure**: If there are more than 10 sibling directories at the same level, reorganize using the MECE principle (see "Directory Management").

### Writing `prompt.md`

Do not simply paste the user's raw question. Polish it:

- Make it LLM-friendly for future content generation
- Improve structure and clarity
- Fix typos or grammatical issues
- Fill in obvious gaps — but do not add unrelated content without asking the user first

### Handling Attachments

When the user's question involves images, files, or other attachments:

1. **Save location**: Place attachments in `<knowledge-directory>/assets/` (sibling to `prompt.md`).
2. **Rename files**: If the original filename is meaningless (e.g., `IMG_001.jpg`, `screenshot_2026.png`), rename it to something semantic based on context (e.g., `hangzhou-floor-plan.jpg`). Use lowercase English + hyphens.
3. **Reference in `prompt.md`**:
   - Images: `![description](assets/filename.jpg)` — ensure Markdown preview renders the image.
   - Other files: `[file description](assets/filename.ext)`.

### Naming Conventions

- **Directory names**: In the user's language, describing the knowledge topic.
- **HTML filenames**: Lowercase English words + hyphens (`-`), 2–5 words, e.g., `css-grid-layout.html`.
- **PDF / PNG filenames**: Match the directory name (in the user's language), e.g., `<dir-name>.pdf`, `<dir-name>.png`.
- **Attachment filenames**: Lowercase English + hyphens, semantically named.
- **If name and content diverge**: Update the directory name and HTML filename; consider whether the directory hierarchy needs adjustment.

### Directory Management

- Knowledge directories start flat at the repository root.
- No more than 10 subdirectories at any level.
- When exceeded, reorganize using the **MECE principle** (Mutually Exclusive, Collectively Exhaustive):
  - **Mutually Exclusive**: Each subdirectory has clear topical boundaries; entries belong to exactly one category.
  - **Collectively Exhaustive**: All existing entries fit into some category — nothing is left out.
  - Aim for 3–10 entries per subdirectory.
  - After reorganization, update any relative path references in affected files.

---

## HTML Generation Guidelines

Generated HTML must follow the design spec in `DESIGN.md`. Key points:

- **Fonts**: Inter (body) + Geist Mono (code / technical labels)
- **Colors**: White background `#ffffff`, body text `#0d0d0d`, brand green `#18E299`
- **Border radius**: Buttons/inputs `9999px` (pill), cards `16px`, large cards `24px`
- **Borders**: `rgba(0,0,0,0.05)` at 5% opacity
- **Font weights**: Only 400 / 500 / 600
- **Spacing**: 8px base unit, section gaps 48px–96px
- **Each HTML must be a complete standalone page** — viewable by opening directly in a browser
- **Static HTML by default**: No animations (CSS animation, transition, JS animation) to ensure exported PDF/PNG content is fully readable. Only add animations when the user explicitly requests them.
- **Include `@media print` pagination-friendly styles**: Prevent figures, tables, and code blocks from being split across pages during PDF export (see `DESIGN.md` for details).

### Responsive Layout

- **Must support both desktop and mobile**, using CSS `@media` queries for responsive layout.
- Desktop: max content width per spec. Mobile (`max-width: 768px`): adjust font sizes, spacing, card layouts for readability.
- Images, tables, and code blocks must adapt to small screens — no horizontal overflow.
- Card grids should switch to single-column on mobile.

### Dark Mode

- **Dark mode is off by default** — pages load in light theme.
- **Provide a toggle button in the top-right corner**: Fixed position (`position: fixed`), using sun/moon icons (pure CSS or inline SVG — no external icon libraries). Clicking toggles light/dark theme.
- Implement by toggling `data-theme="dark"` on `<html>`, using CSS variables for all colors.
- **Dark theme palette**: Background `#0d0d0d`, body text `#e5e5e5`, card background `#1a1a1a`, borders `rgba(255,255,255,0.1)`. Brand green `#18E299` stays unchanged.
- Persist user's theme preference via `localStorage`.
- **PDF/PNG export always uses light theme** — force light colors in `@media print` styles.

---

## Export Guidelines

After any HTML creation or change, PDF and PNG export is mandatory.

### File Naming

- PDF and PNG filenames match the directory name (in the user's language): `<dir-name>.pdf`, `<dir-name>.png`.
- Export files are placed in the same directory as the HTML.

### Export Method

Use the Go tool provided by the `webmind-export` skill. A single command exports both PDF and PNG in one Chrome session:

```bash
<webmind-export-skill-dir>/scripts/webmind-export \
  -html ./<knowledge-dir>/<topic-slug>.html \
  -pdf ./<knowledge-dir>/<dir-name>.pdf \
  -png ./<knowledge-dir>/<dir-name>.png
```

The tool is built on Go + chromedp (Chrome DevTools Protocol) — no Node.js dependency. Export parameters:

- **PDF**: Vector output, preserves background colors, no headers/footers, supports CSS `@media print` pagination rules.
- **PNG**: Full-page screenshot, viewport width 1440px, `deviceScaleFactor: 2`, output 2880px physical pixel width, waits for `document.fonts.ready` to ensure web fonts are loaded.

### Notes

- Ensure the HTML file is saved and complete before exporting.
- If the HTML uses web fonts, network access must be available during export.
