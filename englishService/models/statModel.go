package models

import (
	"time"
)

type Stat struct {
	Ip      string
	Ua      string
	Created time.Time
}
