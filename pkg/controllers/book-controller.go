package controllers
import(
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/kushagra-gupta01/bookStore/pkg/models"
	"github.com/kushagra-gupta01/bookStore/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
)
var NewBook models.Book
func getBook(w http.ResponseWriter,r *http.Request){
	NewBook:=models.GetAllBooks()
	res,_ := json.Marshal(NewBook)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:= vars["BookId"]
	Id,err :=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:=models.GetBookById(Id)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func createBook(w http.ResponseWriter,r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.createBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func deleteBook(w http.ResponseWriter,r *http.Request){
	Vars:=mux.Vars(r)
	bookId:=Vars["bookId"]
	id,err:=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("parse error")
	}
	book:=models.DeleteBook(id)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookId:=vars["bookId"]
	id,err:= strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("parse error")
	}
	bookDetails,db:=models.GetBookById(id)
	if updateBook.Name !=""{
		bookDetails.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		bookDetails.Author=updateBook.Author
	}
	if updateBook.Publication!=""{
		bookDetails.Publication=updateBook.Publication
	}
	db.Save(&bookDetails)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}