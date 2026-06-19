# Copilot instructions for Gator

## Build, test, and lint commands

- Build (whole project):
  - go build ./...
- Run (local dev / one-off):
  - go run . -- <command> <args>
    - Example: go run . register alice
- Tests (whole repo):
  - go test ./...
- Run a single package's tests:
  - go test ./internal/database
- Run a single test by name (in package):
  - go test ./internal/database -run TestMyFunction
  - or: go test ./... -run TestMyFunction
- Lint / vet (builtin):
  - go vet ./...
- SQL codegen (sqlc):
  - sqlc generate --config sqlc.yaml
  - (Requires sqlc installed; sqlc.yaml points sql/schema and sql/queries -> internal/database)

Notes: go.mod specifies `go 1.26.3`. Use the Go toolchain matching that version.

## High-level architecture

- Entry point: main.go
  - Reads config (internal/config/config.go), opens a PostgreSQL connection (lib/pq), instantiates generated DB queries, and runs CLI commands.
- CLI command pattern:
  - commands.go defines a `commands` registry and `state` (db + cfg).
  - Commands are registered by name and invoked from os.Args. Handlers live in handler_*.go.
- Database layer:
  - SQL is in `sql/schema/` (migrations/schema) and `sql/queries/` (query files).
  - sqlc (config: sqlc.yaml) generates typed Go DB code into `internal/database`.
  - DB usage is via the generated `Queries` type and an abstract DBTX interface in internal/database.
- Config:
  - internal/config/config.go reads/writes a JSON config at ~/.gatorconfig.json
  - Expected JSON keys: `db_url` (Postgres DSN) and `current_user_name`.
- RSS handling and aggregation logic:
  - rss_feed.go contains feed fetching/parsing.
  - Aggregation and feed/user commands implemented in handler_agg.go, handler_feed.go, handler_user.go.

## Key conventions / repository-specific patterns

- Command handlers: signature pattern is `func(*state, command) error`. Register handlers via `cmds.register("name", handler)`.
- DB-first workflow: add or change queries in `sql/queries/*.sql` and regenerate `internal/database` with `sqlc generate`.
- Schema changes: add numbered SQL files under `sql/schema/` (project uses simple numbered schema files).
- Config lives in the user's home dir as `.gatorconfig.json` (created/updated via config API).
- Postgres is the supported DB backend (github.com/lib/pq used in go.mod).
- Use the generated `internal/database` package for all DB access — avoid hand-rolled SQL in code.

## Useful developer commands

- Regenerate DB code (after changing SQL): sqlc generate --config sqlc.yaml
- Start a quick command (example addfeed): go run . addfeed https://example.com/feed
- Create config (example):
  - echo '{"db_url":"postgres://user:pass@localhost:5432/gator?sslmode=disable","current_user_name":"alice"}' > ~/.gatorconfig.json

## AI assistant notes for Copilot sessions

- Focus on: handlers (handler_*.go), sql/queries/*.sql, and internal/database (generated) when tracing data flow.
- When modifying DB access, update SQL files and run `sqlc generate` rather than editing generated files.
- Config changes require updating ~/.gatorconfig.json; main.go reads DB URL from that file.

