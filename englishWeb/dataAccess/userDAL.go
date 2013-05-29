package dataAccess

import (
	. "../models"
	"time"
)

func AddUser(email string, nickname string, realname string, password string, role int64) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&User{Email: email, Nickname: nickname, Realname: realname, Password: password, Role: role, Created: time.Now()})

	return err
}

func SaveUser(usr User) error {
	q, _ := ConnDb()
	defer q.Close()
	_, e := q.Save(&usr)
	return e
}

func UpdateUser(uid int, ur User) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.WhereEqual("id", int64(uid)).Update(&ur)
	return err
}

func GetUser(id int64) (user User) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("id=?", id).Find(&user)
	return user
}

func DelUser(uid int64) error {
	q, _ := ConnDb()
	defer q.Close()
	usr := GetUser(uid)
	_, err := q.Delete(&usr)

	return err
}

func GetUserByRole(role int) (user User) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("role=?", int64(role)).Find(&user)
	return user
}

func GetAllUserByRole(role int) (user []*User) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("role=?", int64(role)).OrderByDesc("id").FindAll(&user)
	return user
}

func GetUserByNickname(nickname string) (user User) {
	q, _ := ConnDb()
	defer q.Close()
	q.Where("nickname=?", nickname).Find(&user)
	return user
}
