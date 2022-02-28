package model

import (
	"errors"
	"fmt"

	"github.com/crackeer/goweb/common"
	"github.com/crackeer/goweb/container"
)

// Update
//  @receiver object
//  @return error
func (object *Object) Update() error {
	if object.ID < 1 {
		retData := &Object{
			Title:      object.Title,
			Content:    object.Content,
			Tag:        object.Tag,
			Type:       object.Type,
			CreateTime: common.GetNowTimeString(),
			UpdateTime: common.GetNowTimeString(),
		}

		return container.GetDatabase().Create(retData).Error
	}

	return container.GetDatabase().Model(&Object{}).Where(map[string]interface{}{
		"id": object.ID,
	}).Updates(map[string]interface{}{
		"title":       object.Title,
		"content":     object.Content,
		"tag":         object.Tag,
		"update_time": common.GetNowTimeString(),
	}).Error
}

// ShareObject
//  @param objectID
//  @param expire
//  @return string
//  @return error
func ShareObject(objectID int64, expire int64) (string, error) {

	object, _ := GetObjectByID(objectID)
	if object.ID < 1 {
		return "", errors.New("不存在记录")
	}

	retData := &Object{
		Title:      fmt.Sprintf("%d", object.ID),
		Content:    fmt.Sprintf("%d", expire),
		Tag:        common.RandomString(10),
		Type:       TypeShare,
		CreateTime: common.GetNowTimeString(),
		UpdateTime: common.GetNowTimeString(),
	}

	err := container.GetDatabase().Create(retData).Error

	return retData.Tag, err
}

// GetAll
//  @param objectType
//  @return []Object
//  @return error
func GetAll(objectType string) ([]*Object, error) {

	list := []*Object{}
	tx := container.GetDatabase().Model(&Object{}).Where(map[string]interface{}{
		"type": objectType,
	})

	err := tx.Select([]string{
		"id", "title", "content", "tag", "type", "create_time", "update_time",
	}).Find(&list).Error
	return list, err
}

// GetObjectByTag
//  @param objectType
//  @param theTag
//  @return *Object
//  @return error
func GetObjectByTag(objectType string, theTag string) (*Object, error) {
	return QueryOne(map[string]interface{}{
		"type": objectType,
		"tag":  theTag,
	})
}

// GetTags
//  @param objectType
//  @return []string
//  @return error
func GetTags(objectType string) ([]string, error) {
	list := []string{}
	err := container.GetDatabase().Model(&Object{}).Where(map[string]interface{}{
		"type": objectType,
	}).Select("distinct tag").Pluck("tag", &list).Error
	return list, err
}

// GetObjectList
//  @param objectType
//  @param tag
//  @param page
//  @return []string
//  @return error
func GetObjectList(objectType, tag, keyword string, page int64, pageSize int64) ([]*Object, int64, error) {
	list := []*Object{}
	var total int64
	offset := (page - 1) * pageSize
	tx := container.GetDatabase().Model(&Object{})
	if len(tag) > 0 {
		tx = tx.Where("tag = ?", tag).Where("type = ?", objectType)
	} else if len(keyword) > 0 {
		tx = tx.Where("type = ? and (title like ? or content like ?)", objectType, fmt.Sprintf("%%%s%%", keyword), fmt.Sprintf("%%%s%%", keyword))
	}
	tx.Count(&total)

	err := tx.Select([]string{
		"id", "title", "tag", "type", "create_time", "update_time",
	}).Offset(int(offset)).Limit(int(pageSize)).Order("id desc").Find(&list).Error
	return list, total, err
}

// GetObjectByID
//  @param id
//  @return *Object
//  @return error
func GetObjectByID(id int64) (*Object, error) {
	return QueryOne(map[string]interface{}{
		"id": id,
	})
}

// DeleteObjectByID
//  @param id
//  @return error
func DeleteObjectByID(id int64) error {
	return container.GetDatabase().Where(map[string]interface{}{
		"id": id,
	}).Delete(&Object{}).Error
}

// QueryOne
//  @param query
//  @return *Object
//  @return error
func QueryOne(query map[string]interface{}) (*Object, error) {
	retData := &Object{}
	err := container.GetDatabase().Model(&Object{}).Where(query).Order("id desc").First(retData).Error
	return retData, err
}

// QueryList
//  @param query
//  @return []Object
//  @return error
func QueryList(query map[string]interface{}) ([]Object, error) {
	retData := []Object{}
	err := container.GetDatabase().Model(&Object{}).Where(query).Order("id desc").Find(&retData).Error
	return retData, err
}
