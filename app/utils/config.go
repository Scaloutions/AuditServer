package utils

import "github.com/spf13/viper"

func SetUpExternalConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigType("toml")
	v.SetConfigName("app")
	v.AddConfigPath("app/config")
	err := v.ReadInConfig()
	CheckAndHandleError(err)
	v.SetConfigName("messages")
	err = v.MergeInConfig()
	CheckAndHandleError(err)
	return v
}
