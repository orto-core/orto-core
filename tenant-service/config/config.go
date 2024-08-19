package config

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("tenant-config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.Set("database.dsn", os.ExpandEnv(viper.GetString("database.dsn")))

	return nil
}
