package main
import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
	"github.com/kushagra-gupta01/bookStore/pkg/routes"
)

func main(){
	r:=mux.NewRouter()
	routes.RegisterBokStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8000",r))
}