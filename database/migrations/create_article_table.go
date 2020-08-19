package main

import (
	"github.com/jinzhu/gorm"
	// import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "fumi:abc123@/wikigo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println("接続エラー")
	}
	defer db.Close()
}
