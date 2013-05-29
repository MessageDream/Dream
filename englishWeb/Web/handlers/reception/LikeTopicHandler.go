package reception

import (
	bll "../../../businessLogic"
	"../../common"
	"../base"
)

type LikeTopicHandler struct {
	base.BaseHandler
}

func (self *LikeTopicHandler) Get() {
	//inputs := self.Input()
	//id, _ := strconv.Atoi(inputs.Get("id"))
	if common.IsSpider(self.Ctx.Request.UserAgent()) != true {

		id, _ := self.GetInt(":tid")

		tp := bll.GetTopic(id)
		tp.Hotup = tp.Hotup + 1
		tp.Hotness = common.Hotness(tp.Hotup, tp.Hotdown, tp.Created)

		bll.SaveTopic(tp)

		self.Ctx.WriteString("success")

	} else {
		self.Ctx.WriteString("R u spider?")
	}

}
