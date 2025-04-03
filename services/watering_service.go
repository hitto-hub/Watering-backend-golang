package services

import (
	"time"

	"github.com/hitto-hub/PlantaTalk/models"
)

// GetTimestamp は現在のタイムスタンプを "YYYY/MM/DD HH:MM:SS" 形式で返す
func GetTimestamp() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// ExecuteWatering は指定されたアドレスに対して水やり指示を生成
func ExecuteWatering(address int) models.Instruction {
	return models.Instruction{
		Timestamp:   GetTimestamp(),
		Address:     address,
		Instruction: 1, // 1 を給水指示とする
	}
}

// 他のビジネスロジックもここに実装可能
