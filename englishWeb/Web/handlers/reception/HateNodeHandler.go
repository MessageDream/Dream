package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
)

type HateNodeHandler struct {
	base.BaseHandler
}

func (self *HateNodeHandler) Get() {
	//inputs := self.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if common.IsSpider(self.Ctx.Request.UserAgent()) != true {

		id, _ := self.GetInt(":nid")

		nd := bll.GetNode(id)
		nd.Hotdown = nd.Hotdown + 1
		nd.Hotness = common.Hotness(nd.Hotup, nd.Hotdown, nd.Created)

		bll.SaveNode(nd)

		self.Ctx.WriteString("success")
		//self.Ctx.Redirect(302, "/")

	} else {
		self.Ctx.WriteString("R u spider?")
	}

}
