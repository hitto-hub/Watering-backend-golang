// package repository - アプリケーションのデータベース操作を定義するパッケージです。
package repository

import (
	"github.com/hitto-hub/PlantaTalk/database"
	"github.com/hitto-hub/PlantaTalk/models"
)

// GetAllAddresses - GORMを使って全てのアドレス情報を取得します
func GetAllAddresses() ([]models.Address, error) {
	var addresses []models.Address

	// GORM の Find で全件取得
	if err := database.DB.Find(&addresses).Error; err != nil {
		return nil, err
	}

	return addresses, nil
}

// 他の CRUD 操作関数も以下に追加可能
