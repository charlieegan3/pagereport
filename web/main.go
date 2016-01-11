package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting...")
	if len(os.Args) == 1 {
		log.Fatal("Missing PORT parameter")
	}
	panic(http.ListenAndServe(":"+os.Args[1], http.FileServer(http.Dir("static"))))
}
