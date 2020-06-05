package util

import (
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
}