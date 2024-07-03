package util

import (
	"time"
	
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	HTTPServerAddress   string        `mapstructure:"SWD392_HTTP_SERVER_ADDRESS"`
	DBSource            string        `mapstructure:"SWD392_DB_SOURCE"`
	TokenSymmetricKey   string        `mapstructure:"SWD392_TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"SWD392_ACCESS_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	
	viper.SetEnvPrefix("SWD392")
	viper.AutomaticEnv()
	
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	
	err = viper.Unmarshal(&config)
	return
}
