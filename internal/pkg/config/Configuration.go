package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	//Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port   string
	Mode   string
}

func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration

	fmt.Println(configuration)
}

func GetConfig() *Configuration {
	return Config
}