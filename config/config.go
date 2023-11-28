package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Address string
	Token   string
}

func Init() Config {
	// *viper.Viper 초기화
	viperConfig := viper.New()
	// 설정 파일의 디렉토리 세팅
	viperConfig.AddConfigPath(".")
	// 설정 파일명 세팅
	viperConfig.SetConfigFile("app.env")
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Println("Error on Reading Viper Config")
		panic(err)
	}

	var config Config
	// 읽어온 설정값을 config 로 언마샬
	err = viperConfig.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error on Unmarshal Viper Config")
		panic(err)
	}
	return config
}
