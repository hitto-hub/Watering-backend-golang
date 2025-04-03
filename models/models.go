// package models - アプリケーションのデータモデルを定義するパッケージです。
package models

// Address はアドレス情報を表す構造体
type Address struct {
	Address int    `json:"address"`
	Name    string `json:"name"`
}

// WetnessValue は湿度値のレコードを表す
type WetnessValue struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
	Address   int    `json:"address"`
}

// TemperatureValue は温度値のレコードを表す
type TemperatureValue struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
	Address   int    `json:"address"`
}

// HumidityValue は湿度のレコードを表す
type HumidityValue struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Value     int    `json:"value"`
	Address   int    `json:"address"`
}

// WaterSupply は給水記録を表す
type WaterSupply struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      int    `json:"type"`
	Address   int    `json:"address"`
}

// WateringRegular は定期給水スケジュールを表す
type WateringRegular struct {
	ID          int    `json:"id"`
	Timestamp   string `json:"timestamp"`
	TimeHour    int    `json:"time_hour"`
	TimeMinutes int    `json:"time_minutes"`
	Weekday     string `json:"weekday"`
	Address     int    `json:"address"`
}

// Instruction は水やり指示を表す構造体
// RedeployRequest - 再展開リクエストを表す構造体（例）
type Instruction struct {
	ID          int    `json:"id"`
	Timestamp   string `json:"timestamp"`
	Address     int    `json:"address"`
	Instruction int    `json:"instruction"`
}
