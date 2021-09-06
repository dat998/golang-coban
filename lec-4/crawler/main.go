package main

import (
	// "github.com/gocolly/colly"
	"fmt"
	// "os"
	"log"
	// "encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	
)

type Film struct {
	Name     string `json:"name"`
	Year     string `json:"year"`
	Rating   string `json:"Rating"`
}
func main(){
	//var films []string
	//connect database
	db, err:=sql.Open("mysql","dat:@Dat123456@/crawler")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close();
	fmt.Println("Successfull connect to mysql database");
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	res, _ := http.Get(url)
	data, _ := goquery.NewDocumentFromReader(res.Body)
	data.Find(".lister-list tr").Each(func(i int, s *goquery.Selection) {
		// bien tam de  them du lieu vao
		temp := Film{}
		temp.Rating, _ = s.Find(".posterColumn [name = ir]").Attr("data-value")
		temp.Year = s.Find(".titleColumn span").Text()
		info := s.Find(".titleColumn a")
		temp.Name = info.Text()
		//films[i+1] = temp

		fmt.Println(temp.Name, temp.Year,temp.Rating)
		sqlQuery := "INSERT Items SET movie_name=?, movie_year=?, rate=?"
		//insert,err:=db.Prepare("insert into movie(movie_name,movie_year,rate) values(?,?,?)")
		insert,err:=db.Prepare(sqlQuery)
		defer insert.Close()
		if err != nil {
			panic(err)
		}
		
		result,err2:=insert.Exec(temp.Name,temp.Year,temp.Rating)
		if err2 != nil {
			panic(err2)
		}
		// defer result.Close()
		lastId, err3 := result.LastInsertId() //Lấy ra ID vừa được insert
		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Println(lastId)
		// defer lastId.Close()

	})
	result,err:=db.Query("select movie_name,movie_year,rate from movie")
	if err != nil {
		panic(err.Error())
	}
	for result.Next(){
		var film Film
		err:= result.Scan(&film.Name,&film.Year,&film.Rating)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("-Name: ",film.Name,"-year: ",film.Year,"-rate: ",film.Rating)
	}
	defer result.Close()
	
}
// func ConnectDB(*sql.DB, error){
// 	conn := "tcdat:@Dat123456@/crawler"
// 	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("connect")
// 	return db,nil
// }
// func InsertDB(db *sql.DB, film Film) {
// 	insert,err:=db.Query("insert into movie(movie_name,movie_year,rate) value(?,?,?)")
// 	if err != nil {
// 		panic(err)
// 	}
// 	result,err2:=insert.Exec(film.Name,film.Year,film.rate)
// 	if err2 != nil {
// 		panic(err2)
// 	}
// }
// func SelectDB(db *gorm.DB){
// 	result,err:=db.Query("select movie_name from movie")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for result.Next(){
// 		fmt.Println(result.movie_name)
// 	}
// 	defer result.Close()
// }