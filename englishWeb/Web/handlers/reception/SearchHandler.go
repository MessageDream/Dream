package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
)

type SearchHandler struct {
	base.BaseHandler
}

func (self *SearchHandler) Get() {
	if keyword := self.GetString("keyword"); keyword != "" {
		page, _ := self.GetInt("page")
		limit := 25

		rcs := len(bll.SearchTopic(keyword, 0, 0, "id"))
		pages, pageout, beginnum, endnum, offset := common.Pages(rcs, int(page), limit)
		self.Data["search_hotness"] = bll.SearchTopic(keyword, offset, limit, "hotness")

		keywordz := "keyword=" + keyword + "&"
		self.Data["pagesbar"] = common.Pagesbar(keywordz, rcs, pages, pageout, beginnum, endnum, 1)

	}
	self.TplNames = "search.html"
	self.Layout = "layout.html"

	self.Render()
}
