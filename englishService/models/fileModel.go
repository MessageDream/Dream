package models

import (
	"time"
)

type File struct {
	Id              int64
	Cid             int64 `qbs:"index"`
	Nid             int64 `qbs:"index"`
	Uid             int64 `qbs:"index"`
	Pid             int64 `qbs:"index"`
	Ctype           int64
	Filename        string
	Content         string
	Hash            string
	Location        string
	Url             string
	Size            int64
	Created         time.Time `qbs:"index"`
	Updated         time.Time `qbs:"index"`
	Hotness         float64   `qbs:"index"`
	Hotup           int64     `qbs:"index"`
	Hotdown         int64     `qbs:"index"`
	Views           int64     `qbs:"index"`
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}
