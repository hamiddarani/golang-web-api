package main

import (
	"golang-web-api/api"
	"golang-web-api/config"
	"golang-web-api/data/cache"
	"log"
)

func main() {
	cfg := config.GetConfig()
	
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	api.InitServer(cfg)
}
