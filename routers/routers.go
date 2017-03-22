package routers

import (
	"fmt"
	"ginServer/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() http.Handler {
	fmt.Println("init router")
	router := gin.Default()
	router.POST("/simple/server/post", controllers.PostHandler)
	return router
}
