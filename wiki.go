package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Page struct {
	PageID      int64
	Title       string
	Body        []byte
	CreatedDate time.Time
	UpdatedDate time.Time
	AuthorID    int64
}

type User struct {
	AuthorID       int64
	AuthorUserName string
	AuthorName     string
	CreatedDate    string
}

func main() {
	// Create the DB connection
	dbCXN()

	// Test page load function
	loadPage(1)

	// Test create page function
	new := createPage(Page{
		Title:       "Lorem Ipsum ++",
		Body:        []byte("Sed nec bibendum quam. Nulla auctor euismod sapien, at congue orci tincidunt nec. Sed a tortor pretium, ornare nisl nec, volutpat lacus. Etiam tincidunt nulla ligula, id pretium felis vulputate et. Suspendisse potenti. Nunc non metus eu felis semper pellentesque in dictum magna. Proin quis dignissim eros, vitae interdum lectus. Aenean tempor at dolor quis lacinia. Mauris aliquam massa et lacus dapibus convallis. Donec eget est libero."),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		AuthorID:    1,
	})
	log.Printf("Page ID '%v' created!", new)

	// Test update page function
	edt := editPage(Page{
		PageID:      15,
		Title:       "Lorem Ipsum Edited",
		Body:        []byte("Etiam porta euismod ligula. Morbi varius, dui a finibus vestibulum, risus leo varius odio, vel volutpat elit purus ultrices neque. Vivamus libero risus, gravida vitae nulla ut, gravida suscipit tellus. Donec dapibus placerat orci, sit amet lobortis justo semper in. Nulla tincidunt diam quis viverra malesuada. Curabitur tristique enim eu semper egestas. Mauris interdum malesuada pretium. Etiam enim ligula, tristique in viverra sed, fringilla sit amet orci. Cras mauris libero, accumsan at egestas non, posuere sed felis. Fusce nec est dui. Cras sed auctor lacus. Praesent vel vehicula metus, non ultrices lectus. Duis non molestie ipsum. "),
		UpdatedDate: time.Now(),
	})
	log.Printf("Page ID '%v' updated!", edt)

	// Test last 10 pages
	pages := recentPages()
	for _, p := range pages {
		fmt.Printf("%v, %s \n", p.PageID, p.Title)
	}
	// serve http on port 8080
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbCXN() {
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
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if err != nil {
		log.Fatal(pingErr)
	}
	log.Println("Database is reachable!")
}

func createPage(p Page) int64 {
	result, err := db.Exec("INSERT INTO articles (title, body, createdDate, updatedDate, authorID) VALUES (?, ?, ?, ?, ?)", p.Title, p.Body, p.CreatedDate, p.UpdatedDate, p.AuthorID)
	if err != nil {
		log.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
	}

	return id
}

func loadPage(id int64) Page {
	var p Page

	row := db.QueryRow("SELECT * FROM articles WHERE articleID = 1;")
	if err := row.Scan(&p.PageID, &p.Title, &p.Body, &p.CreatedDate, &p.UpdatedDate, &p.AuthorID); err != nil {
		log.Println(err)
	}
	log.Printf("Article '%s' loaded! \n", p.Title)

	return p
}

func editPage(p Page) sql.Result {
	result, err := db.Exec("UPDATE articles SET title = ?, body = ?, updatedDate = ? WHERE articleID = ?;", p.Title, p.Body, p.UpdatedDate, p.PageID)
	if err != nil {
		log.Println(err)
	}

	return result
}

func recentPages() []Page {
	var pages []Page

	// Get the 10 most recent wiki pages
	rows, err := db.Query("SELECT * FROM articles ORDER BY createdDate DESC LIMIT 10;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	// Populate the Page struct with query results
	for rows.Next() {
		var p Page
		if err := rows.Scan(&p.PageID, &p.Title, &p.Body, &p.CreatedDate, &p.UpdatedDate, &p.AuthorID); err != nil {
			log.Println(err)
		}
		pages = append(pages, p)
	}

	return pages
}

func recentUsers() {
	// Returns the 10 most recent page IDs & Titles

}

func addUser() {
	// Add user code here
}
