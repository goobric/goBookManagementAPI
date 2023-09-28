package routes

import (
	"github.com/goobric/goBookManagementAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

// import (
// 	"github.com/gorilla/mux"
// 	"github.com/goobric/goBookManagementAPI"
// )

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}