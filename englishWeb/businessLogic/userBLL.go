package businessLogic

import (
	dal "../dataAccess"
	. "../models"
	"errors"
)

func AddUser(email string, nickname string, realname string, password string, role int64) error {
	if nickname == "" || password == "" || email == "" {
		return errors.New("nickname and password and email can not be empty")
	}
	return dal.AddUser(email, nickname, realname, password, role)
}

func SaveUser(usr User) error {
	if usr.Nickname == "" || usr.Password == "" || usr.Email == "" {
		return errors.New("nickname and password and email can not be empty")
	}
	return dal.SaveUser(usr)
}

func UpdateUser(uid int, ur User) error {
	if uid < 0 {
		return errors.New("uid can not be less than zero")
	}
	old := dal.GetUser(int64(uid))
	if old.Password == "" {
		return errors.New("model is not exist")
	}
	old.Email = ur.Email
	old.Nickname = ur.Nickname
	old.Realname = ur.Realname
	old.Avatar = ur.Avatar
	old.Avatar_min = ur.Avatar_min
	old.Avatar_max = ur.Avatar_max
	old.Birth = ur.Birth
	old.Province = ur.Province
	old.City = ur.City
	old.Company = ur.Company
	old.Address = ur.Address
	old.Postcode = ur.Postcode
	old.Mobile = ur.Mobile
	old.Website = ur.Website
	old.Sex = ur.Sex
	old.Qq = ur.Qq
	old.Msn = ur.Msn
	old.Weibo = ur.Weibo
	return dal.UpdateUser(uid, old)
}

func GetUser(id int64) (user User) {
	user = dal.GetUser(id)
	return
}

func DelUser(uid int64) error {
	if uid < 0 {
		return errors.New("uid can not be less than zero")
	}
	return dal.DelUser(uid)
}

func GetUserByRole(role int) (user User) {
	user = dal.GetUserByRole(role)
	return
}

func GetAllUserByRole(role int) (user []*User) {
	user = dal.GetAllUserByRole(role)
	return
}

func GetUserByNickname(nickname string) (user User) {
	user = dal.GetUserByNickname(nickname)
	return
}
