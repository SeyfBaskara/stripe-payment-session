package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	ContentfulSpaceID 		string `mapstructure:"CONTENTFUL_SPACE_ID"`
	ContentfulAccesToken 	string `mapstructure:"CONTENTFUL_ACCESS_TOKEN"`
	EnvironmentID			string `mapstructure:"CONTENTFUL_ENVIRONMENT_ID"`
	ContentTypes			string `mapstructure:"CONTENTFUL_CONTENT_TYPES"`
	ServerPort     			string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}