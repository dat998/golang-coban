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
	db, err:=sql.Open("mysql","dat:@Dat123456@/product")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close();

	fmt.Println("Successfull connect to mysql database");
	product:= Product{2,"dien thoai2",30000000}
	insertDB(db,product)

	
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
func insertDB(db *sql.DB, product Product){
	_,err:=db.Query("insert into products(id, name, price) values(?,?,?)",product.Id,product.Name,product.Price)
	if err != nil {
		panic(err.Error())
	}
	
}