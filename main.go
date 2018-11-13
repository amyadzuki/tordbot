package main

import (
	"fmt"
	"io/ioutil"
)

func main(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage:", args[0], "/path/to/token.dat /path/to/database.db")
	}
	b, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", args[2])
	if err != nil {
		panic(err)
	}
}
