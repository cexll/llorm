package llorm

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestLLOrm(t *testing.T) {
	db, _ := sql.Open("sqlite3", "ll.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}

func TestOne(t *testing.T) {
	e, _ := NewEngine("sqlite3", "ll.db")
	defer e.Close()
	s := e.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) Values(?), (?)", "Cexll", "iMorta").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
