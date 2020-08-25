package models

import (
	"time"

	"github.com/jinzhu/gorm"

	// import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Article struct
type Article struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Title     string `gorm:"unique;not null"`
	Body      string `gorm:"not null"`
	Author    string `gorm:"default:'匿名'"`
	UpdatedAt time.Time
}

// Save func
func (a *Article) Save() error {
	db, _ := gorm.Open("mysql", "fumi:abc123@/gowiki?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	db.Where(Article{Title: a.Title}).FirstOrCreate(&a)
	return nil
}

// Find func
func Find(title string) *Article {
	db, _ := gorm.Open("mysql", "fumi:abc123@/gowiki?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	var a Article
	result := db.Where("Title = ?", title).Take(&a)
	println(result)
	if result == nil {
		return nil
	}
	return nil
}
