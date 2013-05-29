package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type TopicDeleteHandler struct {
	base.RootAuthHandler
}

func (self *TopicDeleteHandler) Get() {
	tid, _ := self.GetInt(":tid")
	bll.DelTopic(tid)
	self.Ctx.Redirect(302, "/")
}
