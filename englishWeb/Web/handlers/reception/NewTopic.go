package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type NewTopicHandler struct {
	base.AuthHandler
}

func (self *NewTopicHandler) Get() {
	self.TplNames = "topic_new.html"
	self.Layout = "layout.html"
	self.Data["nodes"] = bll.GetAllNode()
	self.Render()
}

func (self *NewTopicHandler) Post() {
	nid, _ := self.GetInt("nodeid")
	cid := bll.GetNode(nid).Pid
	uid, _ := self.GetSession("userid").(int64)
	tid_title := self.GetString("title")
	tid_content := self.GetString("content")
	if tid_title != "" && tid_content != "" {
		bll.AddTopic(self.GetString("title"), self.GetString("content"), cid, nid, uid)
		self.Ctx.Redirect(302, "/node/"+self.GetString("nodeid"))
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
