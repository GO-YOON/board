package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yoon/board/database"
	"yoon/board/utils"
	"yoon/board/vo"
)

/*===================================================================
=
= [controller] FindAllUsers
=
===================================================================*/

func FindAllUsers(ctx *gin.Context) {
	users := []vo.User{}
	err := database.DB.Select(&users," SELECT * FROM users ")

	if err != nil {
		log.Fatalln(err)
	}


	pagination := utils.Pagination{
		Page: 1,
		Vertial: 10,
		Total: len(users),
	}

	pages := pagination.Pages()
	fmt.Println("pages",  pages)


	var arr []int

	for i := 1; i<=pages; i++ {
		arr = append(arr, i)
	}

	fmt.Println("arr__", arr)

	ctx.HTML(http.StatusOK, "user.html", gin.H{
		"users": users,
		"pages": arr,
	})
}





