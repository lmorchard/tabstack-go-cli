# MCP server smoke-test results

**Date:** 2026-04-28
**Branch:** feat/embedded-mcp-server
**Transport:** http (127.0.0.1:7777)
**API key:** redacted (temporary staging key, rotate after)

Session initialized OK; `tools/list` returned all five `tabstack_*` tools.

## Per-tool results

```
==== tabstack_extract_markdown ====
{
  "content": "---\ntitle: Example Domain\nurl: https://example.com\n---\n\nThis domain is for use in documentation examples without needing permission. Avoid use in operations.",
  "url": "https://example.com",
  "metadata": {
    "author": "",
    "created_at": "",
    "creator": "",
    "description": "",
    "image": "",
    "keywords": null,
    "modified_at": "",
    "page_count": 0,
    "pdf_version": "",
    "producer": "",
    "publisher": "",
    "site_name": "",
    "subject": "",
    "title": "",
    "type": "",
    "url": ""
  }
}

==== tabstack_extract_json ====
{
  "title": "Example Domain"
}

==== tabstack_generate_json ====
{
  "summary": "This page is an example domain used for documentation, instructing users to avoid its use in operations and providing a link to learn more."
}

==== tabstack_research ====
Paris is the capital and largest city of France [1][2][3]. It has served as the seat of France's national government since **508 AD**, when Clovis the Frank made it his capital, and it later regained this status under King Philippe Auguste (1180-1223) [1][3]. While Paris has largely remained the capital, there was a **four-year interruption** during World War II, from **1940 to 1944**, when Vichy functioned as the capital [3]. As of January 2026, Paris has an estimated city population of **2.04 million** within an area of **105.4 km<sup>2</sup>**, and a metropolitan population of **13.2 million** [1]. It holds the distinction of being the largest metropolitan area and fourth-most populous city in the European Union [1], and is home to the official residences of the French President and Pri
... [truncated, full length: 816 chars]

==== tabstack_automate ====
The page title is: "Example Domain"

```
