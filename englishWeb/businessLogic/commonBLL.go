package businessLogic

import (
	dal "../dataAccess"
)

func Counts() (categorys int, nodes int, topics int, menbers int) {
	categorys, nodes, topics, menbers = dal.Counts()
	return
}
