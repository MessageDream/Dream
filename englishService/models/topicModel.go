package models

import (
	"time"
)

//topic,Pid:node
type Topic struct {
	Id              int64
	Cid             int64 `qbs:"index"`
	Nid             int64 `qbs:"index"`
	Uid             int64 `qbs:"index"`
	Ctype           int64
	Title           string
	Content         string
	Attachment      string
	Created         time.Time `qbs:"index"`
	Updated         time.Time `qbs:"index"`
	Hotness         float64   `qbs:"index"`
	Hotup           int64     `qbs:"index"`
	Hotdown         int64     `qbs:"index"`
	Views           int64     `qbs:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}
