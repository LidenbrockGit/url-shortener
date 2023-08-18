package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	if !viper.IsSet("jwt_key") {
		log.Fatalln("Param 'jwt_key' doesn't set in config file")
	}
}
