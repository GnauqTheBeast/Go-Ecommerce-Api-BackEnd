package initialize

import (
	"fmt"

	"github.com/GnauqTheBeast/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fail to read configuration: %w", err)
	}

	if err = viper.Unmarshal(&global.Config); err != nil {
		return fmt.Errorf("fail to decode configuration: %w", err)
	}

	return nil
}
