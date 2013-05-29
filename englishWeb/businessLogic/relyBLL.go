package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func AddReply(tid int64, uid int64, content string, author string, email string, website string) error {
	if tid < 0 {
		return errors.New("tid can not be  less than zero")
	}
	return dal.AddReply(tid, uid, content, author, email, website)
}

func DelReply(tid int64) error {
	if tid < 0 {
		return errors.New("tid can not be less than zero")
	}
	return dal.DelReply(tid)
}

func GetReply(id int64) (reply Reply) {
	reply = dal.GetReply(id)
	return
}

func GetReplyByPid(tid int64, offset int, limit int, path string) (allr []*Reply) {
	allr = dal.GetReplyByPid(tid, offset, limit, path)
	return
}
