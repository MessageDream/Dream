package dataAccess

import (
	"../models"
	"fmt"
	"github.com/coocood/qbs"
	//_ "github.com/mattn/go-sqlite3"
	//_ "github.com/coocood/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DbName         = "english"
	DbUser         = "root"
	mysqlDriver    = "mymysql"
	mysqlDrvformat = "%v/%v/"
	pgDriver       = "postgres"
	pgDrvFormat    = "user=%v dbname=%v sslmode=disable"
	sqlite3Driver  = "sqlite3"
	dbtypeset      = "mysql"
)

func RegisterDb() {

	switch {
	case dbtypeset == "sqlite":
		qbs.Register("sqlite3", "../data/sqlite.db", "", qbs.NewSqlite3())

	case dbtypeset == "mysql":
		qbs.Register("mysql", "root:@/english?charset=utf8&parseTime=true&loc=Local", DbName, qbs.NewMysql())

	case dbtypeset == "pgsql":
		qbs.Register("postgres", "qbs_test@/qbs_test?charset=utf8&parseTime=true&loc=Local", DbName, qbs.NewPostgres())
	}

}

func ConnDb() (q *qbs.Qbs, err error) {
	q, err = qbs.GetQbs()
	return q, err
}

func SetMg() (mg *qbs.Migration, err error) {
	mg, err = qbs.GetMigration()
	return mg, err
}

func CreateDb() bool {
	RegisterDb()
	q, err := ConnDb()
	defer q.Close()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		mg, _ := SetMg()
		defer mg.Close()

		mg.CreateTableIfNotExists(new(models.User))
		mg.CreateTableIfNotExists(new(models.Category))
		mg.CreateTableIfNotExists(new(models.Node))
		mg.CreateTableIfNotExists(new(models.Topic))
		mg.CreateTableIfNotExists(new(models.Reply))
		mg.CreateTableIfNotExists(new(models.Kvs))
		mg.CreateTableIfNotExists(new(models.File))

		//用户等级划分：正数是普通用户，负数是管理员各种等级划分，为0则尚未注册
		if GetUserByRole(-1000).Role != -1000 {
			AddUser("root@localhost", "root", "系统默认管理员", "rootpass", -1000)
			fmt.Println("Default User:root,Password:rootpass")

			if GetAllTopic(0, 0, "id") == nil {
				//分類默認數據
				AddCategory("English！", "This is EnglishCategory！")

				AddNode("Node！", "This is Node!", 1, 1)
				SetTopic(0, 1, 1, 1, 0, "Topic Title", `<p>This is Topic!</p>`, "root", "")

			}
		}

		/*if GetKV("author") != "Insion" {
			SetKV("author", "Insion")
			SetKV("title", "Toropress")
			SetKV("title_en", "Toropress")
			SetKV("keywords", "Toropress,")
			SetKV("description", "Toropress,")

			SetKV("company", "Toropress")
			SetKV("copyright", "2013 Copyright Toropress .All Right Reserved")
			SetKV("site_email", "info@verywave.com")

			SetKV("tweibo", "http://t.qq.com/yours")
			SetKV("sweibo", "http://weibo.com/yours")
		}*/

		return true
	}
	return false
}
