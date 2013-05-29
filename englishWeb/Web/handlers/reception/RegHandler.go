package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
	"fmt"
)

type RegHandler struct {
	base.BaseHandler
}

func (self *RegHandler) Get() {
	self.TplNames = "reg.html"
	self.Render()
}

func (self *RegHandler) Post() {
	self.TplNames = "reg.html"
	self.Ctx.Request.ParseForm()
	username := self.Ctx.Request.Form.Get("username")
	password := self.Ctx.Request.Form.Get("password")
	usererr := common.CheckUsername(username)

	fmt.Println(usererr)
	if usererr == false {
		self.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := common.CheckPassword(password)
	if passerr == false {
		self.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	pwd := common.Encrypt_password(password, nil)

	//now := torgo.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := bll.GetUserByNickname(username)

	if userInfo.Nickname == "" {
		bll.AddUser(username+"@insion.co", username, "", pwd, 1)

		//登录成功设置session
		self.SetSession("userid", userInfo.Id)
		self.SetSession("username", userInfo.Nickname)
		self.SetSession("userrole", userInfo.Role)
		self.SetSession("useremail", userInfo.Email)

		self.Ctx.Redirect(302, "/login")
	} else {
		self.Data["UsernameErr"] = "User already exists"
	}
	self.Render()
}
