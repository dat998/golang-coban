package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-diver/mysql"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	fmt.Println("Go MySQL Tutorial")
	db, err := sql.Open("mysql", "root:passeord@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Successfully connected to mysql database")
	//
	// baseURL := "https://template-homedecor.onshopbase.com/collections/new-arrivals"
	// config := &tls.Config{
	// 	InsecureSkipVerify: true,
	// }
	// transport := &http.Transport{
	// 	TLSClientConfig: config,
	// }
	// netClient := &http.Client{
	// 	Transport: transport,
	// }
	// response, err := netClient.Get(baseURL)
	// checkErr(err)
	// body, err := ioutil.ReadAll(response.Body)
	// checkErr(err)
	// //fmt.Println(string(body))
	// response.Body.Close()

	// //hello
	// //var hello = "hello dat"
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "hi")
	// })
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, string(body))
	// })
	// fmt.Print("hello")
	// http.ListenAndServe(":8080", nil)

	// db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()
	// insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
