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
	// If we'e connected already, just reuse the connection.
	if db != nil {
		return nil
	}

	// Capture connection properties.
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

	// TODO: in real cases there will be bunch of results, this
	// should use "LIMITS" clause or support pagination.
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

func QueryBookById(id int) (*model.Book, error) {
	var book model.Book

	row := db.QueryRow("SELECT * FROM t_books WHERE id = ?", id)
	if err := row.Scan(&book.Id, &book.Title, &book.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("queryBookById %d: no such book", id)
		}

		return nil, fmt.Errorf("queryBookById %d: %v", id, err)
	}

	return &book, nil
}

func QueryAllComments() ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := db.Query("SELECT * FROM t_comments")
	if err != nil {
		return nil, fmt.Errorf("queryAllComments: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.Id, &comment.BookId, &comment.Content); err != nil {
			return nil, fmt.Errorf("queryAllComments: %v", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryAllComments: %v", err)
	}

	return comments, nil
}

func QueryCommentsByBookId(bookId int) ([]model.Comment, error) {
	var comments []model.Comment

	rows, err := db.Query("SELECT * FROM t_comments WHERE book_id = ?", bookId)
	if err != nil {
		return nil, fmt.Errorf("queryCommentsByBookId: %v", err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.Id, &comment.BookId, &comment.Content); err != nil {
			return nil, fmt.Errorf("queryCommentsByBookId: %v", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryCommentsByBookId: %v", err)
	}

	return comments, nil
}
