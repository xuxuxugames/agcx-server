package main

import (
	"fmt"
	"github.com/agcx_server/config"
	"github.com/agcx_server/routers"
	"github.com/agcx_server/utils/database"
	"log"
	"net/http"
)

func main() {
	// 初始化数据库和数据表
	database.Connect()
	database.Migrate()
	database.Seed()

	// 初始化路由
	router := routers.Router()

	// 初始化 HTTP 服务器
	server := &http.Server{

		Addr:           fmt.Sprintf(":%d", config.Server.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动 HTTP 服务器
	log.Printf("Server listening at port: %d", config.Server.HTTPPort)
	server.ListenAndServe()
}
