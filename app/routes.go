package app

import (
	"github.com/Toflex/Wallet-API/controllers"
	"github.com/Toflex/Wallet-API/middlewares"
	"github.com/Toflex/Wallet-API/repository"
	"github.com/Toflex/Wallet-API/services"
	"github.com/Toflex/Wallet-API/utility"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"sync"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {

		repo := repository.NewRepository(a.DB, a.RDB)
		serv := services.NewService(repo, a.Config)
		controller := controllers.DefaultController(serv, repo)
		mdd:=middlewares.Middleware{Serv: serv}
		a.Router.Use(gin.Recovery())

		a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


		a.Router.POST("/auth", controller.Authenticate)
		a.Router.POST("/auth/verify", controller.VerifyAuth)


		apiRouter := a.Router.Group("api/v1")

		apiRouter.Use(mdd.LogRequestMiddleware, mdd.AuthMiddleware)

		apiRouter.GET("/wallets/:wallet_id/balance", controller.GetWalletBalance)
		apiRouter.POST("/wallets/:wallet_id/credit", controller.CreditWalletAccount)
		apiRouter.POST("/wallets/:wallet_id/debit", controller.DebitWalletAccount)


		a.Router.NoRoute(func(c *gin.Context) {
			apiResponse:=utility.NewResponse()
			c.JSON(http.StatusNotFound, apiResponse.Error("Page not found"))
		})

	})
	log.Info("Routes have been initialized")
}
