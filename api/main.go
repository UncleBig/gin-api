package main

import (
	"fmt"

	"github.com/UncleBig/gin-api/log"
	"github.com/UncleBig/gin-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//init logger
	log.SetLogger()
	//init router
	router := routers.Init()
	fmt.Println("Init server ....")
	router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
