package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id   int8
	Name string
	Age  int8
}
type Students []Student

func main() {
	fmt.Printf("my website")

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
