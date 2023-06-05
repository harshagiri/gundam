package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
	}
}

func LoadConfig() (Config, error) {
	viper.SetDefault("Server.Port", 8080)
	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", 3306)
	viper.SetDefault("Database.User", "root")
	viper.SetDefault("Database.Password", "")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // You can specify additional configuration paths here if needed.

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config file: %s", err))
	}

	return config, nil
}

func (c Config) DatabaseConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/database_name", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port)
}

func (c Config) GetServerPort() int {
	return c.Server.Port
}
