package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo/internal/api/routers"
	"todo/internal/pkg/config"
)

func Run(configPath string){

	fmt.Println("In api/run")
	if configPath == "" {
		configPath = "data/config.yml"
	}

	setConfiguration(configPath)
	conf := config.GetConfig()
	web := routers.Setup()

	fmt.Println(conf.Server.Port)
	_ = web.Run(":" + conf.Server.Port)
}

func setConfiguration(configPath string) {
	config.Setup(configPath)
	//db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}