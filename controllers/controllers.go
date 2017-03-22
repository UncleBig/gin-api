package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {

	fmt.Println("Post Handler")
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}

	c.JSON(http.StatusOK, holder)
	return
}
