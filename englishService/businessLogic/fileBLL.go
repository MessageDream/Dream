package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddFile(ctype int64, location string, url string) (result RestResult) {
	if ctype < 0 {
		result.IsValid = false
		result.ErrorMessage = "ctype can not be less than zero"
		return
	}
	if location == "" || url == "" {
		result.IsValid = false
		result.ErrorMessage = "neither location nor url can be empty"
		return
	}
	err := dal.AddFile(ctype, location, url)
	result.Wrap(err, true)
	return
}

func DelFile(id int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not be less than zero"
		return
	}
	err := dal.DelFile(id)
	result.Wrap(err, true)
	return
}

func GetFile(id int64) (result RestResult) {
	f := dal.GetFile(id)
	result.Wrap(nil, f)
	return
}

func GetAllFile() (result RestResult) {
	f := dal.GetAllFile()
	result.Wrap(nil, f)
	return
}

func GetAllFileByCtype(ctype int64) (result RestResult) {
	f := dal.GetAllFileByCtype(ctype)
	result.Wrap(nil, f)
	return
}

func SaveFile(f File) (result RestResult) {
	if f.Ctype < 0 {
		result.IsValid = false
		result.ErrorMessage = "ctype can not be less than zero"
		return
	}
	if f.Location == "" || f.Url == "" {
		result.IsValid = false
		result.ErrorMessage = "neither location nor url can be empty"
		return
	}
	err := dal.SaveFile(f)
	result.Wrap(err, true)
	return
}

func SetFile(id int64, pid int64, ctype int64, filename string, content string, hash string, location string, url string, size int64) (result RestResult) {
	if ctype < 0 {
		result.IsValid = false
		result.ErrorMessage = "ctype can not be less than zero"
		return
	}
	if location == "" || url == "" {
		result.IsValid = false
		result.ErrorMessage = "neither location nor url can be empty"
		return
	}
	err := dal.SetFile(id, pid, ctype, filename, content, hash, location, url, size)
	result.Wrap(err, true)
	return
}
