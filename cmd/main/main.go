package main

import (
	"log"
	"net/http"

	"github.com/goobric/goBookManagementAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Create Server
//tell Golang where the Routers are

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r) //links to routes
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
