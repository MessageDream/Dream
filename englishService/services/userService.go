package services

import (
	. "../businessLogic"
	"../common"
	"../models"
	"fmt"
	"github.com/coocood/jas"
	"time"
)

type Users struct{}

func (*Users) PostAdd(ctx *jas.Context) {
	fmt.Println("UserAdd")
	if role, check := common.CheckFigure(ctx, "role"); check {
		ctx.Data = AddUser(ctx.FormValue("email"), ctx.FormValue("nickname"), ctx.FormValue("realname"), ctx.FormValue("password"), role)
	}
}

/*func (*Users) PostSave(ctx *jas.Context) {
	fmt.Println("UserSave")
	ctx.Data =
}*/

func (*Users) PostUpdate(ctx *jas.Context) {
	fmt.Println("UserUpdate")
	var uid int64
	if id, check := common.CheckFigure(ctx, "uid"); check {
		uid = id
	}
	if sex, check1 := common.CheckFigure(ctx, "uid"); check1 {
		birth, _ := time.Parse("2006-01-02 15:04:05", ctx.FormValue("birth"))
		ctx.Data = UpdateUser(int(uid), models.User{
			Email:      ctx.FormValue("email"),
			Nickname:   ctx.FormValue("nickname"),
			Realname:   ctx.FormValue("realname"),
			Avatar:     ctx.FormValue("avatar"),
			Avatar_min: ctx.FormValue("avatar_min"),
			Avatar_max: ctx.FormValue("avatar_max"),
			Birth:      birth,
			Province:   ctx.FormValue("province"),
			City:       ctx.FormValue("city"),
			Company:    ctx.FormValue("company"),
			Address:    ctx.FormValue("address"),
			Postcode:   ctx.FormValue("postcode"),
			Mobile:     ctx.FormValue("mobile"),
			Website:    ctx.FormValue("website"),
			Sex:        sex,
			Qq:         ctx.FormValue("qq"),
			Msn:        ctx.FormValue("msn"),
			Weibo:      ctx.FormValue("weibo")})
	}

}

func (*Users) PostGetuser(ctx *jas.Context) {
	fmt.Println("UserGetUser")
	if id, check := common.CheckFigure(ctx, "uid"); check {
		ctx.Data = GetUser(id)
	}
}

func (*Users) PostDel(ctx *jas.Context) {
	fmt.Println("UserDel")
	if id, check := common.CheckFigure(ctx, "uid"); check {
		ctx.Data = DelUser(id)
	}
}

func (*Users) PostGetuserbyrole(ctx *jas.Context) {
	fmt.Println("UserGetUserByRole")
	if role, check := common.CheckFigure(ctx, "role"); check {
		ctx.Data = GetUserByRole(int(role))
	}

}

func (*Users) PostGetalluserbyrole(ctx *jas.Context) {
	fmt.Println("UserGetAllUserByRole")
	if role, check := common.CheckFigure(ctx, "role"); check {
		ctx.Data = GetAllUserByRole(int(role))
	}

}

func (*Users) PostGetuserbynickname(ctx *jas.Context) {
	fmt.Println("UserGetUserByNickname")
	ctx.Data = GetUserByNickname(ctx.FormValue("nickname"))
}
