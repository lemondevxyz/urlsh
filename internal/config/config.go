package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"github.com/thanhpk/randstr"
)

type Config struct {
	Domain        string `validate:"required" mapstructure:"domain"`
	Sluglength    int    `validate:"required" mapstructure:"sluglength"`
	Characters    string `validate:"required" mapstructure:"characters"`
	SessionSecret string `validate:"required,len=32" mapstructure:"session_secret"`
}

var validate *validator.Validate

func NewConfig() (c Config, err error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetDefault("domain", "http://localhost:8080")
	viper.SetDefault("session_secret", randstr.String(32))
	viper.SetDefault("sluglength", 4)
	viper.SetDefault("characters", "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
