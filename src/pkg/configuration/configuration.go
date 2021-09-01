package configuration

import (
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Configuration interface {
}

type Settings struct {
	ConsulAddress string
	ConsulKey     string
	LocalPath     string
}

func NewConfigService(settings Settings) (*viper.Viper, error) {
	config := viper.New()
	config.AllowEmptyEnv(true)
	config.AutomaticEnv()

	config.AddRemoteProvider("consul", settings.ConsulAddress, settings.ConsulKey)
	config.SetConfigType("json")
	err := config.ReadRemoteConfig()

	if err != nil {
		fmt.Println(err)
	} else {
		return config, nil
	}

	config.SetConfigName("settings.json")
	config.AddConfigPath(settings.LocalPath)
	err = config.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
	return config, err
}
