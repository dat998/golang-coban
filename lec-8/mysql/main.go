package main

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Review struct {
	Vote    string
	Title   string
	Content string
}

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "dat:@Dat123456@/paypay")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(1200)
	defer db.Close()

	time_start := time.Now()
	ch := make(chan Review)
	var wg sync.WaitGroup

	wg.Add(2)
	go GetReview(&wg, ch)
	go SetReview(db, &wg, ch)
	wg.Wait()
	time_end := time.Now()
	log.Println(time_end.Sub(time_start))

	Search(db)
	log.Println("success")
}

func GetReview(wg *sync.WaitGroup, ch chan Review) {
	defer wg.Done()
	defer close(ch)
	file, err := os.Open(".Downloads/train.csv/train.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		item := Review{
			Vote:    record[0],
			Title:   record[1],
			Content: record[2],
		}
		ch <- item
	}
}

func SetReview(db *sql.DB, wg *sync.WaitGroup, ch chan Review) {
	defer wg.Done()
	for {
		item, ok := <-ch
		if !ok {
			return
		}

		insert, err := db.Query("INSERT INTO reviews(vote, title, content) VALUES ( ?, ?, ? )", item.Vote, item.Title, item.Content)

		if err != nil {
			panic(err.Error())
		}
		insert.Close()
	}
}

func Search(db *sql.DB) {
	time_start := time.Now()
	str := "0"
	var review Review

	err := db.QueryRow("SELECT vote, title, content FROM reviews WHERE content LIKE ?", "%"+str+"%").Scan(&review.Vote, &review.Title, &review.Content)
	if err != nil {
		log.Fatal(err)
		return
	}
	time_end := time.Now()
	log.Println(time_end.Sub(time_start))
	log.Println(review)
}
