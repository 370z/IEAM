package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)

func History(router *gin.RouterGroup) {
	router.GET("/user", Gethistoryuser)
	router.GET("/approve", Gethistoryapprove)
	router.GET("/finance", Gethistoryfinance)
}


func Gethistoryuser(c *gin.Context)  {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.Gethistoryuser(r)
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

func Gethistoryapprove(c *gin.Context)  {
	
	res, err := repo.Gethistoryapprove()
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

func Gethistoryfinance(c *gin.Context)  {
	
	res, err := repo.Gethistoryfinance()
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

