package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func Counts() (result RestResult) {
	categorys, nodes, topics, menbers := dal.Counts()
	data := [4]int{categorys, nodes, topics, menbers}
	result.Wrap(nil, data)
	return
}
