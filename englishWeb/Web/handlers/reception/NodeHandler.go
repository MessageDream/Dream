package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
)

type NodeHandler struct {
	base.BaseHandler
}

func (self *NodeHandler) Get() {
	page, _ := self.GetInt("page")
	nid, _ := self.GetInt(":nid")

	nid_handler := bll.GetNode(nid)
	nid_handler.Views = nid_handler.Views + 1
	bll.UpdateNode(nid, nid_handler)

	limit := 25
	rcs := len(bll.GetAllTopicByNid(nid, 0, 0, 0, "hotness"))
	pages, pageout, beginnum, endnum, offset := common.Pages(rcs, int(page), limit)
	self.Data["pagesbar"] = common.Pagesbar("", rcs, pages, pageout, beginnum, endnum, 1)
	self.Data["nodeid"] = nid
	self.Data["topics_hotness"] = bll.GetAllTopicByNid(nid, offset, limit, 0, "hotness")
	self.Data["topics_latest"] = bll.GetAllTopicByNid(nid, offset, limit, 0, "id")

	self.TplNames = "node.html"
	self.Layout = "layout.html"

	if nid != 0 {
		self.Render()
		/*
			if sess_userrole, _ := self.GetSession("userrole").(int64); sess_userrole == -1000 {
				self.Render()
			} else {
				nid_path := strconv.Itoa(int(nid_handler.Pid)) + "/" + strconv.Itoa(int(nid_handler.Id)) + "/"
				nid_name := "index.html"
				rs, _ := self.RenderString()
				common.Writefile("./archives/"+nid_path, nid_name, rs)
				self.Redirect("/archives/"+nid_path+nid_name, 301)
			}*/
	} else {
		self.Redirect("/", 302)
	}

}
