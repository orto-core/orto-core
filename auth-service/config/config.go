package config

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("auth-config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.Set("database.dsn", os.ExpandEnv(viper.GetString("database.dsn")))
	viper.Set("mail.host", os.ExpandEnv(viper.GetString("mail.host")))
	viper.Set("mail.username", os.ExpandEnv(viper.GetString("mail.username")))
	viper.Set("mail.password", os.ExpandEnv(viper.GetString("mail.password")))

	return nil
}
