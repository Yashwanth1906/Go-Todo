package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Yashwanth1906/Go-Todo/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.Routes(r)
	http.Handle("/", r)
	fmt.Println("Listening on Port 8000....")
	log.Fatal(http.ListenAndServe(":8000", r))
}
