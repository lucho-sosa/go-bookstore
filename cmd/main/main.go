package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/lucho-sosa/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
