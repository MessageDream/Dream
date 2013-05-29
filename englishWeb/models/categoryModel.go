package models

import (
	"time"
)

//category,Pid:root
type Category struct {
	Id             int64
	Pid            int64 `qbs:"index"`
	Uid            int64 `qbs:"index"`
	Ctype          int64
	Title          string
	Content        string
	Attachment     string
	Created        time.Time `qbs:"index"`
	Hotness        float64   `qbs:"index"`
	Hotup          int64     `qbs:"index"`
	Hotdown        int64     `qbs:"index"`
	Views          int64     `qbs:"index"`
	Author         string
	NodeTime       time.Time
	NodeCount      int64
	NodeLastUserId int64
}
