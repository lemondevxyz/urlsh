package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"github.com/thanhpk/randstr"
	"github.com/toms1441/urlsh/internal/shortener"
)

type Config struct {
	Domain        string           `validate:"required" mapstructure:"domain"`
	SessionSecret string           `validate:"required,len=32" mapstructure:"session_secret"`
	Shortener     shortener.Config `validate:"required" mapstructure:"shortener"`
}

var validate *validator.Validate

func NewConfig() (c Config, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("domain", "http://localhost:8080")
	viper.SetDefault("session_secret", randstr.String(32))
	viper.SetDefault("shortener", shortener.Config{
		Length:     4,
		Characters: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	})

	// Initiate viper for our config
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			os.Create("config.yaml")
			viper.WriteConfig()
		}
	}

	// Unmarshal the config
	err = viper.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("viper.Unmarshal: %v", err)
	}

	if validate == nil {
		validate = validator.New()
	}

	err = validate.Struct(&c)
	if err != nil {
		return c, fmt.Errorf("validate.Struct: %v", err)
	}

	return
}
