package dataAccess

import (
	. "../models"
)

func AddKV(k string, v string) error {
	q, _ := ConnDb()
	defer q.Close()
	_, err := q.Save(&Kvs{K: k, V: v})
	return err
}

func SetKV(k string, v string) error {
	q, _ := ConnDb()
	defer q.Close()
	var kvs Kvs
	if q.Where("k=?", k).Find(&kvs); kvs.Id == 0 {
		_, err := q.Save(&Kvs{K: k, V: v})
		return err
	} else {
		type Kvs struct {
			K string
			V string
		}

		_, err := q.WhereEqual("k", k).Update(&Kvs{K: k, V: v})

		return err
	}
	return nil
}

func GetKV(k string) (v string) {
	q, _ := ConnDb()
	defer q.Close()
	var kvs Kvs
	q.Where("k=?", k).Find(&kvs)
	return kvs.V
}
