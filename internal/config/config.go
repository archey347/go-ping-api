package config

import "github.com/spf13/viper"

type Config struct {
	ApiServer ApiServerConfig `mapstructure:"apiServer"`
}

type ApiServerConfig struct {
	Address string `mapstructure:"address" default:""`
	Port    int    `mapstructure:"port"`
}

func Load(configFile string) (config Config, err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
