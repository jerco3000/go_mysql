package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	)

func main() {

	db, err := sql.Open("mysql", "root:kfsdppf50uopcx@/jerco?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM agenda")
	checkErr(err)

	for rows.Next() {
		var id int
		var nombre string
		var apellido string
		err = rows.Scan(&id, &nombre, &apellido)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(nombre)
		fmt.Println(apellido)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}