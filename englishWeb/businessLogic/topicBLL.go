package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func TopicCount() (today int, this_week int, this_month int) {
	today, this_week, this_month = dal.TopicCount()
	return
}

func SetTopic(id int64, cid int64, nid int64, uid int64, ctype int64, title string, content string, author string, attachment string) error {
	if id < 0 || cid < 0 || nid < 0 || uid < 0 {
		return errors.New("id and cid and nid and uid can not be less than zero")
	}
	if title == "" {
		return errors.New("title can not be null")
	}
	return dal.SetTopic(id, cid, nid, uid, ctype, title, content, author, attachment)
}

func AddTopic(title string, content string, cid int64, nid int64, uid int64) error {
	if cid < 0 || nid < 0 || uid < 0 {
		return errors.New(" cid and nid and uid can not be less than zero")
	}
	if title == "" {
		return errors.New("title can not be null")
	}
	return dal.AddTopic(title, content, cid, nid, uid)
}

func DelTopic(id int64) error {
	if id < 0 {
		return errors.New(" id can not be less than zero")
	}
	return dal.DelTopic(id)
}

func GetAllTopic(offset int, limit int, path string) (allt []*Topic) {
	allt = dal.GetAllTopic(offset, limit, path)
	return
}

func GetAllTopicByCid(cid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {
	allt = dal.GetAllTopicByCid(cid, offset, limit, ctype, path)
	return
}

func GetAllTopicByCidNid(cid int64, nid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {
	allt = dal.GetAllTopicByCidNid(cid, nid, offset, limit, ctype, path)
	return
}

func GetAllTopicByNid(nodeid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {
	allt = dal.GetAllTopicByNid(nodeid, offset, limit, ctype, path)
	return
}

func SearchTopic(content string, offset int, limit int, path string) (allt []*Topic) {
	allt = dal.SearchTopic(content, offset, limit, path)
	return
}

func GetTopic(id int64) (topic Topic) {
	topic = dal.GetTopic(id)
	return
}

func SaveTopic(tp Topic) error {
	if tp.Cid < 0 || tp.Nid < 0 {
		return errors.New("neither cid nor nid can  be less than zero")
	}
	if tp.Title == "" {
		return errors.New("title can not be null")
	}
	return dal.SaveTopic(tp)
}

func UpdateTopic(tid int64, tp Topic) error {
	if tid < 0 {
		return errors.New("tid can not be less than zero")
	}
	return dal.UpdateTopic(tid, tp)
}

func EditTopic(tid int64, nid int64, cid int64, uid int64, title string, content string) error {
	if tid < 0 || cid < 0 || nid < 0 || uid < 0 {
		return errors.New("tid and cid and nid and uid can not be less than zero")
	}
	if title == "" {
		return errors.New("title can not be null")
	}
	return dal.EditTopic(tid, nid, cid, uid, title, content)
}
