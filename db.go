// db.go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "kv{~1l5K$0zO6rVrR3seE?zxD"
	DB_NAME     = "DomoffonDB"
)

func SelectDB(input string) (*sql.Rows, error) {
	dbinfo := fmt.Sprintf("host=localhost port=49044 user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	//fmt.Println("# Querying")
	rows, err := db.Query(input)
	checkErr(err)
	fmt.Println("3")

	return rows, err
}

func UpdateDB(input string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	//fmt.Println("# Inserting values")
	_, err = db.Query(input)
	fmt.Println(input)

	return err
}

func DeleteFromDB(input string) error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname, Created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)

	//fmt.Println("# Deleting")
	stmt, err := db.Prepare(input)
	checkErr(err)

	res, err := stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	return err
}
