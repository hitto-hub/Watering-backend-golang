// package database - GORMを使ったデータベース接続・初期化を行うパッケージです。
package database

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hitto-hub/PlantaTalk/models"
)

var DB *gorm.DB

// InitDB - GORMを使ってDB初期化
func InitDB() {
	var err error

	// SQLiteに接続（なければファイルが作成される）
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続に失敗しました: %v", err)
	}

	fmt.Println("GORM: SQLiteに接続成功")

	// モデルに基づいてテーブルを自動作成（マイグレーション）
	err = DB.AutoMigrate(
		&models.Address{},
		&models.WetnessValue{},
		&models.TemperatureValue{},
		&models.HumidityValue{},
		&models.WaterSupply{},
		&models.WateringRegular{},
		&models.Instruction{},
	)
	if err != nil {
		log.Fatalf("テーブル作成に失敗しました: %v", err)
	}

	fmt.Println("GORM: テーブル自動作成完了")
}
