package model

import (
	"errors"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
)

const (
	insertSQL = "INSERT INTO object(title, content, tag, type, create_time, update_time) values(?, ?, ?, ?, ?, ?)"
	updateSQL = "UPDATE object SET title=?, content=?, tag=?, update_time=? WHERE id= ?"
	selectSQL = "SELECT id, title, content, tag, type, create_time, update_time FROM object where type=?"

	querySQL = "SELECT id, title, content, tag, type, create_time, update_time FROM object where type=? and tag=? and title=? order by id asc"
)

func (object *Object) Update() error {

	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	if object.ID < 1 {

		if len(object.Type) < 1 {
			return errors.New("please set object type")
		}
		stmt, err := db.Prepare(insertSQL)
		if err != nil {
			return err
		}
		result, err := stmt.Exec(object.Title, object.Content, object.Tag, object.Type, common.GetNowTimeString(), common.GetNowTimeString())
		stmt.Close()
		db.Close()
		if err != nil {
			return err
		}
		object.ID, _ = result.LastInsertId()
		return nil
	}

	stmt, err := db.Prepare(updateSQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(object.Title, object.Content, object.Tag, common.GetNowTimeString(), object.ID)
	stmt.Close()
	db.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetAll(objectType string) ([]Object, error) {
	list := []Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(selectSQL, objectType)
	if err != nil {
		return list, err
	}

	var (
		id                                                   int64
		title, content, tag, objType, createTime, updateTime string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &title, &content, &tag, &objType, &createTime, &updateTime); err == nil {
			list = append(list, Object{
				ID:         id,
				Title:      title,
				Content:    content,
				Tag:        tag,
				Type:       objType,
				CreateTime: createTime,
				UpdateTime: updateTime,
			})
		}
	}
	rows.Close()
	db.Close()
	return list, nil
}

func GetTheOne(objectType string, theTag, theTitle string) (*Object, error) {
	list := []*Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(querySQL, objectType, theTag, theTitle)
	if err != nil {
		return nil, err
	}

	var (
		id                                                   int64
		title, content, tag, objType, createTime, updateTime string
	)
	rows.Scan()
	for rows.Next() {
		if err := rows.Scan(&id, &title, &content, &tag, &objType, &createTime, &updateTime); err == nil {
			list = append(list, &Object{
				ID:         id,
				Title:      title,
				Content:    content,
				Tag:        tag,
				Type:       objType,
				CreateTime: createTime,
				UpdateTime: updateTime,
			})
		}
	}
	rows.Close()
	db.Close()
	if len(list) > 0 {
		return list[len(list)-1], nil
	}
	return nil, nil
}
