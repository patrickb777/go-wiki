package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Page struct {
	ArticleID   int
	Title       string
	Body        []byte
	CreatedDate time.Time
	UpdatedDate time.Time
	AuthorID    int
}

type User struct {
	AuthorID       int
	AuthorUserName string
	AuthorName     string
	CreatedDate    string
}

func main() {
	db := dbCXN()
	wikiPage, err := loadPage(1, db)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(wikiPage.Title)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbCXN() (db *sql.DB) {
	// Database connection properties
	cfg := mysql.Config{
		User:                 "app",
		Passwd:               "jVqb2aren2Gm", //Hardcoded for testing
		Net:                  "tcp",
		Addr:                 "172.17.0.2:3306",
		DBName:               "wiki",
		ParseTime:            true,
		AllowNativePasswords: true, // Stackoverflow fix for connection issue
	}

	// Get database handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if err != nil {
		log.Fatal(pingErr)
	}
	log.Print("Database is reachable!")

	return db
}

func loadPage(id int, db *sql.DB) (Page, error) {
	var p Page

	row := db.QueryRow("SELECT * FROM articles WHERE articleID = 1;")
	if err := row.Scan(&p.ArticleID, &p.Title, &p.Body, &p.CreatedDate, &p.UpdatedDate, &p.AuthorID); err != nil {
		return p, err
	}
	return p, nil
}
