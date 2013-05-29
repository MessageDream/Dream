package dataAccess

import (
	. "../models"
	"time"
)

func AddReply(tid int64, uid int64, content string, author string, email string, website string) error {
	q, _ := ConnDb()
	defer q.Close()
	if _, err := q.Save(&Reply{Pid: tid, Uid: uid, Content: content, Created: time.Now(), Author: author, Email: email, Website: website}); err != nil {
		return err
	}

	type Topic struct {
		ReplyTime       time.Time
		ReplyCount      int64
		ReplyLastUserId int64
	}

	if _, err := q.WhereEqual("id", tid).Update(&Topic{ReplyTime: time.Now(), ReplyCount: int64(len(GetReplyByPid(tid, 0, 0, "id"))), ReplyLastUserId: uid}); err != nil {
		return err
	}
	/*
		tp := GetTopic(tid)
		tp.ReplyCount = int64(len(GetReplyByPid(tid, 0, 0, "id")))
		tp.ReplyTime = time.Now()
		tp.ReplyLastUserId = int64(uid)
		if _, err := q.Save(&tp); err != nil {
			return err
		}
	*/
	return nil
}

func DelReply(tid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	reply := GetReply(tid)
	_, err := q.Delete(&reply)

	return err
}

func GetReply(id int64) (reply Reply) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&reply)
	return reply
}

func GetReplyByPid(tid int64, offset int, limit int, path string) (allr []*Reply) {
	q, _ := ConnDb()
	defer q.Close()
	if tid == 0 {
		q.Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allr)
	} else {
		//最热回复
		//q.Where("pid=?", tid).Offset(offset).Limit(limit).OrderByDesc("hotness").FindAll(&allr)
		q.WhereEqual("pid", tid).Offset(offset).Limit(limit).OrderByDesc(path).FindAll(&allr)
	}
	return allr
}
