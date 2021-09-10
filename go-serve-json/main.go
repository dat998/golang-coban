package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type Student struct {
	Id   int8
	Name string
	Age  int8
}
type Product struct{
	Id  	int `json: "id"`
	Name 	string `json: "name"`
	Price 	int `json: "price"`
}
type Students []Student
// var Products []Product
var db *sql.DB

func main() {
	conn()
	//get(db)
	//http.HandleFunc("/getdb",get)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/api/music", MusicPageAPI)
	http.HandleFunc("/api/student", StudentAPI)
	http.HandleFunc("/api/students", StudentsAPI)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
func StudentAPI(w http.ResponseWriter, r *http.Request) {
	var student = Student{1, "diep", 18}
	json.NewEncoder(w).Encode(student)
}
func StudentsAPI(w http.ResponseWriter, r *http.Request) {
	var listStudent = Students{
		Student{1, "diep", 28},
		Student{2, "dat", 23},
	}
	json.NewEncoder(w).Encode(listStudent)
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home page")
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This about page")
}
func MusicPageAPI(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"name": "heloo",
		"sing": "fhsvi",
	}
	json.NewEncoder(w).Encode(data)
}
func get(db *sql.DB){
	
	result,err:=db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}
	for result.Next(){
		var product Product

		err = result.Scan(&product.Id,&product.Name,&product.Price)
		if err != nil {
			panic(err.Error())
		}
		// Products  += append(product.Id,product.Name,product.Price)
		fmt.Println(product.Id)
		fmt.Println(product.Name)
		fmt.Println(product.Price)
		// fmt.Println(err)
	}
	defer result.Close();
	//json.NewEncoder(w).Encode(result)
}
func conn(){
	fmt.Printf("my website")
	db,err:=sql.Open("mysql","dat:@Dat123456@/product")
	if err != nil {
		panic(err.Error())
	}
	pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}