package main

import (
	"fmt"
	"os"

	"github.com/azer/crud"
	_ "github.com/lib/pq"
)

var DB *crud.DB

func main() {
	DB, err := crud.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}
