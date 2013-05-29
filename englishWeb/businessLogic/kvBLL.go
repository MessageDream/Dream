package businessLogic

import (
	dal "../dataAccess"
	"errors"
)

func AddKV(k string, v string) error {
	if k == "" || v == "" {
		return errors.New("neither key nor value can be empty")
	}
	return dal.AddKV(k, v)
}

func SetKV(k string, v string) error {
	if k == "" || v == "" {
		return errors.New("neither key nor value can be empty")
	}
	return dal.SetKV(k, v)
}

func GetKV(k string) string {
	return dal.GetKV(k)
}
