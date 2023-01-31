package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)

func Withdraw_borrow_return(router *gin.RouterGroup) {
	router.GET("/all", GetallWBR)
	router.GET("/approve", GetapproveWBR)

}

func GetallWBR(c *gin.Context)  {
	res, err := repo.GetallformWer()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Approve List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}
func GetapproveWBR(c *gin.Context)  {
	res, err := repo.GetapproveWBR()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Approve List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}