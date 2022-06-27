package model

import (
	"encoding/json"

	"gorm.io/gorm"
)

const (
	TypeLink  = "link"
	TypeMD    = "markdown"
	TypeShare = "share"
	TypeImage = "image"
)

/*
CREATE TABLE object(
  id    INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT,
  content TEXT,
  tag TEXT,
  type TEXT,
  create_time text,
  update_time text
);
*/
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

func (object *Object) ToMap() map[string]interface{} {
	bytes, _ := json.Marshal(object)
	retData := map[string]interface{}{}
	json.Unmarshal(bytes, &retData)
	return retData
}

// AllTables ...
//  @param key
//  @return []map
func AllTables(db *gorm.DB) []map[string]interface{} {
	list := []map[string]interface{}{}
	db.Table("sqlite_master").Where(map[string]interface{}{
		"type": "table",
	}).Find(&list)
	return list
}

// ExecSQL
//  @param db
//  @return []map
func ExecSQL(db *gorm.DB, sql string) error {
	return db.Exec(sql).Error
}
