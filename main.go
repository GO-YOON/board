package main

import (
	"github.com/gin-gonic/gin"
	"yoon/board/database"
	"yoon/board/routes"
)

func main() {
	/*----------------------------------------------------------------------------
	-
	-	Default()
	-
	----------------------------------------------------------------------------*/
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/static", "./static")
	/*----------------------------------------------------------------------------
	-
	-	Group
	-
	----------------------------------------------------------------------------*/
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "MY %s", "index")
	})

	admin := r.Group("/admin")
	routes.AddUserRouter( admin )

	database.New()
	/*----------------------------------------------------------------------------
	-
	-	Run
	-
	----------------------------------------------------------------------------*/
	r.Run(":8080")
}

