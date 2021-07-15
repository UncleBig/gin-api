package routers

import (
	"fmt"

	"github.com/UncleBig/gin-api/controllers"
	"github.com/UncleBig/gin-api/log"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	fmt.Println("init router")
	r := gin.Default()

	//r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log.Logger, true))
	r.POST("/simple/server/post", controllers.PostHandler)
	return r
}
