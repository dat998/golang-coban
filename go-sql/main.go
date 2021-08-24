package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
type Product struct{
	Id  	int `json: "id"`
	Name 	string `json: "name"`
	Price 	int `json: "price"`
}
func main()  {
	fmt.Println("Go sql")
	db, err:=sql.Open("mysql","tcdat:12345678@/product")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close();

	fmt.Println("Successfull connect to mysql database");


	
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
		fmt.Println(product.Id)
		fmt.Println(product.Name)
		fmt.Println(product.Price)
	}
	defer result.Close();
	
}