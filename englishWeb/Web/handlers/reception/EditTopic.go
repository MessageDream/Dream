package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type TopicEditHandler struct {
	base.RootAuthHandler
}

func (self *TopicEditHandler) Get() {
	tid, _ := self.GetInt(":tid")
	tid_handler := bll.GetTopic(tid)
	self.Data["topic"] = tid_handler
	self.Data["inode"] = bll.GetNode(tid_handler.Nid)

	self.Layout = "layout.html"
	self.TplNames = "topic_edit.html"
	self.Render()
}

func (self *TopicEditHandler) Post() {
	tid, _ := self.GetInt(":tid")
	nid, _ := self.GetInt("nodeid")
	cid := bll.GetNode(nid).Pid
	uid, _ := self.GetSession("userid").(int64)
	tid_title := self.GetString("title")
	tid_content := self.GetString("content")

	if tid_title != "" && tid_content != "" {
		bll.EditTopic(tid, nid, cid, uid, tid_title, tid_content)
		self.Ctx.Redirect(302, "/view/"+self.Ctx.Params[":tid"])
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
