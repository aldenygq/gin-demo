package httpclient

import "github.com/spf13/viper"

func Init() {
	cfgUtil := viper.Sub("util")
	if cfgUtil != nil {
		InitHttpClient()
	}
}
