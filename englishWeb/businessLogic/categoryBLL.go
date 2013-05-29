package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func AddCategory(title string, content string) error {
	if title == "" {
		return errors.New("title is empty")
	}
	return dal.AddCategory(title, content)
}

func SaveCategory(cat Category) error {
	return dal.SaveCategory(cat)
}

func DelCategory(id int64) error {
	if id < 0 {
		return errors.New("id can not less than zero")
	}
	return dal.DelCategory(id)
}

func GetAllCategory() (allc []*Category) {
	allc = dal.GetAllCategory()
	return
}

func GetCategory(id int64) (category Category) {
	category = dal.GetCategory(id)
	return
}

func UpdateCategory(cid int64, cg Category) error {
	if cid < 0 {
		return errors.New("id can not less than zero")
	}
	return dal.UpdateCategory(cid, cg)
}
