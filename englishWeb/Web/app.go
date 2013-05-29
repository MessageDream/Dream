package main

import (
	"../dataAccess"
	"./handlers/reception"
	"./handlers/root"
	"github.com/astaxie/beego"
)

func main() {
	dataAccess.CreateDb()
	beego.SetStaticPath("/static", "./static")
	beego.SetStaticPath("/archives", "./archives")

	beego.Router("/", &reception.MainHandler{})
	beego.Router("/category/:cid([0-9]+)", &reception.MainHandler{})
	beego.Router("/search", &reception.SearchHandler{})

	beego.Router("/node/:nid([0-9]+)", &reception.NodeHandler{})
	beego.Router("/view/:tid([0-9]+)", &reception.ViewHandler{})

	beego.Router("/register", &reception.RegHandler{})
	beego.Router("/login", &reception.LoginHandler{})
	beego.Router("/logout", &reception.LogoutHandler{})

	beego.Router("/like/topic/:tid([0-9]+)", &reception.LikeTopicHandler{})
	beego.Router("/hate/topic/:tid([0-9]+)", &reception.HateTopicHandler{})

	beego.Router("/like/node/:nid([0-9]+)", &reception.LikeNodeHandler{})
	beego.Router("/hate/node/:nid([0-9]+)", &reception.HateNodeHandler{})

	beego.Router("/new/category", &reception.NewCategoryHandler{})
	beego.Router("/new/node", &reception.NewNodeHandler{})
	beego.Router("/new/topic", &reception.NewTopicHandler{})
	beego.Router("/new/reply/:tid([0-9]+)", &reception.NewReplyHandler{})

	beego.Router("/modify/category", &reception.ModifyCategoryHandler{})
	beego.Router("/modify/node", &reception.ModifyNodeHandler{})

	beego.Router("/topic/delete/:tid([0-9]+)", &reception.TopicDeleteHandler{})
	beego.Router("/topic/edit/:tid([0-9]+)", &reception.TopicEditHandler{})

	beego.Router("/node/delete/:nid([0-9]+)", &reception.NodeDeleteHandler{})
	beego.Router("/node/edit/:nid([0-9]+)", &reception.NodeEditHandler{})

	beego.Router("/delete/reply/:rid([0-9]+)", &reception.DeleteReplyHandler{})

	//root routes
	beego.Router("/root", &root.RMainHandler{})
	beego.Router("/root-login", &root.RLoginHandler{})
	beego.Router("/root/account", &root.RAccountHandler{})

	beego.SessionOn = true
	beego.Run()
}
