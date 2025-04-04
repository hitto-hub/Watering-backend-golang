// package models - アプリケーションのデータモデルを定義するパッケージです。
package models

import "gorm.io/gorm"

// Address - アドレス情報を表す構造体
type Address struct {
	gorm.Model
	Address int    `gorm:"uniqueIndex;not null" json:"address"`
	Name    string `gorm:"uniqueIndex;not null" json:"name"`
}

// WetnessValue - 湿度値のレコードを表す
type WetnessValue struct {
	gorm.Model
	Timestamp string `gorm:"not null" json:"timestamp"`
	Value     int    `gorm:"not null" json:"value"`
	Address   int    `gorm:"not null;index" json:"address"`
}

// TemperatureValue - 温度値のレコードを表す
type TemperatureValue struct {
	gorm.Model
	Timestamp string `gorm:"not null" json:"timestamp"`
	Value     int    `gorm:"not null" json:"value"`
	Address   int    `gorm:"not null;index" json:"address"`
}

// HumidityValue - 湿度のレコードを表す
type HumidityValue struct {
	gorm.Model
	Timestamp string `gorm:"not null" json:"timestamp"`
	Value     int    `gorm:"not null" json:"value"`
	Address   int    `gorm:"not null;index" json:"address"`
}

// WaterSupply - 給水記録を表す
type WaterSupply struct {
	gorm.Model
	Timestamp string `gorm:"not null" json:"timestamp"`
	Type      int    `gorm:"not null" json:"type"`
	Address   int    `gorm:"not null;index" json:"address"`
}

// WateringRegular - 定期給水スケジュールを表す
type WateringRegular struct {
	gorm.Model
	Timestamp   string `gorm:"not null" json:"timestamp"`
	TimeHour    int    `gorm:"not null" json:"time_hour"`
	TimeMinutes int    `gorm:"not null" json:"time_minutes"`
	Weekday     string `gorm:"not null" json:"weekday"`
	Address     int    `gorm:"not null;index" json:"address"`
}

// Instruction - 水やり指示を表す構造体
type Instruction struct {
	gorm.Model
	Timestamp   string `gorm:"not null" json:"timestamp"`
	Address     int    `gorm:"not null;index" json:"address"`
	Instruction int    `gorm:"not null" json:"instruction"`
}
