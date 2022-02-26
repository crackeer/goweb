package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Object struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Tag        string `json:"tag"`
	Type       string `json:"type"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (Object) TableName() string {
	return "object"
}
func main() {
	db, err := gorm.Open(sqlite.Open("you.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//err = db.AutoMigrate(&Object{})
	db.Create(&Object{
		Title:   "test",
		Content: "test",
	})
	data := map[string]interface{}{
		"title":   "simple",
		"content": "think",
	}
	db.Table("objet").Create(&data)
	list := []map[string]interface{}{}
	db.Table("object").Find(&list)

	for _, v := range list {
		fmt.Println(v)
	}
}
