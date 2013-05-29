package dataAccess

import (
	. "../models"
	"./common"
	"fmt"
	"os"
)

func AddFile(ctype int64, location string, url string) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&File{Ctype: ctype, Location: location, Url: url})
	return err
}

func DelFile(id int64) error {
	q, _ := ConnDb()
	defer q.Close()
	f := GetFile(id)

	if common.Exist("." + f.Location) {
		if err := os.Remove("." + f.Location); err != nil {
			return err
			fmt.Println(err)
		}
	}

	//不管实际路径中是否存在文件均删除该数据库记录，以免数据库记录陷入死循环无法删掉
	_, err := q.Delete(&f)
	fmt.Println(err)
	return err
}

func GetFile(id int64) (f File) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&f)
	return f
}

func GetAllFile() (f []*File) {
	q, _ := ConnDb()
	defer q.Close()
	q.OrderByDesc("id").FindAll(&f)
	return f
}

func GetAllFileByCtype(ctype int64) (f []*File) {
	q, _ := ConnDb()
	defer q.Close()
	q.WhereEqual("ctype", ctype).OrderByDesc("id").FindAll(&f)
	return f
}

func SaveFile(f File) error {
	q, _ := ConnDb()
	defer q.Close()
	_, e := q.Save(&f)
	return e
}

func SetFile(id int64, pid int64, ctype int64, filename string, content string, hash string, location string, url string, size int64) error {
	q, _ := ConnDb()
	defer q.Close()
	var f File
	if q.WhereEqual("id", id).Find(&f); f.Id == 0 {
		_, err := q.Save(&File{Id: id, Pid: pid, Ctype: ctype, Filename: filename, Content: content, Hash: hash, Location: location, Url: url, Size: size})
		return err
	} else {
		type File struct {
			Pid      int64
			Ctype    int64
			Filename string
			Content  string
			Hash     string
			Location string
			Url      string
			Size     int64
		}
		_, err := q.WhereEqual("id", id).Update(&File{Pid: pid, Ctype: ctype, Filename: filename, Content: content, Hash: hash, Location: location, Url: url, Size: size})

		return err
	}
	return nil
}
