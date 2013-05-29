package reception

import (
	"../base"
)

type LogoutHandler struct {
	base.BaseHandler
}

func (self *LogoutHandler) Get() {
	//退出，销毁session
	self.DelSession("userid")
	self.DelSession("username")
	self.DelSession("userrole")
	self.DelSession("useremail")
	self.Ctx.Redirect(302, "/")

}
