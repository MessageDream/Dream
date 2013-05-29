package models

import (
	"time"
)

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64 `qbs:"index"`
	Pid        int64 `qbs:"index"` //Topic id
	Ctype      int64
	Content    string
	Attachment string
	Created    time.Time `qbs:"index"`
	Hotness    float64   `qbs:"index"`
	Hotup      int64     `qbs:"index"`
	Hotdown    int64     `qbs:"index"`
	Views      int64     `qbs:"index"`
	Author     string
	Email      string
	Website    string
}
