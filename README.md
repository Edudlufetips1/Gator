Gator
=====
Introduction
-----------

Gator is a multi-user command-line RSS feed aggregator. Follow your favorite blogs and news sources, and let Gator fetch new posts in the background so you can browse them all in one place.

Quick start
-----------

Requirements
- Go (>= 1.26.3)
- PostgreSQL

Install the CLI
- Install the binary (recommended):

  go install github.com/Edudlufetips1/Gator@latest

  After installing, the binary will be placed in $GOBIN (or $GOPATH/bin). Add that directory to your PATH if needed.

- Or run directly during development:

  go run . <command> [args]

Configuration
- Create a JSON config at ~/.gatorconfig.json with at least these keys:

  {
    "db_url": "postgres://user:pass@localhost:5432/gator?sslmode=disable",
    "current_user_name": "alice"
  }

  Example (shell):

  echo '{"db_url":"postgres://user:pass@localhost:5432/gator?sslmode=disable","current_user_name":"alice"}' > ~/.gatorconfig.json

Running the program
- With installed binary (example):

  Gator register alice
  Gator addfeed https://example.com/feed
  Gator listfeeds
  Gator aggregate (time interval [seconds]) --> Gator agg 10s

- With go run (development):

  go run . register alice
  go run . addfeed https://example.com/feed

Useful commands
- go build ./...        # build the whole project
- go test ./...         # run tests
- go vet ./...          # vet/lint checks

Notes
- The project expects a running Postgres instance reachable with the db_url in the config file.
- For schema or DB code changes, see sql/ and regenerate DB code with `sqlc generate --config sqlc.yaml` if applicable.

License
- See the repository for license information.
