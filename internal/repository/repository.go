package repository

import (
	"log"

	"github.com/hitto-hub/PlantaTalk/internal/database"
	"github.com/hitto-hub/PlantaTalk/internal/models"
)

// GetAllAddresses は DB から全てのアドレス情報を取得します
func GetAllAddresses() ([]models.Address, error) {
	rows, err := database.DB.Query("SELECT address, name FROM addresses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var addr models.Address
		if err := rows.Scan(&addr.Address, &addr.Name); err != nil {
			log.Printf("Error scanning address: %v", err)
			continue
		}
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

// 他の CRUD 操作関数も以下に追加可能
