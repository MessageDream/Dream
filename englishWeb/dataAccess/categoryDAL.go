package dataAccess

import (
	. "../models"
	"time"
)

func AddCategory(title string, content string) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&Category{Title: title, Content: content, Created: time.Now()})

	return err
}

func SaveCategory(cat Category) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&cat)
	return err
}

func DelCategory(id int64) error {
	q, _ := ConnDb()
	defer q.Close()
	category := GetCategory(id)
	_, err := q.Delete(&category)

	return err
}

func GetAllCategory() (allc []*Category) {
	q, _ := ConnDb()
	defer q.Close()
	q.FindAll(&allc)
	return allc
}

func GetCategory(id int64) (category Category) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&category)
	return category
}

func UpdateCategory(cid int64, cg Category) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.WhereEqual("id", int64(cid)).Update(&cg)
	return err
}
