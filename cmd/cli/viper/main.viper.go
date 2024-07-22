package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func ViperConfig() error {
	viper := viper.New()
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	viper.SetConfigType("yml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fail to read configuration: %w", err)
	}

	// read server configuration
	fmt.Println("Server Port: ", viper.GetInt("server.port"))
	fmt.Println("Server Port: ", viper.GetString("security.jwt.key"))

	var config *Config
	if err = viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("fail to decode configuration: %w", err)
	}

	for _, val := range config.Databases {
		fmt.Println(val.User)
		fmt.Println(val.Password)
		fmt.Println(val.Host)
	}

	return nil
}

func main() {
	err := ViperConfig()
	if err != nil {
		fmt.Println(err)
	}
}
