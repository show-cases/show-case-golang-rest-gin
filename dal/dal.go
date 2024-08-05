package dal

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/jiafangtao/showcases/model"
)

var db *sql.DB

func Connect() error {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root", //TODO
		Passwd: "root", //TODO
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "showcases",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return pingErr
	}

	fmt.Println("Connected!")
	return nil
}

func Disconnect() error {
	return db.Close()
}

func QueryAllBooks() ([]model.Book, error) {

	var books []model.Book

	rows, err := db.Query("SELECT * FROM t_books")
	if err != nil {
		return nil, fmt.Errorf("queryAllBooks: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Description); err != nil {
			return nil, fmt.Errorf("queryAllBooks: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryAllBooks: %v", err)
	}

	return books, nil
}
