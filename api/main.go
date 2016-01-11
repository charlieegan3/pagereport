package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/charlieegan3/pagereport/api/service"
)

func respondWithError(err error, w http.ResponseWriter) {
	data := struct {
		Message string `json:"message"`
	}{err.Error()}
	jsonString, _ := json.Marshal(data)
	io.WriteString(w, string(jsonString))
}

func parseQuery(rawUrl string) (query service.Query, err error) {
	queryUrl, err := url.Parse(rawUrl)
	if err != nil {
		return service.Query{}, err
	}

	params, err := url.ParseQuery(queryUrl.RawQuery)
	if err != nil {
		return service.Query{}, err
	}

	requestedUrl, exists := params["url"]
	if !exists {
		return service.Query{}, errors.New("Required param missing: url")
	}
	selector, exists := params["selector"]
	if !exists {
		return service.Query{}, errors.New("Required param missing: selector")
	}

	return service.Query{Url: requestedUrl[0], Selector: selector[0]}, err
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.String())
	if err != nil {
		respondWithError(err, w)
		return
	}
	response, err := service.ProcessQuery(query)
	if err != nil {
		respondWithError(err, w)
		return
	}

	jsonString, _ := json.Marshal(response)
	io.WriteString(w, string(jsonString))
}

func main() {
	fmt.Println("Starting...")
	if len(os.Args) == 1 {
		log.Fatal("Missing PORT parameter")
	}
	http.HandleFunc("/service", serviceHandler)
	http.ListenAndServe(":"+os.Args[1], nil)
}
