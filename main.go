package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("Connecting to database")
	connection()
	router :=  cors.Default().Handler(NewRouter())
	log.Fatal(http.ListenAndServe(":8025", router))
}

