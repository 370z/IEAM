package controllers

import (
	"net/http"

	"ieam-backend/models"
	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)

func Form(router *gin.RouterGroup) {
	router.GET("/all", Getallform)
	router.POST("/add", Createform)
	router.POST("/addreturn", Createformreturn)
	router.GET("/detail", FormGetbyid)
	router.PUT("/update", FormUpdate)
	router.DELETE("/delete", FormDelete)
}

func Getallform(c *gin.Context) {
	res, err := repo.Getallform()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Form List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func Createform(c *gin.Context) {
	var d models.Form
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.Createform(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Form",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Form is Successfully Created.",
		"data": res,
	})
}
func Createformreturn(c *gin.Context) {
	var d models.Form
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.Createformreturn(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Form",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Form is Successfully Created.",
		"data": res,
	})
}

func FormGetbyid(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.Getdetailform(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Form",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}

func FormUpdate(c *gin.Context) {
	var d models.Form
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	res, err := repo.Updateform(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Edit Blog",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Blog is Successfully Edite.",
		"data": res,
	})
}

func FormDelete(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}

	err := repo.Deleteform(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Delete Form",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 204,
		"msg":  "Form is Successfully Delete.",
	})
}
