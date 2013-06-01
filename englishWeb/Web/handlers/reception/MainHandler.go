package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
)

type MainHandler struct {
	base.BaseHandler
}

func (self *MainHandler) Get() {
	cid, _ := self.GetInt(":cid")
	page, _ := self.GetInt("page")
	curtab, _ := self.GetInt("tab")
	limit := 20
	home := "false"
	if cid == 0 {
		home = "true"
		self.Data["topics_read"] = bll.GetAllTopicByCid(1, 0, 3, 0, "id")
		self.Data["topics_micro"] = bll.GetAllTopicByCid(2, 0, 3, 0, "id")
		self.Data["topics_joke"] = bll.GetAllTopicByCid(3, 0, 3, 0, "id")
		self.Data["topics_song"] = bll.GetAllTopicByCid(4, 0, 3, 0, "id")
	} else {
		self.Data["curtab"] = curtab
		topics_rcs := len(bll.GetAllTopicByCid(cid, 0, 0, 0, "hotness"))
		topics_pages, topics_pageout, topics_beginnum, topics_endnum, offset := common.Pages(topics_rcs, int(page), limit)
		switch cid {
		case 1:
			self.Data["topics_read"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "id")
			self.Data["topics_pagesbar_tab1"] = common.Pagesbar("tab=1&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
		case 2:
			self.Data["topics_micro"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "id")
		case 3:
			self.Data["topics_joke"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "id")
		case 4:
			self.Data["topics_song"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "id")
		}
	}

	self.Data["home"] = home
	self.Data["curcate"] = cid

	self.Layout = "layout.html"
	self.TplNames = "index.html"

}
