package dataAccess

import (
	"./common"
	. "../models"
	"fmt"
	"github.com/coocood/qbs"
	"os"
	"time"
)

func TopicCount() (today int, this_week int, this_month int) {
	q, _ := ConnDb()
	defer q.Close()
	var topict, topicw, topicm []*Topic
	k := time.Now()

	//一天之前
	d, _ := time.ParseDuration("-24h")
	t := k.Add(d)
	e := q.Where("created>?", t).FindAll(&topict)
	if e != nil {
		today = 0
		fmt.Println(e)
	} else {
		today = len(topict)
	}

	//一周之前
	w := k.Add(d * 7)
	e = q.Where("created>?", w).FindAll(&topicw)
	if e != nil {
		this_week = 0
		fmt.Println(e)
	} else {
		this_week = len(topicw)
	}

	//一月之前
	m := k.Add(d * 30)
	e = q.Where("created>?", m).FindAll(&topicm)
	if e != nil {
		this_month = 0
		fmt.Println(e)
	} else {
		this_month = len(topicm)
	}

	return today, this_week, this_month
}

func SetTopic(id int64, cid int64, nid int64, uid int64, ctype int64, title string, content string, author string, attachment string) error {
	q, _ := ConnDb()
	defer q.Close()
	var tp Topic
	if q.WhereEqual("id", id).Find(&tp); tp.Id == 0 {
		_, err := q.Save(&Topic{Id: id, Cid: cid, Nid: nid, Uid: uid, Ctype: ctype, Title: title, Content: content, Author: author, Attachment: attachment})
		return err
	} else {
		type Topic struct {
			Cid        int64
			Nid        int64
			Uid        int64
			Ctype      int64
			Title      string
			Content    string
			Author     string
			Attachment string
		}

		_, err := q.WhereEqual("id", id).Update(&Topic{Cid: cid, Nid: nid, Uid: uid, Ctype: ctype, Title: title, Content: content, Author: author, Attachment: attachment})
		return err
	}
	return nil
}

func AddTopic(title string, content string, cid int64, nid int64, uid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	if _, err := q.Save(&Topic{Cid: cid, Nid: nid, Title: title, Content: content, Created: time.Now()}); err != nil {
		return err
	}

	type Node struct {
		TopicTime       time.Time
		TopicCount      int64
		TopicLastUserId int64
	}

	if _, err := q.WhereEqual("id", nid).Update(&Node{TopicTime: time.Now(), TopicCount: int64(len(GetAllTopicByNid(nid, 0, 0, 0, "id"))), TopicLastUserId: uid}); err != nil {
		return err
	}
	/*
		nd := GetNode(nid)
		nd.TopicTime = time.Now()
		nd.TopicCount = int64(len(GetAllTopicByNid(nid, 0, 0, "id")))
		nd.TopicLastUserId = int64(uid)
		if _, err := q.Save(&nd); err != nil {
			return err
		}
	*/
	return nil
}

func DelTopic(id int64) error {
	q, _ := ConnDb()
	defer q.Close()
	topic := GetTopic(id)
	if common.Exist("." + topic.Attachment) {
		if err := os.Remove("." + topic.Attachment); err != nil {
			//return err
			//可以输出错误，但不要反回错误，以免陷入死循环无法删掉
			fmt.Println("DEL TOPIC", id, err)
		}
	}

	//不管实际路径中是否存在文件均删除该数据库记录，以免数据库记录陷入死循环无法删掉
	_, err := q.Delete(&topic)

	return err
}

func GetAllTopic(offset int, limit int, path string) (allt []*Topic) {
	q, _ := ConnDb()
	defer q.Close()
	q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("created").FindAll(&allt)
	return allt
}

func GetAllTopicByCid(cid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {
	//排序首先是热值优先，然后是时间优先。
	q, _ := ConnDb()
	defer q.Close()

	switch {
	case path == "asc":
		if ctype != 0 {
			condition := qbs.NewCondition("cid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).FindAll(&allt)

		} else {
			q.Where("cid=?", cid).Offset(offset).Limit(limit).FindAll(&allt)

		}
	case path == "views" || path == "reply_count":
		if ctype != 0 {
			condition := qbs.NewCondition("cid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&allt)

		} else {
			if cid == 0 {
				q.OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&allt)
			} else {
				q.WhereEqual("cid", cid).OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&allt)
			}

		}
	default:
		if ctype != 0 {

			condition := qbs.NewCondition("cid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

		} else {
			if cid == 0 {
				q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

			} else {
				q.WhereEqual("cid", cid).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)
			}
		}

	}
	return allt
}

func GetAllTopicByCidNid(cid int64, nid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {

	q, _ := ConnDb()
	defer q.Close()

	switch {
	case path == "asc":
		if ctype != 0 {
			condition := qbs.NewCondition("cid=?", cid).And("nid=?", nid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).FindAll(&allt)

		} else {

			condition := qbs.NewCondition("cid=?", cid).And("nid=?", nid)
			q.Condition(condition).Offset(offset).Limit(limit).FindAll(&allt)

		}
	default:
		if ctype != 0 {
			condition := qbs.NewCondition("cid=?", cid).And("nid=?", nid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

		} else {

			condition := qbs.NewCondition("cid=?", cid).And("nid=?", nid)
			q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

		}

	}
	return allt
}

func GetAllTopicByNid(nodeid int64, offset int, limit int, ctype int64, path string) (allt []*Topic) {
	//排序首先是热值优先，然后是时间优先。
	q, _ := ConnDb()
	defer q.Close()

	switch {
	case path == "asc":
		if nodeid == 0 {
			//q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)
			return nil
		} else {
			if ctype != 0 {
				condition := qbs.NewCondition("nid=?", nodeid).And("ctype=?", ctype)
				q.Condition(condition).Offset(offset).Limit(limit).FindAll(&allt)

			} else {
				q.Where("nid=?", nodeid).Offset(offset).Limit(limit).FindAll(&allt)

			}
		}
	default:
		if nodeid == 0 {
			//q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)
			return nil
		} else {
			if ctype != 0 {
				condition := qbs.NewCondition("nid=?", nodeid).And("ctype=?", ctype)
				q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

			} else {
				q.Where("nid=?", nodeid).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)

			}
		}
	}
	return allt
}

func SearchTopic(content string, offset int, limit int, path string) (allt []*Topic) {
	//排序首先是热值优先，然后是时间优先。
	if content != "" {
		q, _ := ConnDb()
		defer q.Close()
		keyword := "%" + content + "%"
		condition := qbs.NewCondition("title like ?", keyword).Or("content like ?", keyword)
		q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("reply_count").OrderByDesc("created").FindAll(&allt)
		//q.Where("title like ?", keyword).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("created").FindAll(&allt)
		return allt
	}
	return nil
}

func GetTopic(id int64) (topic Topic) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&topic)
	return topic
}

func SaveTopic(tp Topic) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&tp)
	return err
}

func UpdateTopic(tid int64, tp Topic) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.WhereEqual("id", int64(tid)).Update(&tp)
	return err
}

func EditTopic(tid int64, nid int64, cid int64, uid int64, title string, content string) error {
	tpc := GetTopic(tid)
	tpc.Cid = int64(cid)
	tpc.Nid = int64(nid)
	tpc.Title = title
	tpc.Content = content
	tpc.Updated = time.Now()

	if err := UpdateTopic(tid, tpc); err != nil {
		return err
	}

	q, _ := ConnDb()
	defer q.Close()

	type Node struct {
		TopicTime       time.Time
		TopicCount      int64
		TopicLastUserId int64
	}

	if _, err := q.WhereEqual("id", nid).Update(&Node{TopicTime: tpc.Created, TopicCount: int64(len(GetAllTopicByNid(nid, 0, 0, 0, "id"))), TopicLastUserId: int64(uid)}); err != nil {
		return err
	}

	return nil
}
