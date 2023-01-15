package main

import (
	"api/configs"
	"api/router"
)

func main() {
	configs.StartConfig()
	router.StartRouter()
}
