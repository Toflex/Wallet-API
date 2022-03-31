package controllers

import (
	"github.com/Toflex/Wallet-API/repository"
	"github.com/Toflex/Wallet-API/services"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	CreditWalletAccount(*gin.Context)
	GetWalletBalance(*gin.Context)
	DebitWalletAccount(*gin.Context)
	Authenticate(*gin.Context)
	VerifyAuth(*gin.Context)
}

type NewController struct {
	Service    services.Service
	Repo	   repository.Repository
}

func DefaultController(serv services.Service, repo repository.Repository) Controller {
	return &NewController{
		Service:    serv,
		Repo:       repo,
	}
}