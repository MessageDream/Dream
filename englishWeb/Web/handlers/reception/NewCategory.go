package reception

import (
	bll "../../../businessLogic"
	"../base"
)

type NewCategoryHandler struct {
	base.RootAuthHandler
}

func (self *NewCategoryHandler) Get() {
	self.TplNames = "new_category.html"
	self.Layout = "layout.html"

	self.Render()
}

func (self *NewCategoryHandler) Post() {
	inputs := self.Input()
	t := inputs.Get("title")
	c := inputs.Get("content")
	if t != "" && c != "" {
		bll.AddCategory(t, c)
		//后续修改让它跳转到新创建的分类id去
		self.Ctx.Redirect(302, "/")
	} else {
		self.Ctx.Redirect(302, "/")
	}

}
