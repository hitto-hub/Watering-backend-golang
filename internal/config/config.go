package config

import "os"

// Config は設定情報を表す
type Config struct {
	ServerPort string
	// 他の設定項目
}

// LoadConfig は環境変数などから設定を読み込み Config を返す
func LoadConfig() *Config {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // デフォルトポート
	}
	return &Config{
		ServerPort: port,
	}
}
