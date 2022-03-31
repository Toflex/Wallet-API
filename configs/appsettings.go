package configs

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Config struct {
	DBName string
	DBHost string
	DBPort string
	DBUser string
	DBPass string

	RedisHost string
	RedisPort string
	RedisPass string
	RedisDB   int

	ServerPort string
	ServerHost string

	Secret string
}

func Configuration() *Config {
	redisDB, err := strconv.Atoi(os.Getenv("RedisDB"))
	if err != nil {
		log.Errorf("Redis DB expects integer value, DB has been set to '0'")
		redisDB = 0
	}

	config := Config{
		DBName:     os.Getenv("DBName"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPass:     os.Getenv("DBPass"),
		RedisHost:  os.Getenv("RedisHost"),
		RedisPort:  os.Getenv("RedisPort"),
		RedisPass:  os.Getenv("RedisPass"),
		RedisDB:    redisDB,
		ServerPort: os.Getenv("ServerPort"),
		ServerHost: os.Getenv("ServerHost"),
		Secret: os.Getenv("Secret"),
	}

	return &config
}