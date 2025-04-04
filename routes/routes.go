// package routes - アプリケーションのルーティングを管理するパッケージです。
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hitto-hub/PlantaTalk/controllers"
)

// SetupRoutes - 全 HTTP ルートを設定してハンドラを返します
func SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// アドレス関連のルート
	router.HandleFunc("/api/v1/addresses", controllers.GetAddressesHandler).Methods("GET")
	// ※ POST, PUT, DELETE などのルートは必要に応じて追加

	// 水やり指示関連のルート
	router.HandleFunc("/api/v1/instructions", controllers.SetInstructionHandler).Methods("POST")
	// ※ GET (最新取得など) のルートも追加可能

	// 定期水やりスケジュールのルート
	router.HandleFunc("/api/v1/watering_regular", controllers.SetWateringRegularHandler).Methods("POST")
	// ※ GET, DELETE のルートも追加可能

	return router
}
