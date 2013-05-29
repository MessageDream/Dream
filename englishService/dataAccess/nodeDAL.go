package dataAccess

import (
	. "../models"
	"github.com/coocood/qbs"
	"time"
)

func AddNode(title string, content string, cid int64, uid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	if _, err := q.Save(&Node{Pid: cid, Uid: uid, Title: title, Content: content, Created: time.Now()}); err != nil {
		return err
	}

	type Category struct {
		NodeTime       time.Time
		NodeCount      int64
		NodeLastUserId int64
	}

	if _, err := q.WhereEqual("id", cid).Update(&Category{NodeTime: time.Now(), NodeCount: int64(len(GetAllNodeByCid(cid, 0, 0, 0, "id"))), NodeLastUserId: uid}); err != nil {
		return err
	}
	/*
		ctr := GetCategory(cid)
		ctr.NodeTime = time.Now()
		ctr.NodeCount = int64(len(GetAllNodeByCid(cid, 0, 0, "id")))
		ctr.NodeLastUserId = int64(uid)
		if _, err := q.Save(&ctr); err != nil {
			return err
		}
	*/
	return nil
}

func SetNode(id int64, title string, content string, cid int64, uid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	var nd Node
	if q.WhereEqual("id", id).Find(&nd); nd.Id == 0 {
		_, err := q.Save(&Node{Id: id, Pid: cid, Uid: uid, Title: title, Content: content})
		return err
	} else {
		type Node struct {
			Pid     int64
			Uid     int64
			Title   string
			Content string
		}

		_, err := q.WhereEqual("id", id).Update(&Node{Pid: cid, Uid: uid, Title: title, Content: content})
		return err
	}
	return nil
}

func SaveNode(nd Node) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&nd)
	return err
}

func DelNodePlus(nid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	node := GetNode(nid)
	_, err := q.Delete(&node)

	for i, v := range GetAllTopicByNid(nid, 0, 0, 0, "id") {
		if i > 0 {
			DelTopic(v.Id)
			for ii, vv := range GetReplyByPid(v.Id, 0, 0, "id") {
				if ii > 0 {
					DelReply(vv.Id)
				}
			}
		}
	}

	return err
}

func DelNode(nid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	node := GetNode(nid)
	_, err := q.Delete(&node)

	return err
}

func GetAllNode() (alln []*Node) {
	q, _ := ConnDb()
	defer q.Close()
	//q.OrderByDesc("id").FindAll(&alln)
	q.OrderByDesc("created").FindAll(&alln)
	return alln
}

func GetAllNodeByCid(cid int64, offset int, limit int, ctype int64, path string) (alln []*Node) {
	//排序首先是热值优先，然后是时间优先。
	q, _ := ConnDb()
	defer q.Close()
	switch {
	case path == "asc":
		if ctype != 0 {
			condition := qbs.NewCondition("pid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).FindAll(&alln)
		} else {
			if cid == 0 {
				q.Offset(offset).Limit(limit).FindAll(&alln)
			} else {
				q.WhereEqual("pid", cid).Offset(offset).Limit(limit).FindAll(&alln)
			}

		}
	case path == "views" || path == "topic_count":
		if ctype != 0 {
			condition := qbs.NewCondition("pid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&alln)

		} else {
			if cid == 0 {
				q.OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&alln)
			} else {
				q.WhereEqual("pid", cid).OrderByDesc(path).Offset(offset).Limit(limit).FindAll(&alln)
			}

		}
	default:
		if ctype != 0 {

			condition := qbs.NewCondition("pid=?", cid).And("ctype=?", ctype)
			q.Condition(condition).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("topic_count").OrderByDesc("created").FindAll(&alln)

		} else {
			if cid == 0 {
				q.Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("topic_count").OrderByDesc("created").FindAll(&alln)
			} else {
				q.WhereEqual("pid", cid).Offset(offset).Limit(limit).OrderByDesc(path).OrderByDesc("views").OrderByDesc("topic_count").OrderByDesc("created").FindAll(&alln)
			}
		}

	}
	return alln
}

func GetNode(id int64) (node Node) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&node)
	return node
}

func UpdateNode(nid int64, nd Node) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.WhereEqual("id", int64(nid)).Update(&nd)

	return err
}

func EditNode(nid int64, cid int64, uid int64, title string, content string) error {
	nd := GetNode(nid)
	nd.Pid = cid
	nd.Title = title
	nd.Content = content
	nd.Updated = time.Now()
	if err := UpdateNode(nid, nd); err != nil {
		return err
	}

	q, _ := ConnDb()
	defer q.Close()

	type Category struct {
		NodeTime       time.Time
		NodeCount      int64
		NodeLastUserId int64
	}

	if _, err := q.WhereEqual("id", cid).Update(&Category{NodeTime: time.Now(), NodeCount: int64(len(GetAllNodeByCid(cid, 0, 0, 0, "id"))), NodeLastUserId: int64(uid)}); err != nil {
		return err
	}

	return nil
}
