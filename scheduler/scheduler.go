// package scheduler - アプリケーションのスケジューリング機能を管理するパッケージです。
package scheduler

import (
	"log"
	"time"

	"github.com/hitto-hub/PlantaTalk/database"
	"github.com/hitto-hub/PlantaTalk/services"
)

// StartScheduler は毎分実行されるスケジューラを起動
func StartScheduler() {
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for t := range ticker.C {
			log.Printf("Scheduler triggered at %v", t)
			executeWateringRegular()
		}
	}()
}

// executeWateringRegular は定期水やりスケジュールに基づく処理を実行
func executeWateringRegular() {
	// 実際には DB からスケジュール情報を取得し、条件に合う場合は水やり指示を発行する処理を実装
	// ここでは例として固定のアドレスに対して処理を実行します

	// 例: アドレス 1, 2 に対して水やり実行
	addressesToWater := []int{1, 2}
	for _, addr := range addressesToWater {
		instruction := services.ExecuteWatering(addr)
		// DB への保存処理（省略）
		_ = database.DB
		log.Printf("Executed watering for address %d: %+v", addr, instruction)
	}
}
