package model

import "encoding/json"

const ObjectTypeLink = "link"

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

func (object *Object) ToMap() map[string]interface{} {
	bytes, _ := json.Marshal(object)
	retData := map[string]interface{}{}
	json.Unmarshal(bytes, &retData)
	return retData
}
