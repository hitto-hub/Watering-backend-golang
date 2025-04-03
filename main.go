package main

import (
	"log"
	"net/http"

	"github.com/hitto-hub/PlantaTalk/config"
	"github.com/hitto-hub/PlantaTalk/database"
	"github.com/hitto-hub/PlantaTalk/routes"
	"github.com/hitto-hub/PlantaTalk/scheduler"
)

// main 関数はアプリケーションのエントリーポイントです
func main() {
	// 設定読み込み
	cfg := config.LoadConfig()

	// DB 接続初期化
	database.InitDB()

	// スケジューラ起動
	scheduler.StartScheduler()

	// ルーティング設定
	router := routes.SetupRoutes()

	// APIサーバ起動
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
