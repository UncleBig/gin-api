package main

import (
	"fmt"
	"net/http"

	"ginServer/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	//init router
	router := routers.Init()
	fmt.Println("Init server ....")
	http.ListenAndServe(":8080", router)

}
