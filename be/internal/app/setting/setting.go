package setting

import (
	"github.com/spf13/viper"
	"strings"
)

var AppConfig Configuration

type Configuration struct {
	Database database `mapstructure:"database"`
}

type database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

func InitConfiguration() {
	ymlConfig := viper.New()
	ymlConfig.AddConfigPath("./configs")
	ymlConfig.SetConfigName("app")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	ymlConfig.SetEnvKeyReplacer(replacer)
	ymlConfig.AutomaticEnv()

	err := ymlConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = ymlConfig.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
