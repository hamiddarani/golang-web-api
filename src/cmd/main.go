package main

import (
	"golang-web-api/api"
	"golang-web-api/config"
	"golang-web-api/data/cache"
	"golang-web-api/data/db"
	"log"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	api.InitServer(cfg)
}
