package main

import (
	"net/http"
	"os"

	"github.com/azer/crud"
	"github.com/charlieegan3/pagereport/web/controllers/index"
	_ "github.com/lib/pq"
)

var DB *crud.DB

func init() {
	DB, _ := crud.Connect("postgres", os.Getenv("DATABASE_URL"))
	DB.Ping()
}

func main() {
	http.HandleFunc("/", index.View)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
