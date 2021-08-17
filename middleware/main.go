package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	//"github.com/gorilla/handlers"
)

//const JsonContentType = "application/json"

// struct book
type Book struct {
	ID 		string `json:"id"`
	Isbn 	string `json:"isbn"`
	Title 	string `json:"title"`
	Author 	*Author `json:"author"`

}
type Author struct{
	Name 	string `json:"name"`
	Age 	int `json:age`
}
var books []Book
func main()  {
	router := mux.NewRouter().StrictSlash(true)
	books = append(books,Book{ID: "1",Isbn: "123",Title: "The First Book",Author: &Author{Name:"John",Age: 25}})
	books = append(books,Book{ID: "2",Isbn: "124",Title: "The Second Book",Author: &Author{Name:"Đạt",Age: 23}})
	// var age int = 23
	router.Methods(http.MethodGet).Path("/api/books/{id}").HandlerFunc(getBook)
	router.Methods(http.MethodGet).Path("/api/books").HandlerFunc(getBooks)
	router.Methods(http.MethodPost).Path("/api/books/add").HandlerFunc(createBook)
	router.Methods(http.MethodPut).Path("/api/books/update/{id}").HandlerFunc(updateBook)
	router.Methods(http.MethodDelete).Path("/api/books/delete/{id}").HandlerFunc(deleteBook)

	// router.HandlerFunc("/api/books/{id}",getBook).Methods("GET")
	// router.HandlerFunc("/api/books",createBook).Methods("POST")
	// router.HandlerFunc("/api/books/{id}",updateBook).Methods("PUT")
	// router.HandlerFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
	// router.Use(contentTypeCheckingMiddleware)

	

	// var Student1 = Student{"Dat","123"}
	// var Student2 = Student{"123","123"}
	// checkUser(studentCheck,Student1)
	// checkUser(studentCheck,Student2)
	fmt.Println("hello")
	
	
}
// var studentCheck = Student{"Dat","123"}
// func checkUser(student1 Student,student2 Student)  {
// 	if student1.name == student2.name && student1.pass == student2.pass{
// 		fmt.Println(student1.name)
// 		fmt.Println(student1.pass)
// 	}else {
// 		fmt.Println("Bạn đăng nhập k đúng")
// 	}
// }
// func contentTypeCheckingMiddleware()  {
// 	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request)  {
// 		reqContentType := r.Header.Get("Content-Type")
// 		if reqContentType != JsonContentType{
// 			fmt.Fprintf("request only allow content type application/json")
// 			return
// 		}
// 		next.ServeHTTP(w,r)
// 	})
// }
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)//get params
	//loop through book and find with id
	for _, item:= range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index, item:=range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID =params["id"]
			books = append(books,book)
			json.NewEncoder(w).Encode(book)
			return
		}
		
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for index, item:=range books{
		if item.ID == params["id"]{
			books = append(books[:index],books[index+1:]...)
			break
		}
		
	}
	json.NewEncoder(w).Encode(books)
}