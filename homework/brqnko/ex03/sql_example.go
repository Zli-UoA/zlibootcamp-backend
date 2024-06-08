package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// SQLite3データベースに接続
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}

	// 1
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
		name TEXT,
		age INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, "Alice", 20)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, "Bob", 25)
	if err != nil {
		log.Fatal(err)
	}

	// 2
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	// 3
	_, err = db.Exec(`UPDATE users SET name = ?`, "Charlie")
	if err != nil {
		log.Fatal(err)
	}

	// 4
	_, err = db.Exec("DELETE FROM users where id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}

	// 5
	rows, err = db.Query("SELECT * FROM users WHERE id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	// 6
	rows, err = db.Query("SELECT * FROM users WHERE age >= ?", 20)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	// 7
	rows, err = db.Query("SELECT * FROM users WHERE name = ? and age >= 20", "Alice", 20)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	// 8
	rows, err = db.Query("SELECT * FROM users WHERE age >= ? or name LIKE ?", 20, "%b")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, age)
	}

	_, err = db.Exec("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
