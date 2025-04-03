package main

import (
	"log"
	"net/http"

	"github.com/hitto-hub/PlantaTalk/internal/config"
	"github.com/hitto-hub/PlantaTalk/internal/database"
	"github.com/hitto-hub/PlantaTalk/internal/routes"
	"github.com/hitto-hub/PlantaTalk/internal/scheduler"
)

func main() {
	// 設定読み込み
	cfg := config.LoadConfig()

	// DB 接続初期化
	database.InitDB()

	// スケジューラ起動
	scheduler.StartScheduler()

	// ルーティング設定
	router := routes.SetupRoutes()

	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
