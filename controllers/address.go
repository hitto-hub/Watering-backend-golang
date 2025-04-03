// Package controllers - アプリケーションのコントローラを定義するパッケージです。
// ここでは、アドレス情報を取得するハンドラを定義しています。
package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hitto-hub/PlantaTalk/repository"
)

// GetAddressesHandler は全アドレス情報を取得するハンドラです
// @Summary 全アドレス取得
// @Description DB から全アドレスを取得し JSON で返します
// @Produce json
// @Success 200 {object} []models.Address
// @Failure 500 {object} map[string]string
// @Router /api/addresses [get]
func GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	addresses, err := repository.GetAllAddresses()
	if err != nil {
		http.Error(w, "Failed to retrieve addresses", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(addresses)
}

// 他のアドレス操作（POST, PUT, DELETE）のハンドラ
