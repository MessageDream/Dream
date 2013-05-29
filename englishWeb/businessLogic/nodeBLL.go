package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func AddNode(title string, content string, cid int64, uid int64) error {
	if title == "" {
		return errors.New("title can not be empty")
	}
	if cid < 0 || uid < 0 {
		return errors.New("neither cid nor uid can be less than zero")
	}
	return dal.AddNode(title, content, cid, uid)
}

func SetNode(id int64, title string, content string, cid int64, uid int64) error {
	if id < 0 {
		return errors.New("id can not be less than zero")
	}
	if title == "" {
		return errors.New("title can not be empty")
	}
	if cid < 0 || uid < 0 {
		return errors.New("neither cid nor uid can be less than zero")
	}
	return dal.SetNode(id, title, content, cid, uid)
}

func SaveNode(nd Node) error {
	if nd.Id < 0 {
		return errors.New("id can not be less than zero")
	}
	if nd.Title == "" {
		return errors.New("title can not be empty")
	}
	if nd.Pid < 0 || nd.Uid < 0 {
		return errors.New("neither cid nor uid can be less than zero")
	}
	return dal.SaveNode(nd)
}

func DelNodePlus(nid int64) error {
	if nid < 0 {
		return errors.New("id can not be less than zero")
	}
	return dal.DelNodePlus(nid)
}

func DelNode(nid int64) error {
	if nid < 0 {
		return errors.New("nid can not be less than zero")
	}
	return dal.DelNode(nid)
}

func GetAllNode() (alln []*Node) {
	alln = dal.GetAllNode()
	return
}

func GetAllNodeByCid(cid int64, offset int, limit int, ctype int64, path string) (alln []*Node) {
	alln = dal.GetAllNodeByCid(cid, offset, limit, ctype, path)
	return
}

func GetNode(id int64) (node Node) {
	node = dal.GetNode(id)
	return
}

func UpdateNode(nid int64, nd Node) error {
	if nid < 0 {
		return errors.New("nid can not be less than zero")
	}
	return dal.UpdateNode(nid, nd)
}

func EditNode(nid int64, cid int64, uid int64, title string, content string) error {
	if nid < 0 {
		return errors.New("nid can not be less than zero")
	}
	if title == "" {
		return errors.New("title can not be empty")
	}
	if cid < 0 || uid < 0 {
		return errors.New("neither cid nor uid can be less than zero")
	}
	return dal.EditNode(nid, cid, uid, title, content)
}
