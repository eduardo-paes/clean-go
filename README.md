# Clean Arch Go Project

Clean Architecture project of a RESTful API using Go language for study purposes.

## Database and Libraries

**Database:**

- PostgresSQL

**Libraries:**

- [Pgx](https://github.com/jackc/pgx): Database connection
- [Mux](https://github.com/gorilla/mux): Request router and a dispatcher to combine incoming requests with their respective handlers.
- [Go-paginate](https://github.com/booscaaa/go-paginate): Creation of queries for Postgres
- [Viper](https://github.com/spf13/viper): Settings for the dev/prod environment
- [Testify](https://github.com/stretchr/testify): Test
- [Pgx Mock](https://github.com/pashagolub/pgxmock): Mock for the pgx connection pool
- [Migrate](https://github.com/golang-migrate/migrate): Run updates to our database

## Commands

Initialize Go Modules:

```bash
go mod init github.com/eduardo-paes/clean-go
```
