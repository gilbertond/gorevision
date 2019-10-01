Docker image
`gilbertond/mssql-server-linux`

execute
`go build -o goose *.go`
`goose mssql "sqlserver://sa:myPassw0rd@localhost:1433?database=master" status`  to check status
`goose mssql "sqlserver://sa:myPassw0rd@localhost:1433?database=master" up`  to migrate
`goose mssql "sqlserver://sa:myPassw0rd@localhost:1433?database=master" up-to 00001` rollback to revision

can use SQL or can use Go file  like in example below
`00001*.sql` or 

can create migration files like so
`goose create fetch_user_data go`  using go
`goose create add_some_column sql` using sql


```
Drivers:
    postgres
    mysql
    sqlite3
    mssql
    redshift

Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up

    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
    goose tidb "user:password@/dbname?parseTime=true" status
    goose mssql "sqlserver://user:password@dbname:1433?database=master" status

Options:

```