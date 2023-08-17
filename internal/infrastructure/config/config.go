package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("../../..")
	viper.SetConfigFile("config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
