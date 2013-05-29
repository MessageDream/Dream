package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func TopicCount() (result RestResult) {
	today, this_week, this_month := dal.TopicCount()
	data := [3]int{today, this_week, this_month}
	result.Wrap(nil, data)
	return
}

func SetTopic(id int64, cid int64, nid int64, uid int64, ctype int64, title string, content string, author string, attachment string) (result RestResult) {
	if id < 0 || cid < 0 || nid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "id and cid and nid and uid can not be less than zero"
		return
	}
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be null"
		return
	}
	err := dal.SetTopic(id, cid, nid, uid, ctype, title, content, author, attachment)
	result.Wrap(err, true)
	return
}

func AddTopic(title string, content string, cid int64, nid int64, uid int64) (result RestResult) {
	if cid < 0 || nid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = " cid and nid and uid can not be less than zero"
		return
	}
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be null"
		return
	}
	err := dal.AddTopic(title, content, cid, nid, uid)
	result.Wrap(err, true)
	return
}

func DelTopic(id int64) (result RestResult) {
	if id < 0 {
		result.IsValid = false
		result.ErrorMessage = " id can not be less than zero"
		return
	}
	err := dal.DelTopic(id)
	result.Wrap(err, true)
	return
}

func GetAllTopic(offset int, limit int, path string) (result RestResult) {
	allt := dal.GetAllTopic(offset, limit, path)
	result.Wrap(nil, allt)
	return
}

func GetAllTopicByCid(cid int64, offset int, limit int, ctype int64, path string) (result RestResult) {
	allt := dal.GetAllTopicByCid(cid, offset, limit, ctype, path)
	result.Wrap(nil, allt)
	return
}

func GetAllTopicByCidNid(cid int64, nid int64, offset int, limit int, ctype int64, path string) (result RestResult) {
	allt := dal.GetAllTopicByCidNid(cid, nid, offset, limit, ctype, path)
	result.Wrap(nil, allt)
	return
}

func GetAllTopicByNid(nodeid int64, offset int, limit int, ctype int64, path string) (result RestResult) {
	allt := dal.GetAllTopicByNid(nodeid, offset, limit, ctype, path)
	result.Wrap(nil, allt)
	return
}

func SearchTopic(content string, offset int, limit int, path string) (result RestResult) {
	allt := dal.SearchTopic(content, offset, limit, path)
	result.Wrap(nil, allt)
	return
}

func GetTopic(id int64) (result RestResult) {
	tp := dal.GetTopic(id)
	result.Wrap(nil, tp)
	return
}

func SaveTopic(tp Topic) (result RestResult) {
	if tp.Cid < 0 || tp.Nid < 0 {
		result.IsValid = false
		result.ErrorMessage = "neither cid nor nid can  be less than zero"
		return
	}
	if tp.Title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be null"
		return
	}
	err := dal.SaveTopic(tp)
	result.Wrap(err, true)
	return
}

func UpdateTopic(tid int64, tp Topic) (result RestResult) {
	if tid < 0 {
		result.IsValid = false
		result.ErrorMessage = "tid can not be less than zero"
		return
	}
	err := dal.UpdateTopic(tid, tp)
	result.Wrap(err, true)
	return
}

func EditTopic(tid int64, nid int64, cid int64, uid int64, title string, content string) (result RestResult) {
	if tid < 0 || cid < 0 || nid < 0 || uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "tid and cid and nid and uid can not be less than zero"
		return
	}
	if title == "" {
		result.IsValid = false
		result.ErrorMessage = "title can not be null"
		return
	}
	err := dal.EditTopic(tid, nid, cid, uid, title, content)
	result.Wrap(err, true)
	return
}
