
# Grapton

Grapton is a command-line interface (CLI) tool built in Go for managing user accounts, RSS feeds, and feed aggregation. Like Gator, it allows users to register, log in, add and follow RSS feeds, browse posts from followed feeds, and perform feed aggregation tasks. The tool interacts with a PostgreSQL database for storing users, feeds, feed follows, and posts. It uses sqlc for generating type-safe database queries and goose for database migrations.

## Features

   1. User management: Register, login, list users, reset.
    2. Feed management: Add feeds, list feeds, follow/unfollow feeds, list followed feeds.
    3. Feed aggregation: Aggregate content from RSS feeds.
    4. Post browsing: Browse posts from followed feeds.

## Requirements

To run Grapton, you'll need the following installed on your system:
 Go: Version 1.21 or later (the project uses Go modules).
 PostgreSQL: A running PostgreSQL instance (version 12 or later recommended). You'll need to create a database and have connection details ready.

Additionally, for development or building:
**sqlc** for generating database code (install via go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest).
**goose** for running database migrations (install via go install github.com/pressly/goose/v3/cmd/goose@latest).

## Installation

Install the Grapton CLI using Go's install command. This will compile and install the binary to your $GOPATH/bin directory (ensure it's in your PATH).
text
``` 
go install github.com/natnael-alemayehu/grapton@latest
```
The binary will be named grapton (based on the repository name), but you can rename it to grapton if preferred for convenience.

## Setup
1. Set Up the PostgreSQL Database
   - Create a new PostgreSQL database (e.g., grapton_db).
   - Update your database connection details (host, port, user, password, database name). You'll need these for the config file.

2. Run Database Migrations

The project uses goose for migrations, with schema files located in sql/schema/. Navigate to the project root and run:
text

```bash
goose -dir sql/schema postgres "postgres://user:password@host:port/dbname?sslmode=disable" up
```

Replace the connection string with your PostgreSQL details. This will apply migrations for tables like users, feeds, feed_follows, and posts.

If your Makefile includes a migration target (e.g., make migrate), you can use that instead.
3. Generate Database Code (Optional for Development)

If you've modified SQL queries or schema, regenerate the database code using sqlc:
text
```bash
sqlc generate
```

This reads from sqlc.yaml and outputs Go code to internal/database/.

4. Set Up the Config File

Grapton reads configuration from a file form the home "~" of your directory with a file name called **"~/.gatorconfig.json"**. Create a config file in the project root or your working directory with the following structure:
JSON

{
  "DBURL": "postgres://user:password@host:port/dbname?sslmode=disable"
}

    DBURL: The full PostgreSQL connection string. Replace with your actual details.

The config is loaded via config.Read(), which parses this file. If the file is missing or malformed, the program will log an error and exit.
5. Run the Program

Once installed and configured, run the CLI with commands:
text

```bash
grapton <command> [args...]
```

(If you didn't rename the binary, use grapton instead of grapton.)

Examples of commands you can run (after setup):

    register: Create a new user account.
    text

grapton register <username> <password>

login: Log in as a user (required for protected commands).
text

grapton login <username> <password>

addfeed: Add a new RSS feed (requires login).
text

grapton addfeed <feed_name> <feed_url>

follow: Follow an existing feed (requires login).
text

grapton follow <feed_url>

feeds: List all available feeds.
text

grapton feeds

agg: Run the feed aggregator to fetch and update posts from feeds.
text

grapton agg

browse: Browse posts from your followed feeds (requires login).
text

```bash
grapton browse
```

For a full list of commands, refer to commands.go or run the tool without arguments to see usage.
Project Structure

    main.go: Entry point, sets up config, database, and command handlers.
    commands.go: Defines the command registry.
    handler_*.go: Command handlers for users, feeds, posts, etc.
    internal/config/: Configuration loading.
    internal/database/: sqlc-generated database queries and models.
    sql/: SQL schema migrations (schema/) and queries (queries/).
    sqlc.yaml: Configuration for sqlc code generation.
    Makefile: Build and utility scripts (e.g., for generating code or migrations).

Development

    Clone the repository: git clone https://github.com/natnael-alemayehu/grapton.git
    Build locally: go build -o grapton
    Test: Run tests in internal/config/ and other packages with go test ./...
