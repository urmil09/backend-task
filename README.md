# Flink: Backend interview task

This project contains the Flink backend interview task where we have a Go application with a simplified
version of an HTTP server to process HTTP requests.

## The task


## Development

### Requirements
- [Docker](https://docs.docker.com/compose/install/)
- [pressly/goose](https://github.com/pressly/goose#install)

### Run
To run the application, simply type `docker compose up` in the root folder.

Migrations will be applied and the HTTP server will bind to port `8080` together with a Postgres database which can be accessed on port `5432`.

To check if everything is going fine:
```bash
curl -v http://localhost:8080/health
```
You should get as response:
```
ok!
```

### Database migrations
To manage the database migrations, we are using [pressly/goose](https://github.com/pressly/goose).

You do not need to install the tool locally as we provide the option to run it via docker via the provided makefile commands:

- Check migration status: `make migration_status`
- Create a new migration: `make migration name=my_migration_name_here`
- Execute migrations: `make migrate`
- Rollback the last migration: `make migrate_down`

If you do want to run it locally the relevant credentials and settings can be found in `.env`
