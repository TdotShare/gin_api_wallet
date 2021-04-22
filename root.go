package main

import (
	"root/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routerGroup := r.Group("/api")
	{
		userRouter := routerGroup.Group("/user")
		{

			account := new(controllers.AccountController)
			userRouter.GET("/all", account.ActionUserall)
			userRouter.GET("/view/:id", account.ActionView)
			userRouter.GET("/delete/:id", account.ActionDelete)
			userRouter.POST("/create", account.ActionCreate)
			userRouter.POST("/update", account.ActionUpdate)
		}

		bankRouter := routerGroup.Group("/bank")
		{
			bank := new(controllers.BankController)
			bankRouter.GET("/all", bank.ActionBankGet)
			bankRouter.GET("/view/:id", bank.ActionView)
			bankRouter.GET("/delete/:id", bank.ActionDelete)
			bankRouter.POST("/create", bank.ActionCreate)

		}

		historyRouter := routerGroup.Group("/history")
		{
			history := new(controllers.HistoryController)
			historyRouter.GET("/all", history.ActionHistoryGet)
			historyRouter.GET("/delete/:id", history.ActionDelete)
			historyRouter.POST("/create", history.ActionCreate)

		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
