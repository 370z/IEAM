package main

import (
	"log"

	"ieam-backend/config"
	ctrl "ieam-backend/controllers"
	"ieam-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.ConnectDB(); err != nil {
		log.Panic("Cannot connect database ", err)
	}
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())

	ctrl.Auth(router.Group("/auth"))
	router.Use(middlewares.Auth())
	ctrl.Dashbord(router.Group("/dashbord"))
	ctrl.User(router.Group("/user"))
	ctrl.Form(router.Group("/form"))
	ctrl.Income_expenses(router.Group("/ie"))
	ctrl.Withdraw_borrow_return(router.Group("/wer"))
	ctrl.History(router.Group("/history"))
	ctrl.Approve(router.Group("/approve"))
	ctrl.Report(router.Group("/report"))
	ctrl.Notreturn(router.Group("/notreturn"))
	ctrl.Notreturnbill(router.Group("/notreturnbill"))

	router.Run()
}
