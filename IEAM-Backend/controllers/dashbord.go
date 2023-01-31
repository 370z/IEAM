package controllers

import (
	"net/http"

	repo "ieam-backend/repositories"

	"github.com/gin-gonic/gin"
)


func Dashbord(router *gin.RouterGroup) {
	router.GET("/dashbord", Dashborddata)
	router.GET("/dashbord_all", Dashborddata_all)
	router.GET("/dashbord_historynotreturn", Dashborddata_historynotreturn)
}

func Dashborddata(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.Getdashbord(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Content",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}

func Dashborddata_all(c *gin.Context) {
	res, err := repo.Getdashbord_all()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Content",
		})
		return
	}

	res_numberform, err := repo.Getdashbord_numberform()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Content",
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
		"data_numberform" :res_numberform,
	})
}
func Dashborddata_historynotreturn(c *gin.Context) {
	res, err := repo.Dashborddata_historynotreturn()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Content",
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}
