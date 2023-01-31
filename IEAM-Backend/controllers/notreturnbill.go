package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)

func Notreturnbill(router *gin.RouterGroup) {
	router.GET("/Getnotreturnbill", Getnotreturnbill)
}


func Getnotreturnbill(c *gin.Context)  {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.Getnotreturnbill(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting History",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}

