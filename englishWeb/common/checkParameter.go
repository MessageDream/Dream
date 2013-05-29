package common

import (
	"../models"
	"fmt"
	"github.com/coocood/jas"
	"strconv"
)

func CheckFigure(ctx *jas.Context, name string) (int64, bool) {
	figure, err := strconv.ParseInt(ctx.FormValue("role"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.Data = models.RestResult{IsValid: false, ErrorMessage: name + " must be figure"}
		return figure, false
	}
	return figure, true
}
