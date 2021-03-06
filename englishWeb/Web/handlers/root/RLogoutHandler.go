package root

import (
	"../base"
)

type RLogoutHandler struct {
	base.BaseHandler
}

func (self *RLogoutHandler) Get() {
	//退出，销毁session
	self.DelSession("userid")
	self.DelSession("username")
	self.DelSession("userrole")
	self.DelSession("useremail")
	self.Ctx.Redirect(302, "/root-login")

}
