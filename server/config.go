package server

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBConn     string `mapstructure:"DB_CONN"`
}

func LoadConfig() *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	// config.DBConn = viper.GetString("DB_CONN")
	// config.ServerPort = viper.GetString("SERVER_PORT")

	return &config
}
