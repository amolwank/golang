package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
* Tag a very simple struct
 */
type Tag struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	//open up our database connection.

	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/MYSQLTEST")
	//db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")

	//if there is error in opening the connections handle it

	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	//fmt.Println(db)

	defer db.Close()

	//execute the query
	results, err := db.Query("select id, name, population from cities")

	if err != nil {
		panic(err.Error()) // proper error handling intead of panic in your code
	}

	for results.Next() {
		var tag Tag

		err = results.Scan(&tag.ID, &tag.Name, &tag.Population)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(tag.ID, tag.Name, tag.Population)
	}
}
