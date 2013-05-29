package models

import (
	"time"
)

//node,Pid:category
type Node struct {
	Id              int64
	Pid             int64 `qbs:"index"`
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
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}
