package models

import (
	"log"
)

type RestResult struct {
	IsValid      bool
	DataItem     interface{}
	ErrorMessage string
}

func (r *RestResult) Wrap(err error, data interface{}) {
	if err != nil {
		r.IsValid = false
		r.ErrorMessage = "server error"
		log.Fatal(err)
	} else {
		r.IsValid = true
		r.DataItem = data
	}
}
