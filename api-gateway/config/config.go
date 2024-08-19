package config

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("gateway-config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.Set("services.auth_service.url", os.ExpandEnv(viper.GetString("services.auth_service.url")))
	viper.Set("services.tenant_service.url", os.ExpandEnv(viper.GetString("services.tenant_service.url")))
	viper.Set("services.page_service.url", os.ExpandEnv(viper.GetString("services.page_service.url")))
	viper.Set("authentication.jwt_secret", os.ExpandEnv(viper.GetString("authentication.jwt_secret")))

	return nil
}
