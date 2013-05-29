package businessLogic

import (
	dal "../dataAccess"
	. "../models"
)

func AddUser(email string, nickname string, realname string, password string, role int64) (result RestResult) {
	if nickname == "" || password == "" || email == "" {
		result.IsValid = false
		result.ErrorMessage = "nickname and password and email can not be empty"
		return
	}
	err := dal.AddUser(email, nickname, realname, password, role)
	result.Wrap(err, true)
	return

}

func SaveUser(usr User) (result RestResult) {
	if usr.Nickname == "" || usr.Password == "" || usr.Email == "" {
		result.IsValid = false
		result.ErrorMessage = "nickname and password and email can not be empty"
		return
	}
	err := dal.SaveUser(usr)
	result.Wrap(err, true)
	return
}

func UpdateUser(uid int, ur User) (result RestResult) {
	if uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "uid can not be less than zero"
		return
	}
	old := dal.GetUser(int64(uid))
	if old.Password == "" {
		result.IsValid = false
		result.ErrorMessage = "model is not exist"
		return
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
	err := dal.UpdateUser(uid, old)
	result.Wrap(err, true)
	return
}

func GetUser(id int64) (result RestResult) {
	us := dal.GetUser(id)
	result.Wrap(nil, us)
	return
}

func DelUser(uid int64) (result RestResult) {
	if uid < 0 {
		result.IsValid = false
		result.ErrorMessage = "uid can not be less than zero"
		return
	}
	err := dal.DelUser(uid)
	result.Wrap(err, true)
	return
}

func GetUserByRole(role int) (result RestResult) {
	us := dal.GetUserByRole(role)
	result.Wrap(nil, us)
	return
}

func GetAllUserByRole(role int) (result RestResult) {
	user := dal.GetAllUserByRole(role)
	result.Wrap(nil, user)
	return
}

func GetUserByNickname(nickname string) (result RestResult) {
	user := dal.GetUserByNickname(nickname)
	result.Wrap(nil, user)
	return
}
