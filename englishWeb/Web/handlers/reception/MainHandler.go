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
	page, _ := self.GetInt("page")
	curtab, _ := self.GetInt("tab")
	cid, _ := self.GetInt(":cid")
	limit := 25
	home := "false"
	if cid == 0 {
		home = "true"
	}

	self.Data["home"] = home
	self.Data["curcate"] = cid
	self.Data["curtab"] = curtab

	topics_rcs := len(bll.GetAllTopicByCid(cid, 0, 0, 0, "hotness"))
	topics_pages, topics_pageout, topics_beginnum, topics_endnum, offset := common.Pages(topics_rcs, int(page), limit)

	self.Data["topics_latest"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "id")
	self.Data["topics_hotness"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "hotness")
	self.Data["topics_views"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "views")
	self.Data["topics_reply_count"] = bll.GetAllTopicByCid(cid, offset, limit, 0, "reply_count")

	self.Data["topics_pagesbar_tab1"] = common.Pagesbar("tab=1&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	self.Data["topics_pagesbar_tab2"] = common.Pagesbar("tab=2&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	self.Data["topics_pagesbar_tab3"] = common.Pagesbar("tab=3&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)
	self.Data["topics_pagesbar_tab4"] = common.Pagesbar("tab=4&", topics_rcs, topics_pages, topics_pageout, topics_beginnum, topics_endnum, 1)

	nodes_rcs := len(bll.GetAllNodeByCid(cid, 0, 0, 0, "hotness"))
	nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, offset := common.Pages(nodes_rcs, int(page), limit)

	self.Data["nodes_latest"] = bll.GetAllNodeByCid(cid, offset, limit, 0, "id")
	self.Data["nodes_hotness"] = bll.GetAllNodeByCid(cid, offset, limit, 0, "hotness")
	self.Data["nodes_views"] = bll.GetAllNodeByCid(cid, offset, limit, 0, "views")

	self.Data["nodes_pagesbar_tab5"] = common.Pagesbar("tab=5&", nodes_rcs, nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, 1)
	self.Data["nodes_pagesbar_tab6"] = common.Pagesbar("tab=6&", nodes_rcs, nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, 1)
	self.Data["nodes_pagesbar_tab7"] = common.Pagesbar("tab=7&", nodes_rcs, nodes_pages, nodes_pageout, nodes_beginnum, nodes_endnum, 1)

	self.Layout = "layout.html"
	self.TplNames = "index.html"
	//self.Render()

}
