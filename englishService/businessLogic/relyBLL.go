package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddReply(tid int64, uid int64, content string, author string, email string, website string) (result RestResult) {
	if tid < 0 {
		result.IsValid = false
		result.ErrorMessage = "tid can not be  less than zero"
		return
	}
	err := dal.AddReply(tid, uid, content, author, email, website)
	result.Wrap(err, true)
	return
}

func DelReply(tid int64) (result RestResult) {
	if tid < 0 {
		result.IsValid = false
		result.ErrorMessage = "tid can not be less than zero"
		return
	}
	err := dal.DelReply(tid)
	result.Wrap(err, true)
	return
}

func GetReply(id int64) (result RestResult) {
	reply := dal.GetReply(id)
	result.Wrap(nil, reply)
	return
}

func GetReplyByPid(tid int64, offset int, limit int, path string) (result RestResult) {
	allr := dal.GetReplyByPid(tid, offset, limit, path)
	result.Wrap(nil, allr)
	return
}
