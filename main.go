package main

import (
	"main/controllers"

	"main/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @termsOfService http://swagger.io/terms/

func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "GIN WALLET API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
