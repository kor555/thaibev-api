package main

import (
	"api/config"
	initializer "api/init"
	"api/router"
)

func main() {
	_, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	init := initializer.Init(config.AppConfig)

	routerFactory := router.NewRouter(init)
	r := routerFactory.CreateRouter()

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
