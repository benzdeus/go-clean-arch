package configs

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Mysql   Mysql
	Gin     Gin
	Example Example
}
type Example struct {
	ExampleNumber int
	ExampleString string
}
type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}
type Gin struct {
	Host string
	Port string
}

func LoadConfig(mode string) (configs Configs, err error) {
	viper.AddConfigPath("./..")
	viper.SetConfigName(mode)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return configs, err
	}

	err = viper.Unmarshal(&configs)

	return configs, err
}
