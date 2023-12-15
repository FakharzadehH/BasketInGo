package config

import (
	"strings"

	"github.com/spf13/viper"
)

var cfg Config

type Config struct {
	DB        DB     `mapstructure:"db"`
	JWTSecret string `mapstructure:"jwt_secret"`
}

func GetConfig() Config {
	return cfg
}
func GetDB() DB {
	return cfg.DB
}

func Load(configPath string) error {
	v := viper.New()
	v.SetEnvPrefix("BasketInGo")
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return nil
}
