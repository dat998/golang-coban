package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	//
	baseURL := "https://template-homedecor.onshopbase.com/collections/new-arrivals"
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: config,
	}
	netClient := &http.Client{
		Transport: transport,
	}
	response, err := netClient.Get(baseURL)
	checkErr(err)
	body, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	fmt.Println(string(body))
	response.Body.Close()

	//hello
	//var hello = "hello dat"
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(body))
	})
	fmt.Print("hello")
	http.ListenAndServe(":3000", nil)

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
