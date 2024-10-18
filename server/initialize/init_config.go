package initialize

import (
	"fmt"
	"gincrm/global"
	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 加载配置
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("系统配置文件加载失败 %s", err.Error())
	}

	// 将配置文件读取到全局对象
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("全局配置文件加载失败")
	}
}
