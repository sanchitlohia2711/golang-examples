package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type book struct {
	id          int
	name        string
	description string
	created     time.Time
}

func initDB() {
	dataSourceString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "yourpassword", "127.0.0.1:3306", "library")
	var err error
	db, err = sql.Open("mysql", dataSourceString)
	if err != nil {
		log.Fatal("Unable to connect to database. Err: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to ping to database. Err: " + err.Error())
	}
}

func createTable() {
	query := `
	CREATE TABLE book (
		id INT AUTO_INCREMENT,
		name VARCHAR(45) NOT NULL,
		description TEXT NOT NULL,
		created DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func insert(book *book) {
	_, err := db.Exec(`INSERT INTO book (name, description, created) VALUES (?, ?, ?)`, book.name, book.description, book.created)
	if err != nil {
		log.Fatal("Error in inserting book record. Err: " + err.Error())
	}

	fmt.Printf("Book inserted with name: %s", book.name)
}

func read(bookname string) []*book {
	var books []*book
	resultDB, err := db.Query(`SELECT * FROM book WHERE name=?`, bookname)
	if err != nil {
		log.Fatal("Error in query execution. Err: " + err.Error())
	}
	for resultDB.Next() {
		book := &book{}
		err = resultDB.Scan(&book.id, book.name, book.description, book.created)
		if err != nil {
			log.Fatal("Error in scan. Err: " + err.Error())
		}
		books = append(books, book)
	}

	fmt.Println("Book read successful")
	return books
}

func delete(bookname string) {
	_, err := db.Exec(`DELETE FROM book WHERE name = ?`, bookname)
	if err != nil {
		log.Fatal("Error while deleting book")
	}
}

func main() {
	initDB()
	createTable()

	book := &book{
		name:        "test1",
		description: "test1 description",
		created:     time.Now(),
	}
	insert(book)
	books := read("test1")

	fmt.Printf("%#v", books)

	delete(book.name)
	fmt.Println("end")
}
