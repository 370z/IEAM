package controllers

import (
	"ieam-backend/models"
	repo "ieam-backend/repositories"
	"ieam-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup)  {
	router.GET("/", Getalluser) 
	router.GET("/detail", Detailuser) 
	router.POST("/create", Createuser) 
	router.PATCH("/edite", Editeuser) 
	router.DELETE("/delete", UserDelete) 
}

func Getalluser(c *gin.Context)  {
	res, err := repo.Getalluser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting User List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func Detailuser(c *gin.Context)  {
	userId := c.MustGet("userId").(float64)
	res, err := repo.GetuserMe(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting User List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func Createuser(c *gin.Context)  {
	var d models.Account
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}
	hash := utils.PasswordHashing(*d.Password)
	d.Password = &hash
	res, err := repo.Createuser(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created User",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "User is Successfully Created.",
		"data": res,
	})
}

func Editeuser(c *gin.Context)  {
	var d models.Account
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}
	hash := utils.PasswordHashing(*d.Password)
	d.Password = &hash
	res, err := repo.Editeuser(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created User",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "User is Successfully Created.",
		"data": res,
	})
}

func UserDelete(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}

	err := repo.DeleteUser(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Delete User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 204,
		"msg":  "User is Successfully Delete.",
	})
}