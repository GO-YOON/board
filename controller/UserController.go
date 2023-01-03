package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yoon/board/database"
	"yoon/board/vo"
)

/*===================================================================
=
= [service] FindAllUsers
=
===================================================================*/

func FindAllUsers(ctx *gin.Context) {

	users := []vo.User{}

	err := database.DB.Select(&users," SELECT * FROM users ")

	if err != nil {
		log.Fatalln(err)
	}

	for i, user := range users {

		fmt.Println(i, user.Email)

	}

	data := users

	ctx.HTML(http.StatusOK, "user.html", data)
}







