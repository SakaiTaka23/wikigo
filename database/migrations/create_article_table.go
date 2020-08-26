package main

import "database/sql"

func main() {
	db, err := sql.Open("mysql", "fumi:abc123@/wikigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println("接続エラー")
	}
	defer db.Close()
	
}
