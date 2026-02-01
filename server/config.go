package server

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBConn     string `mapstructure:"DB_CONN"`
}

func LoadConfig() *Config {
	// viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var config Config
	bindSystemEnv(config)

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err = viper.ReadInConfig(); err != nil {
			fmt.Printf("error reading config file, using system env %v\n", err)
		}
	} else {
		fmt.Printf("file .env not found, using system env %v\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("error unmarshal config %v\n", err)
	}

	return &config
}

func bindSystemEnv(iface any, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindSystemEnv(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
