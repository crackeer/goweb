package model

import (
	"errors"
	"fmt"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
)

const (
	insertSQL = "INSERT INTO object(title, content, tag, type, create_time, update_time) values(?, ?, ?, ?, ?, ?)"
	updateSQL = "UPDATE object SET title=?, content=?, tag=?, update_time=? WHERE id= ?"
	selectSQL = "SELECT id, title, content, tag, type, create_time, update_time FROM object where type=?"

	querySQL    = "SELECT id, title, content, tag, type, create_time, update_time FROM object where type=? and tag=? order by id asc"
	queryTagSQL = "SELECT distinct(tag) FROM object where type=?"

	queryDiaryListSQL = "SELECT id, title, tag, type, create_time, update_time FROM object where type=? order by id desc limit ? offset ?"

	queryListSQL        = "SELECT id, title, tag, type, create_time, update_time FROM object where type=? and tag=? order by id asc limit ? offset ?"
	queryObjectCountSQL = "SELECT count(*) as count FROM object where type=? and tag=?"

	querySingleSQL = "SELECT id, title, content, tag, type, create_time, update_time FROM object where id = ?"

	deleteObjectSQL = "delete FROM object where id = ?"
)

const (
	pageSize      int64 = 20
	diaryPageSize int64 = 365
)

// Update
//  @receiver object
//  @return error
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

// GetAll
//  @param objectType
//  @return []Object
//  @return error
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

func GetTheDiary(objectType string, theTag string) (*Object, error) {
	list := []*Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(querySQL, objectType, theTag)
	if err != nil {
		return nil, err
	}

	var (
		id                                                   int64
		title, content, tag, objType, createTime, updateTime string
	)
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

// GetTags
//  @param objectType
//  @return []string
//  @return error
func GetTags(objectType string) ([]string, error) {
	list := []string{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(queryTagSQL, objectType)
	if err != nil {
		return nil, err
	}

	var (
		tag string
	)
	for rows.Next() {
		if err := rows.Scan(&tag); err == nil {
			list = append(list, tag)
		}
	}
	rows.Close()
	db.Close()
	return list, nil
}

// GetObjectList
//  @param objectType
//  @param tag
//  @param page
//  @return []string
//  @return error
func GetObjectList(objectType string, queryTag string, page int64) ([]*Object, int64, error) {
	list := []*Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(queryListSQL, objectType, queryTag, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}

	var (
		id, total                                   int64
		title, tag, objType, createTime, updateTime string
	)
	rows.Scan()
	for rows.Next() {
		if err := rows.Scan(&id, &title, &tag, &objType, &createTime, &updateTime); err == nil {
			list = append(list, &Object{
				ID:         id,
				Title:      title,
				Tag:        tag,
				Type:       objType,
				CreateTime: createTime,
				UpdateTime: updateTime,
			})
		}
	}

	stmt, countError := db.Prepare(queryObjectCountSQL)
	if countError == nil {
		stmt.QueryRow(objectType, queryTag).Scan(&total)
	}

	rows.Close()
	db.Close()
	return list, total, nil
}

// GetDiaryList
//  @param page
//  @return []*Object
//  @return error
func GetDiaryList(page int64) ([]*Object, error) {
	list := []*Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(queryDiaryListSQL, TypeDiary, diaryPageSize, (page-1)*diaryPageSize)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var (
		id                                          int64
		title, tag, objType, createTime, updateTime string
	)
	rows.Scan()
	for rows.Next() {
		if err := rows.Scan(&id, &title, &tag, &objType, &createTime, &updateTime); err == nil {
			list = append(list, &Object{
				ID:         id,
				Title:      title,
				Tag:        tag,
				Type:       objType,
				CreateTime: createTime,
				UpdateTime: updateTime,
			})
		} else {
			fmt.Println(err.Error())
		}
	}
	rows.Close()
	db.Close()
	return list, nil
}

// GetObjectByID
//  @param id
//  @return *Object
//  @return error
func GetObjectByID(id int64) (*Object, error) {
	list := []*Object{}
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	rows, err := db.Query(querySingleSQL, id)
	if err != nil {
		return nil, err
	}

	var (
		xid                                                  int64
		title, content, tag, objType, createTime, updateTime string
	)
	rows.Scan()
	for rows.Next() {
		if err := rows.Scan(&xid, &title, &content, &tag, &objType, &createTime, &updateTime); err == nil {
			list = append(list, &Object{
				ID:         xid,
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

// DeleteObjectByID
//  @param id
//  @return error
func DeleteObjectByID(id int64) error {
	db, _ := container.LockDatabase()
	defer container.UnlockDatabase()

	stmt, err := db.Prepare(deleteObjectSQL)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	stmt.Close()
	db.Close()
	if err != nil {
		return err
	}
	return nil
}
