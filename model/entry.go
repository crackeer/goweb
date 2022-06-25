package model

import (
	"errors"

	"gorm.io/gorm"
)

// Table
type Table struct {
	Name string
	DB   *gorm.DB

	primaryKey string
}

// NewTable
//  @param name
//  @param db
//  @return *Table
func NewTable(db *gorm.DB, name string) (*Table, error) {
	if db == nil {
		return nil, errors.New("db nil")
	}
	return &Table{
		Name: name,
		DB:   db,
	}, nil
}

// SetPrimaryKey
//  @receiver table
//  @param key
func (table *Table) SetPrimaryKey(key string) {
	table.primaryKey = key
}

// GetPrimaryKey
//  @receiver table
//  @return string
func (table *Table) GetPrimaryKey() string {
	if len(table.primaryKey) > 0 {
		return table.primaryKey
	}

	return "id"
}

// Create
//  @receiver table
//  @param data
//  @return map
func (table *Table) Create(data map[string]interface{}) (map[string]interface{}, error) {
	err := table.DB.Table(table.Name).Create(&data).Error
	return data, err
}

// Update
//  @receiver table
//  @param id
//  @param data
//  @return int64
func (table *Table) Update(id int64, data map[string]interface{}) int64 {
	return table.DB.Table(table.Name).Where(map[string]interface{}{
		"id": id,
	}).Updates(data).RowsAffected
}

// Query
//  @receiver table
//  @param query
//  @return []map
func (table *Table) Query(query map[string]interface{}) []map[string]interface{} {
	list := []map[string]interface{}{}
	table.DB.Table(table.Name).Where(query).Find(&list)
	return list
}

// GetPageList
//  @receiver table
//  @param query
//  @param page
//  @param pageSize
//  @return []map
func (table *Table) GetPageList(query map[string]interface{}, page, pageSize int64) []map[string]interface{} {
	list := []map[string]interface{}{}
	offset := (page - 1) * pageSize
	table.DB.Table(table.Name).Where(query).Offset(int(offset)).Order("id desc").Limit(int(pageSize)).Find(&list)
	return list
}

// Count
//  @receiver table
//  @param query
//  @return int64
func (table *Table) Count(query map[string]interface{}) int64 {
	var count int64
	table.DB.Table(table.Name).Where(query).Count(&count)
	return count
}
