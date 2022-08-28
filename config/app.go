package config

import "github.com/spf13/viper"

type AppConfig struct {
	DbDriver      string `mapstructure:"DB_DRIVER"`
	DbSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

var SECRET_KEY = []byte("secretkeycommitsmartcbs")

var AppConfiguration *AppConfig

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}

	err = viper.Unmarshal(&AppConfiguration)
	return err
}
