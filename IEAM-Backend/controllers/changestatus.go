package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)


func Approve(router *gin.RouterGroup) {
	router.PATCH("/approve", Change_Approve)
	router.PATCH("/transfermomey", Change_Transfermomey)
	router.PATCH("/addbill", Change_Addbill)
}

func Change_Approve(c *gin.Context) {
	var err error
	var d repo.ApproveRequest
	err = c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	err = repo.Approve(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Change Status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Status is Successfully Change.",
	})
}

func Change_Transfermomey(c *gin.Context) {
	var err error
	var d repo.TransfermomeyRequest
	err = c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	err = repo.Transfermomey(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Change Status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Status is Successfully Change.",
	})
}
func Change_Addbill(c *gin.Context) {
	var err error
	var d repo.AddbillRequest
	err = c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	err = repo.Addbill(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Change Status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Status is Successfully Change.",
	})
}