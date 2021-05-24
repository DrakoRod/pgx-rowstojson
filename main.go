package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/jackc/pgx/v4"
)

type Databases struct {
	DbName  string
	DbOwner string
}

// Convert pgx.Rows to json object in []byte format
func PgSqlRowsToJson(rows pgx.Rows) []byte {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		values, _ := rows.Values()
		for i, v := range values {
			valuePtrs[i] = reflect.New(reflect.TypeOf(v)).Interface() // allocate pointer to type
		}
		break
	}

	for rows.Next() {
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := reflect.ValueOf(valuePtrs[i]).Elem().Interface() // dereference pointer
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, _ := json.Marshal(tableData)

	return jsonData
}

func main() {
	urlExample := "postgres://drakorod:s3cr3tH4cks@localhost:5432/postgres"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, urlExample)

	if err != nil {
		fmt.Println("Unable to connect to database")
		os.Exit(1)
	}

	defer conn.Close(ctx)

	sql := "SELECT datname, u.usename AS datowner FROM pg_database db INNER JOIN pg_user u ON u.usesysid = db.datdba"
	rows, err := conn.Query(ctx, sql)

	json := PgSqlRowsToJson(rows)

	fmt.Println("JSON Result::> ", string(json))

}
