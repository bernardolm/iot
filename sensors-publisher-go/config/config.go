package config

import "github.com/spf13/viper"

func Load() {
	viper.SetConfigFile("config.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
