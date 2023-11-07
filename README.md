# Goose App
This is a simple app that add migration tool to your project. It is based on [goose](https://github.com/pressly/goose) and [cobra](https://github.com/spf13/cobra) libraries.

## Usage
```bash
goose-app [command]
```

## Available Commands
* `help` - Help about any command
* `migrate up` - Migrate up to the most recent migration
* `migrate down` - Roll back a single migration from the current version
* `migrate redo` - Re-run the latest migration
* `migrate reset` - Roll back all migrations
* `migrate version` - Print the current version of the database