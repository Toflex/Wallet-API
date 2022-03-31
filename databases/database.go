package databases

import (
	"fmt"
	"github.com/Toflex/Wallet-API/configs"
	"github.com/Toflex/Wallet-API/models"
	"github.com/Toflex/Wallet-API/repository"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//Database ...
type Database struct {
	DB *gorm.DB
	RDB *redis.Client
	Config *configs.Config
}

// InitRedisDB ...
func (d *Database) InitRedisDB() {

	d.RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", d.Config.RedisHost, d.Config.RedisPort),
		Password: d.Config.RedisPass, // no password set
		DB:       d.Config.RedisDB,  // use default DB
	})

	return
}

// InitSQLDB ...
func (d *Database) InitSQLDB() {

	dbName:= d.Config.DBName
	dbHost:= d.Config.DBHost
	dbPort:= d.Config.DBPort
	dbUser:= d.Config.DBUser
	dbPass:= d.Config.DBPass

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser, dbPass, dbHost, dbPort, dbName),
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})


	if err != nil {
		log.Warning("Failed to connect to SQL DB, %s", err.Error())
	}

	// run migrations on wallets table and pre-populate it with data
	db.AutoMigrate(&models.Wallet{})
	db.Create(&repository.Wallets)

	d.DB = db
	return


}