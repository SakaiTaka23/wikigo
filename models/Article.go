package models

import (
	"database/sql"
	"fmt"

	// import
	_ "github.com/go-sql-driver/mysql"
)

// Article struct
type Article struct {
	ID        int
	Title     string
	Body      string
	Author    string
	UpdatedAt []uint8
}

// Save func
func (a *Article) Save() error {
	db, _ := sql.Open("mysql", "fumi:abc123@/gowiki?charset=utf8")
	defer db.Close()
	//db.Where(Article{Title: a.Title}).FirstOrCreate(&a)
	rows, _ := db.Query("SELECT * FROM articles where title = ? LIMIT 1", a.Title)
	if rows == nil {
		db.Query("INSERT INTO articles (title,body,author,updated_at) values (?,?,?,now())", a.Title, a.Body, a.Author)
	} else {

	}
	println(rows)
	return nil
}

// GetIndex func
func GetIndex() []Article {
	db, err := sql.Open("mysql", "fumi:abc123@/wikigo?charset=utf8")
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

// Find タイトルに応じてデータが存在すれば持ってくる、Articleの形で変換したい
func Find(title string) *sql.Result {
	db, err := sql.Open("mysql", "fumi:abc123@/wikigo?charset=utf8")
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM articles WHERE title = ? LIMIT 1")
	checkErr(err)
	res, err := stmt.Exec(title)
	checkErr(err)
	fmt.Println(res)
	//dbtotype(res)
	return nil
}

// dbで取得したデータをArticleの形に変換

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
