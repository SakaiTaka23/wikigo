package models

import (
	"database/sql"
	"fmt"
	"time"

	// import
	_ "github.com/go-sql-driver/mysql"
)

// Article struct
type Article struct {
	ID        int
	Title     string
	Body      string
	Author    string
	UpdatedAt time.Time
}

// SaveInfo func
func (a *Article) SaveInfo() string {
	db, _ := sql.Open("mysql", "fumi:abc123@/wikigo?parseTime=true&charset=utf8")
	defer db.Close()

	var validate string
	if len(a.Title) == 0 || len(a.Body) == 0{
		validate += "タイトルと本文は入力必須です。\n" 
	}
	if len(a.Title) > 20 {
		validate += "タイトルは20文字以下にしてください。\n"
	}
	if len(a.Body) > 150 {
		validate += "本文は150文字以下にしてください。\n"
	} 
	if len(a.Author) > 10 {
		validate += "名前は10文字いかにしてください。\n"
	}

	if validate != "" {
		return validate
	}

	rows, err := db.Query("SELECT COUNT(*) FROM articles where title = ? LIMIT 1", a.Title)
	checkErr(err)
	if !rows.Next() {
		//新規作成
		_, err := db.Query("INSERT INTO articles (title,body,author,updated_at) values (?,?,?,now())", a.Title, a.Body, a.Author)
		checkErr(err)
	} else {
		//修正
		_, err := db.Query("UPDATE articles set title=?,body=?,author=?,updated_at=now() WHERE title=?", a.Title, a.Body, a.Author, a.Title)
		checkErr(err)
	}
	println("success!")
	return ""
}

// GetIndex func
func GetIndex() []Article {
	db, err := sql.Open("mysql", "fumi:abc123@/wikigo?parseTime=true&charset=utf8")
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT title,body,author,updated_at FROM articles ORDER BY updated_at desc")
	checkErr(err)

	var result []Article
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.Title, &a.Body, &a.Author, &a.UpdatedAt)
		checkErr(err)
		result = append(result, a)
	}
	fmt.Println(result)
	return result
}

// GetInfo func
func GetInfo(title string) Article {
	db, err := sql.Open("mysql", "fumi:abc123@/wikigo?parseTime=true&charset=utf8")
	checkErr(err)
	defer db.Close()
	row, err := db.Query("SELECT title,body,author,updated_at FROM articles WHERE title = ? LIMIT 1", title)

	var a Article
	for row.Next() {
		err := row.Scan(&a.Title, &a.Body, &a.Author, &a.UpdatedAt)
		checkErr(err)
	}
	fmt.Println(a)
	return a
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
