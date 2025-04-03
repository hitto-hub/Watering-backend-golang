package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hitto-hub/PlantaTalk/database"
	"github.com/hitto-hub/PlantaTalk/services"
)

// SetInstructionHandler は水やり指示を作成するハンドラ
// @Summary 水やり指示作成
// @Description 指定アドレスに対して水やり指示を作成
// @Param address query int true "Address"
// @Param instruction body struct{ Instruction int } true "Instruction"
// @Produce json
// @Success 200 {object} models.Instruction
// @Failure 400 {object} map[string]string
// @Router /api/instructions [post]
func SetInstructionHandler(w http.ResponseWriter, r *http.Request) {
	addressStr := r.URL.Query().Get("address")
	address, err := strconv.Atoi(addressStr)
	if err != nil {
		http.Error(w, "Invalid address", http.StatusBadRequest)
		return
	}

	var req struct {
		Instruction int `json:"instruction"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ビジネスロジックで指示を作成
	instruction := services.ExecuteWatering(address)

	// ここで DB への保存処理を行う（省略）
	_ = database.DB

	json.NewEncoder(w).Encode(instruction)
}
