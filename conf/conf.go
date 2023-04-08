package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load config error:%s", err.Error()))
	}
}
