package main

import (
	dal "./dataAccess"
	. "./services"
	"fmt"
	"github.com/coocood/jas"
	"net/http"
)

func main() {
	dal.CreateDb()
	router := jas.NewRouter(new(Users))
	fmt.Println(router.HandledPaths(true))
	http.Handle(router.BasePath, router)
	http.ListenAndServe(":8090", nil)
}
