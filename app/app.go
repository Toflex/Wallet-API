package app

import (
	"github.com/Toflex/Wallet-API/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// App ...
type App struct {
	Router *gin.Engine
	DB *gorm.DB
	RDB *redis.Client
	Config *configs.Config
}
