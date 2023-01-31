package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)

func Income_expenses(router *gin.RouterGroup) {
	router.GET("/all", GetallIE)

}

func GetallIE(c *gin.Context)  {
	res, err := repo.GetallformIe()
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