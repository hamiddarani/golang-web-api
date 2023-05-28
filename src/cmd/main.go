package main

import (
	"golang-web-api/api"
	"golang-web-api/config"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
}
