package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type DeleteReplyHandler struct {
	base.RootAuthHandler
}

func (self *DeleteReplyHandler) Get() {
	rid, _ := self.GetInt(":rid")
	bll.DelReply(rid)
	self.Ctx.Redirect(302, "/")
}
