package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddNode(title string, content string, cid int64, uid int64) (result RestResult) {
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be empty"
		return
	}
	if cid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "neither cid nor uid can be less than zero"
		return
	}
	err := dal.AddNode(title, content, cid, uid)
	result.Wrap(err, true)
	return
}

func SetNode(id int64, title string, content string, cid int64, uid int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not be less than zero"
		return
	}
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be empty"
		return
	}
	if cid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "neither cid nor uid can be less than zero"
		return
	}
	err := dal.SetNode(id, title, content, cid, uid)
	result.Wrap(err, true)
	return
}

func SaveNode(nd Node) (result RestResult) {
	if nd.Id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not be less than zero"
		return
	}
	if nd.Title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be empty"
		return
	}
	if nd.Pid < 0 || nd.Uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "neither cid nor uid can be less than zero"
		return
	}
	err := dal.SaveNode(nd)
	result.Wrap(err, true)
	return
}

func DelNodePlus(nid int64) (result RestResult) {
	if nid < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not be less than zero"
		return
	}
	err := dal.DelNodePlus(nid)
	result.Wrap(err, true)
	return
}

func DelNode(nid int64) (result RestResult) {
	if nid < 0 {
		result.IsValid = false
		result.ErrorMessage = "nid can not be less than zero"
		return
	}
	err := dal.DelNode(nid)
	result.Wrap(err, true)
	return
}

func GetAllNode() (result RestResult) {
	alln := dal.GetAllNode()
	result.Wrap(nil, alln)
	return
}

func GetAllNodeByCid(cid int64, offset int, limit int, ctype int64, path string) (result RestResult) {
	if cid < 0 {
		result.IsValid = false
		result.ErrorMessage = "cid can not be less than zero"
		return
	}
	alln := dal.GetAllNodeByCid(cid, offset, limit, ctype, path)
	result.Wrap(nil, alln)
	return
}

func GetNode(id int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = "id can not be less than zero"
		return
	}
	nod := dal.GetNode(id)
	result.Wrap(nil, nod)
	return
}

func UpdateNode(nid int64, nd Node) (result RestResult) {
	if nid < 0 {
		result.IsValid = false
		result.ErrorMessage = "nid can not be less than zero"
		return
	}
	err := dal.UpdateNode(nid, nd)
	result.Wrap(err, true)
	return
}

func EditNode(nid int64, cid int64, uid int64, title string, content string) (result RestResult) {
	if nid < 0 {
		result.IsValid = false
		result.ErrorMessage = "nid can not be less than zero"
		return
	}
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be empty"
		return
	}
	if cid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "neither cid nor uid can be less than zero"
		return
	}
	err := dal.EditNode(nid, cid, uid, title, content)
	result.Wrap(err, true)
	return
}
