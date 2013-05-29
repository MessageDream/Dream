package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type NodeDeleteHandler struct {
	base.RootAuthHandler
}

func (self *NodeDeleteHandler) Get() {
	nid, _ := self.GetInt(":nid")
	bll.DelNode(nid)
	self.Ctx.Redirect(302, "/")
}
