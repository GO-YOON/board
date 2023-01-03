package routes

import (
	"github.com/gin-gonic/gin"
	"yoon/board/controller"
)

func AddUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")

	user.GET("/", controller.FindAllUsers )



}

