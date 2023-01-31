package controllers

import (
	"net/http"

	// "ieam-backend/models"
	repo "ieam-backend/repositories"
	"ieam-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginObject struct {
	Username string `json:"username"`
	Password *string `json:"password"`
	// Email    string `json:"email"`
}

func Auth(router *gin.RouterGroup) {
	router.POST("/login", login)
}

func login(c *gin.Context) {
	var req LoginObject
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}

	user, err := repo.Getuserdetail(req.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "This username does not exist.",
		})
		return
	}
	err = utils.CheckPasswordHash(*user.Password, *req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Invalid Password",
			"user":  user,
		})
		return
	} 
	token := utils.Token(user)
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Login is Successfully",
		"token":  token,

	})
}
