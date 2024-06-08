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
	defer db.Close()

	// テーブルの作成
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
)`)
	if err != nil {
		log.Fatal(err)
	}

	//データのインサート
	_, err = db.Exec(`INSERT INTO users (name, age) VALUES (?, ?), (?, ?)`, "Alice", 20, "Bob", 25)
	if err != nil {
		log.Fatal(err)
	}

	//全てのデータを取得
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

	//users テーブルの name カラムを Charlie に更新
	_, err = db.Exec(`UPDATE users SET name = ?`, "Charlie")
	if err != nil {
		log.Fatal(err)
	}

	// users テーブルから id が 2 のデータを削除
	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, 2)
	if err != nil {
		log.Fatal(err)
	}

	//users テーブルから name が Alice のデータを取得
	rows, err = db.Query(`SELECT id, name, age FROM users WHERE name = ?`, "Alice")
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

	//users テーブルから age が 20 以上のデータを取得
	rows, err = db.Query(`SELECT id, name, age FROM users WHERE age >= ?`, 20)
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

	//users テーブルから age が 20 以上かつ name が Alice のデータを取得
	rows, err = db.Query(`SELECT id, name, age FROM users WHERE age >= ? and name = ?`, 20, "Alice")
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

	//users テーブルから age が 20 以上または name が e で終わるデータを取得
	rows, err = db.Query(`SELECT id, name, age FROM users WHERE age >= ? or name like ?`, 20, "%e")
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
}
