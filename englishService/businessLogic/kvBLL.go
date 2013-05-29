package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddKV(k string, v string) (result RestResult) {
	if k == "" || v == "" {
		result.IsValid = false
		result.ErrorMessage = "neither key nor value can be empty"
		return
	}
	err := dal.AddKV(k, v)
	result.Wrap(err, true)
	return
}

func SetKV(k string, v string) (result RestResult) {
	if k == "" || v == "" {
		result.IsValid = false
		result.ErrorMessage = "neither key nor value can be empty"
		return
	}
	err := dal.SetKV(k, v)
	result.Wrap(err, true)
	return
}

func GetKV(k string) (result RestResult) {
	if k == "" {
		result.IsValid = false
		result.ErrorMessage = " key can not be empty"
		return
	}
	str := dal.GetKV(k)
	result.Wrap(nil, str)
	return
}
