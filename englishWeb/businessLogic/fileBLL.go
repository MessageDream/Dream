package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func AddFile(ctype int64, location string, url string) error {
	if ctype < 0 {
		return errors.New("ctype can not be less than zero")
	}
	if location == "" || url == "" {
		return errors.New("neither location nor url can be empty")
	}
	return dal.AddFile(ctype, location, url)
}

func DelFile(id int64) error {
	if id < 0 {
		return errors.New("id can not be less than zero")
	}
	return dal.DelFile(id)
}

func GetFile(id int64) (f File) {
	f = dal.GetFile(id)
	return
}

func GetAllFile() (f []*File) {
	f = dal.GetAllFile()
	return
}

func GetAllFileByCtype(ctype int64) (f []*File) {
	f = dal.GetAllFileByCtype(ctype)
	return
}

func SaveFile(f File) error {
	if f.Ctype < 0 {
		return errors.New("ctype can not be less than zero")
	}
	if f.Location == "" || f.Url == "" {
		return errors.New("neither location nor url can be empty")
	}
	return dal.SaveFile(f)
}

func SetFile(id int64, pid int64, ctype int64, filename string, content string, hash string, location string, url string, size int64) error {
	if ctype < 0 {
		return errors.New("ctype can not be less than zero")
	}
	if location == "" || url == "" {
		return errors.New("neither location nor url can be empty")
	}
	return dal.SetFile(id, pid, ctype, filename, content, hash, location, url, size)
}
