# fizzy-cli

A fast, human-friendly CLI for the Fizzy kanban board. Manage boards, cards, comments, tags, users, columns, and notifications from your terminal using Fizzy’s HTTP API.

## Features
- Token or magic-link authentication
- List, create, update, and delete boards/cards/comments/columns/users
- Bulk-friendly output with `--json` and `--plain`
- Config + env precedence for repeatable workflows
- Works on macOS and Linux (single Go binary)

## Install
Build from source:

```bash
go build ./...
```

If your default Go cache path is restricted:

```bash
GOCACHE=/path/to/.gocache go build ./...
```

The binary is `./cmd/fizzy-cli/fizzy-cli` when built from that folder, or use `go build -o fizzy-cli ./cmd/fizzy-cli`.

## Quick Start
1) Authenticate

```bash
fizzy-cli auth login --token $FIZZY_TOKEN
# or
fizzy-cli auth login --email user@example.com
```

2) Set your default account

```bash
fizzy-cli account set 897362094
# or persist base URL + account
fizzy-cli config set --base-url https://app.fizzy.do --account 897362094
```

3) List boards

```bash
fizzy-cli board list
```

## Usage Examples

List cards on a board:

```bash
fizzy-cli card list --board-id 03f5v9zkft4hj9qq0lsn9ohcm
```

Create a card:

```bash
fizzy-cli card create --board-id 03f5v9zkft4hj9qq0lsn9ohcm \
  --title "Add dark mode" \
  --description "Switch theme"
```

Upload a card image:

```bash
fizzy-cli card create --board-id 03f5v9zkft4hj9qq0lsn9ohcm \
  --title "Add hero image" \
  --image ./screenshot.png
```

Update a card:

```bash
fizzy-cli card update 4 --title "Add dark mode (updated)" --tag-id 03f5v9zo9qlcwwpyc0ascnilz
```

Comment on a card:

```bash
fizzy-cli comment create 4 --body "Looks good to me"
```

List notifications (unread only):

```bash
fizzy-cli notification list --unread
```

Machine output:

```bash
fizzy-cli card list --board-id 03f5v9zkft4hj9qq0lsn9ohcm --json
fizzy-cli board list --plain
```

## Configuration
Config file location (default):
- `~/.config/fizzy/config.json`

Precedence (highest to lowest):
1. Flags
2. Environment variables
3. Config file
4. Built-in defaults

Supported env vars:
- `FIZZY_BASE_URL`
- `FIZZY_TOKEN`
- `FIZZY_ACCOUNT`
- `FIZZY_CONFIG`

Inspect config:

```bash
fizzy-cli config show
```

## Output Modes
- Default: human-friendly tables
- `--plain`: line-oriented output (stable for scripts)
- `--json`: JSON output of raw API responses

## Security Notes
- Tokens and session cookies grant access to your account; keep them secret.
- `fizzy-cli config show` never prints secrets, only whether they are set.

## Command Reference
Run `fizzy-cli --help` or `fizzy-cli help <command>`.

Common commands:
- `auth login|logout|status`
- `account list|set`
- `config show|set`
- `board list|get|create|update|delete`
- `card list|get|create|update|delete|close|reopen|not-now|triage|untriage|tag|assign|watch|unwatch`
- `comment list|get|create|update|delete`
- `tag list`
- `column list|get|create|update|delete`
- `user list|get|update|deactivate`
- `notification list|read|unread|read-all`
