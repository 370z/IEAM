package controllers

import (
	"ieam-backend/models"
	repo "ieam-backend/repositories"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Report(router *gin.RouterGroup) {
	router.POST("/all", Getsearch)
}

func Getsearch(c *gin.Context) {
	var d models.Search
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.Getsearch(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Content List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

