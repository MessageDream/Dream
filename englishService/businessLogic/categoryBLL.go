package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddCategory(title string, content string) (result RestResult) {
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title is empty"
		return
	}

	err := dal.AddCategory(title, content)
	result.Wrap(err, true)
	return
}

func SaveCategory(cat Category) (result RestResult) {
	err := dal.SaveCategory(cat)
	result.Wrap(err, true)
	return
}

func DelCategory(id int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not less than zero"
		return
	}
	err := dal.DelCategory(id)
	result.Wrap(err, true)
	return
}

func GetAllCategory() (result RestResult) {
	cas := dal.GetAllCategory()
	result.Wrap(nil, cas)
	return
}

func GetCategory(id int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not less than zero"
		return
	}
	ca := dal.GetCategory(id)
	result.Wrap(nil, ca)
	return
}

func UpdateCategory(cid int64, cg Category) (result RestResult) {
	if cid < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not less than zero"
		return
	}
	err := dal.UpdateCategory(cid, cg)
	result.Wrap(err, true)
	return
}
