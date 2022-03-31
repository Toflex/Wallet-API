package databases

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//Database ...
type Database struct {
	DB *gorm.DB
	RDB *redis.Client
}


// InitRedisDB ...
func (d *Database) InitRedisDB() {
	d.RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return
}

// InitSQLDB ...
func (d *Database) InitSQLDB() {

	dbName:= os.Getenv("DBName")
	dbHost:= os.Getenv("DBHost")
	dbPort:= os.Getenv("DBPort")
	dbUser:= os.Getenv("DBUser")
	dbPass:= os.Getenv("DBPass")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",dbUser, dbPass, dbHost, dbPort, dbName),
		DefaultStringSize: 256,
		DisableDatetimePrecision: true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {

	}

	d.DB = db
	return


}