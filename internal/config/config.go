package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config は設定情報を表す
type Config struct {
	ServerPort string
	// 他の設定項目
}

// LoadConfig は環境変数などから設定を読み込み Config を返す
func LoadConfig() *Config {
	// env := os.Getenv("APP_ENV")
	// viper.SetConfigName("config." + env) // config.production.yaml などを読み込む
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &Config{
		ServerPort: viper.GetString("server.port"),
	}
}
