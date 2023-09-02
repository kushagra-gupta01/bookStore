package routes
import (
	"github.com/gorilla/mux"
	"github.com/kushagra-gupta01/bookStore/pkg/controllers"
)

var RegisterBokStoreRoutes = func(Router* mux.Router){
	Router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	Router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	Router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	Router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	Router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}