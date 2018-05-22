package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func loadDatabase(user, password, schema string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		user,
		password,
		schema,
	))
	return db, err
}
