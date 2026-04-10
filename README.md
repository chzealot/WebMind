# WebMind

A personal knowledge base powered by [Claude Code](https://docs.anthropic.com/en/docs/claude-code), using beautifully designed HTML pages as knowledge carriers.

[简体中文](README.zh-CN.md) | **English**

---

## What is WebMind?

WebMind turns your questions into polished, self-contained HTML knowledge pages — complete with PDF and PNG exports. Instead of answering in plain text, Claude Code generates beautifully designed pages following a [Mintlify](https://mintlify.com)-inspired design system.

Each knowledge entry is a directory containing:

```
<topic-directory>/
├── prompt.md          # Your original question (polished)
├── assets/            # Images & attachments (if any)
├── <topic-slug>.html  # The knowledge page
├── <dir-name>.pdf     # Exported PDF
└── <dir-name>.png     # Exported PNG (Retina 2x long-screenshot)
```

## Features

- **Mintlify-inspired design** — Clean, professional HTML pages with Inter + Geist Mono typography, brand green accents, and pill-shaped components
- **Multilingual** — Automatically adapts to your language (Chinese, English, Japanese, etc.)
- **Dark mode** — Built-in light/dark theme toggle on every page
- **Responsive** — Desktop and mobile friendly
- **PDF & PNG export** — One-command export via a Go tool using Chrome DevTools Protocol
- **AI-native workflow** — Designed to work with Claude Code as your knowledge assistant

## Getting Started

### Prerequisites

- [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI installed
- Google Chrome (for PDF/PNG export)
- Go 1.22+ (optional — only needed to build the export tool from source; pre-built binaries are available on the [Releases](../../releases) page)

### Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/chzealot/WebMind.git
   cd WebMind
   ```

2. Get the export tool (choose one):

   **Option A**: Build from source (requires Go 1.22+):
   ```bash
   cd .claude/skills/webmind-export/scripts
   go build -o webmind-export .
   cd -
   ```

   **Option B**: Download pre-built binary from [Releases](../../releases):
   ```bash
   # macOS/Linux — download and extract the archive for your platform
   gh release download --repo chzealot/WebMind --pattern "webmind-export-*-$(uname -s | tr A-Z a-z)-$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/').*" --dir /tmp/wm
   tar xzf /tmp/wm/*.tar.gz -C /tmp/wm
   cp /tmp/wm/*/webmind-export .claude/skills/webmind-export/scripts/
   chmod +x .claude/skills/webmind-export/scripts/webmind-export
   rm -rf /tmp/wm
   ```

3. Start Claude Code in the project directory:
   ```bash
   claude
   ```

4. Ask any question — Claude Code will create a knowledge page for you following the workflow defined in `CLAUDE.md`.

### Example

```
You: What is the MECE principle?
```

Claude Code will:
1. Create a directory (e.g., `MECE Principle/`)
2. Write `prompt.md` with your polished question
3. Generate a beautiful HTML page (`mece-principle.html`)
4. Export PDF and PNG files

## How It Works

The magic is in `CLAUDE.md` — it instructs Claude Code to follow a structured workflow:

1. **Search** existing knowledge first (avoid duplicates)
2. **Write `prompt.md`** — polish the user's question for LLM-friendliness
3. **Generate HTML** — following the design spec in `DESIGN.md`
4. **Export** — PDF (vector) + PNG (Retina 2x screenshot) via the Go export tool
5. **Organize** — auto-reorganize directories when they exceed 10 entries (MECE principle)

## Project Structure

```
WebMind/
├── CLAUDE.md                  # Workflow & guidelines for Claude Code
├── DESIGN.md                  # Mintlify-inspired design system spec
├── LICENSE                    # MIT License
├── .claude/skills/            # Claude Code skills
│   └── webmind-export/       # HTML → PDF/PNG export tool
└── <knowledge-directories>/   # Your knowledge entries
```

## Customization

- **Design**: Edit `DESIGN.md` to change colors, typography, spacing, or component styles
- **Workflow**: Edit `CLAUDE.md` to adjust the knowledge creation process
- **Export**: Modify the Go tool in `.claude/skills/webmind-export/` for different export settings

## Star History

<a href="https://www.star-history.com/?repos=chzealot%2FWebMind&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/chart?repos=chzealot/WebMind&type=date&theme=dark&legend=top-left" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/chart?repos=chzealot/WebMind&type=date&legend=top-left" />
   <img alt="Star History Chart" src="https://api.star-history.com/chart?repos=chzealot/WebMind&type=date&legend=top-left" />
 </picture>
</a>

## License

[MIT](LICENSE)
