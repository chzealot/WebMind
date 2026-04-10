# WebMind

基于 [Claude Code](https://docs.anthropic.com/en/docs/claude-code) 的个人知识库，以精美的 HTML 页面作为知识载体。

**简体中文** | [English](README.md)

---

## WebMind 是什么？

WebMind 将你的问题转化为精心设计的、独立的 HTML 知识页面——并附带 PDF 和 PNG 导出。Claude Code 不会用纯文本回答，而是按照 [Mintlify](https://mintlify.com) 风格的设计系统生成精美的页面。

每条知识是一个目录，包含：

```
<主题目录>/
├── prompt.md          # 你的原始问题（经过润色）
├── assets/            # 图片和附件（如有）
├── <topic-slug>.html  # 知识正文页面
├── <目录名>.pdf       # 导出的 PDF
└── <目录名>.png       # 导出的 PNG（Retina 2x 长截图）
```

## 特性

- **Mintlify 风格设计** — 干净、专业的 HTML 页面，Inter + Geist Mono 字体，品牌绿点缀，药丸形组件
- **多语言支持** — 自动适配你的语言（中文、英文、日文等）
- **暗黑模式** — 每个页面内置亮色/暗色主题切换
- **响应式布局** — 同时适配桌面端和移动端
- **PDF & PNG 导出** — 基于 Go + Chrome DevTools Protocol 的一键导出
- **AI 原生工作流** — 专为 Claude Code 作为知识助手而设计

## 快速开始

### 前置要求

- 已安装 [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI
- Google Chrome（用于 PDF/PNG 导出）
- Go 1.22+（可选 — 仅从源码编译导出工具时需要；预编译二进制文件可在 [Releases](../../releases) 页面下载）

### 安装

1. 克隆仓库：
   ```bash
   git clone https://github.com/chzealot/WebMind.git
   cd WebMind
   ```

2. 获取导出工具（二选一）：

   **方式 A**：从源码编译（需要 Go 1.22+）：
   ```bash
   cd .claude/skills/webmind-export/scripts
   go build -o webmind-export .
   cd -
   ```

   **方式 B**：从 [Releases](../../releases) 下载预编译二进制文件：
   ```bash
   # macOS/Linux — 下载并解压对应平台的压缩包
   gh release download --repo chzealot/WebMind --pattern "webmind-export-*-$(uname -s | tr A-Z a-z)-$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/').*" --dir /tmp/wm
   tar xzf /tmp/wm/*.tar.gz -C /tmp/wm
   cp /tmp/wm/*/webmind-export .claude/skills/webmind-export/scripts/
   chmod +x .claude/skills/webmind-export/scripts/webmind-export
   rm -rf /tmp/wm
   ```

3. 在项目目录下启动 Claude Code：
   ```bash
   claude
   ```

4. 直接提问 — Claude Code 会按照 `CLAUDE.md` 中定义的工作流为你创建知识页面。

### 示例

```
你：什么是 MECE 原则？
```

Claude Code 会：
1. 创建目录（如 `MECE原则/`）
2. 编写 `prompt.md`，润色你的问题
3. 生成精美的 HTML 页面（`mece-principle.html`）
4. 导出 PDF 和 PNG 文件

## 工作原理

核心在于 `CLAUDE.md` — 它指导 Claude Code 遵循结构化的工作流：

1. **搜索** — 先检索已有知识（避免重复）
2. **编写 prompt.md** — 润色用户问题，使其对 LLM 更友好
3. **生成 HTML** — 遵循 `DESIGN.md` 中的设计规范
4. **导出** — PDF（矢量）+ PNG（Retina 2x 截图），使用 Go 导出工具
5. **整理** — 同级目录超过 10 个时自动按 MECE 原则归类

## 项目结构

```
WebMind/
├── CLAUDE.md                  # Claude Code 的工作流与规范
├── DESIGN.md                  # Mintlify 风格设计系统规范
├── LICENSE                    # MIT 开源协议
├── .claude/skills/            # Claude Code 技能
│   └── webmind-export/       # HTML → PDF/PNG 导出工具
└── <知识目录>/                # 你的知识条目
```

## 自定义

- **设计**：编辑 `DESIGN.md` 来修改配色、字体、间距或组件样式
- **工作流**：编辑 `CLAUDE.md` 来调整知识创建流程
- **导出**：修改 `.claude/skills/webmind-export/` 中的 Go 工具来调整导出参数

## Star History

<a href="https://www.star-history.com/?repos=ch+ze+a+lo+t%2Fch+ze+a+lo+t%2Cchzealot%2FWebMind&type=date&legend=top-left">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/chart?repos=ch ze a lo t/ch ze a lo t%2Cchzealot/WebMind&type=date&theme=dark&legend=top-left" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/chart?repos=ch ze a lo t/ch ze a lo t%2Cchzealot/WebMind&type=date&legend=top-left" />
   <img alt="Star History Chart" src="https://api.star-history.com/chart?repos=ch ze a lo t/ch ze a lo t%2Cchzealot/WebMind&type=date&legend=top-left" />
 </picture>
</a>

## 开源协议

[MIT](LICENSE)
