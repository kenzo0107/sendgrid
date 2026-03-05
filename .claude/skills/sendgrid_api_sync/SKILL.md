name: sendgrid_api_sync
description: >
  Fetch and sync SendGrid API references, detect API changes.

## Fetching API Documentation

Twilio Docs serves Markdown via `<link rel="alternate" type="text/markdown">`. Append `.md` to any URL to get clean Markdown instead of the SPA HTML (600KB+).

- Base URL: `https://www.twilio.com/docs/sendgrid/api-reference/`
- Markdown URL: append `.md` to the page URL
  - Example: `https://www.twilio.com/docs/sendgrid/api-reference/alerts/create-a-new-alert.md`

## Local Documentation

- Scraper: `cmd/scraper/main.go`
- Output: `docs/` directory (`.md` files)
- Filename convention: replace `/` after `/api-reference/` with `_`, use `.md` extension
  - Example: `.../alerts/create-a-new-alert` → `alerts_create-a-new-alert.md`

### Bulk Download

```bash
go run ./cmd/scraper/
```

- Default: 5 concurrent downloads (`-c` flag to change)
- Output dir: `-out` flag (default: `docs/`)

### Single Endpoint Lookup

Use WebFetch with `.md` appended:

```
https://www.twilio.com/docs/sendgrid/api-reference/{category}/{endpoint}.md
```

## Adding New Endpoints

1. Add the URL to `urls` slice in `cmd/scraper/main.go`
2. Run `go run ./cmd/scraper/`
3. Refer to the generated `.md` file in `docs/` for implementation

## Change Detection Rules

- Fetch only endpoints related to the feature being modified (no full re-read)
- Detect:
    1. New endpoints
    2. Changed request parameters
    3. Deprecated fields
- Output only:
    - New feature additions
    - Required code modifications
    - Breaking changes
- Do not rewrite unchanged code
- Keep explanations under 5 lines
