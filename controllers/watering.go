// package controllers - アプリケーションのコントローラを定義するパッケージです。
// ここでは、定期水やりスケジュールを設定するハンドラを定義しています。
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// SetWateringRegularHandler - 定期水やりスケジュールを設定するハンドラ
// @Summary 定期水やりスケジュール設定
// @Description 指定アドレスの定期水やりスケジュールを設定
// @Param address query int true "Address"
// @Param schedule body struct {
//   TimeHour    int    `json:"time_hour"`
//   TimeMinutes int    `json:"time_minutes"`
//   Weekday     string `json:"weekday"`
// } true "Watering schedule"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/watering_regular [post]
func SetWateringRegularHandler(w http.ResponseWriter, r *http.Request) {
	addressStr := r.URL.Query().Get("address")
	address, err := strconv.Atoi(addressStr)
	if err != nil {
		http.Error(w, "Invalid address", http.StatusBadRequest)
		return
	}

	var req struct {
		TimeHour    int    `json:"time_hour"`
		TimeMinutes int    `json:"time_minutes"`
		Weekday     string `json:"weekday"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 入力値のバリデーションなど処理（省略）

	resp := map[string]interface{}{
		"status":       true,
		"message":      "Watering schedule set successfully",
		"address":      address,
		"time_hour":    req.TimeHour,
		"time_minutes": req.TimeMinutes,
		"weekday":      req.Weekday,
	}
	json.NewEncoder(w).Encode(resp)
}
