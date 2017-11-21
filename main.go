package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {

	fmt.Println("Connecting to database")
	connection()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8025", router))
}

