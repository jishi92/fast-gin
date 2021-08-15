package main

import (
	"log"
	"net/http"
	"time"

	"fast-gin/config"
	"fast-gin/dao"
	"fast-gin/router"
)

var sc = config.Cfg.Server

func main() {
	dao.Setup()
	defer dao.Close()

	server := &http.Server{
		Addr:           sc.Addr,
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(sc.ReadTimeout * int(time.Second)), // 转换成时间数据结构
		WriteTimeout:   time.Duration(sc.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
