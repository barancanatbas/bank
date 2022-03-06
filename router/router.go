package router

import (
	"mongoexample/api"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	bankHandler := api.BankInit()

	r.GET("/account/:id", bankHandler.Info)
	r.POST("/account", bankHandler.Create)
	r.PUT("/add/money", bankHandler.AddMoney)
	r.PUT("/reduce/money", bankHandler.ReduceMoney)
}
