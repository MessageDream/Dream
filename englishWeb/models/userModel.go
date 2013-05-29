package models

import (
	"time"
)

type User struct {
	Id            int64
	Email         string `qbs:"index"`
	Password      string
	Nickname      string `qbs:"index"`
	Realname      string
	Avatar        string
	Avatar_min    string
	Avatar_max    string
	Birth         time.Time
	Province      string
	City          string
	Company       string
	Address       string
	Postcode      string
	Mobile        string
	Website       string
	Sex           int64
	Qq            string
	Msn           string
	Weibo         string
	Ctype         int64
	Role          int64
	Created       time.Time `qbs:"index"`
	Hotness       float64   `qbs:"index"`
	Hotup         int64     `qbs:"index"`
	Hotdown       int64     `qbs:"index"`
	Views         int64     `qbs:"index"`
	LastLoginTime time.Time
	LastLoginIp   string
	LoginCount    int64
}
