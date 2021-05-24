# DrakoHacks
## pgx-rowstojson

Is a simple example to convert a PostgreSQL query rows to json using [pgxv4](https://github.com/jackc/pgx) driver. The main function is based in this [stackoverflow question](https://stackoverflow.com/questions/50238439/how-to-convert-pgx-rows-from-query-to-json-array).

### Execute the example

Just clone project and change database connection settings like this: 

```
$ git clone 

```

Change database connections settings in code

```
urlExample := "postgres://postgres:s3cr3t@localhost:5432/postgres"

```
Execute main.go to test

```
$ go run main.go
JSON Result::>  [{"datname":"template1","datowner":"postgres"},{"datname":"template0","datowner":"postgres"},{"datname":"explain","datowner":"drakorod"}{"datname":"drakorod","datowner":"drakorod"},{"datname":"go_project","datowner":"corvus_user"}]
```

#### Enjoy! :)
