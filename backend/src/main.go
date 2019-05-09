package main

import (
	"log"
	"net/http"
	"router"
	"utils"
)

func main() {

	utils.Connect()
	router.Build()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
