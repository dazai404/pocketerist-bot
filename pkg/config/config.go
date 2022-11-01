package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBPath            string `mapstructure:"db_file"`
	Messages          Messages
}

type Messages struct {
	Errors    Errors
	Resconses Resconses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvaligURL   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"una le_to_save"`
}

type Resconses struct {
	Start            string `mapstructure:"start"`
	AlreadyAutorized string `mapstructure:"already_authorized"`
	AuthSuccess      string `mapstructure:"auth_success"`
	SavedSuccess     string `mapstructure:"saved_success"`
	UnknownCommand   string `mapstructure:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Resconses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}
	
	if err := ParseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseEnv(cfg *Config) error {
	os.Setenv("TOKEN", "5638775876:AAG-yQTwSJEvt_g9gSZ2kFINFNdhwWowApA")
	os.Setenv("CONSUMER_KEY", "104377-ec35202ccf5af04091cebb1")
	os.Setenv("AUTH_SERVER_URL", "http://localhost/")

	if err := viper.BindEnv("token", ); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key", ); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url", ); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	cfg.PocketConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerURL = viper.GetString("auth_server_url")

	return nil
}