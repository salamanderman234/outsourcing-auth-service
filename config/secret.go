package config

import "github.com/spf13/viper"

func GetAppSecret() string {
	return viper.GetString("APP_SECRET")
}

func GetAppVersion() string {
	return string(viper.GetString("APP_VERSION")[0])
}