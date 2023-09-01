package routes
import (
	"github.com/gorilla/mux"
	"github.com/kushagra-gupta01/bookStore/pkg/controllers"
)

var RegisterBokStoreRoutes = func(Router* mux.Router){
	Router.HandleFunc("/book/",controllers.getBook).Methods("GET")
	Router.HandleFunc("/book/",controllers.createBook).Methods("POST")
	Router.HandleFunc("/book/{bookId}",controllers.getBookById).Methods("GET")
	Router.HandleFunc("/book/{bookId}",controllers.updateBook).Methods("PUT")
	Router.HandleFunc("/book/{bookId}",controllers.deleteBook).Methods("DELETE")
}