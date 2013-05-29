package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type NewNodeHandler struct {
	base.AuthHandler
}

func (self *NewNodeHandler) Get() {
	self.TplNames = "new_node.html"
	self.Layout = "layout.html"

	self.Data["categorys"] = bll.GetAllCategory()
	self.Render()
}

func (self *NewNodeHandler) Post() {
	cid, _ := self.GetInt("category")
	uid, _ := self.GetSession("userid").(int64)
	nid_title := self.GetString("title")
	nid_content := self.GetString("content")

	if nid_title != "" && nid_content != "" && cid != 0 {
		bll.AddNode(nid_title, nid_content, cid, uid)
		self.Ctx.Redirect(302, "/category/"+self.GetString("category"))
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
