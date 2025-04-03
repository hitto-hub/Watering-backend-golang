// package database - アプリケーションのデータベース接続を管理するパッケージです。
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// グローバルなデータベース接続
var DB *sql.DB

// InitDB はデータベース接続を初期化
func InitDB() {
	var err error
	// SQLite の DB ファイルに接続
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// 以下にテーブル作成処理を追加
	// createTables()
}

/*
// createTables はテーブルを作成（省略可）
func createTables() {
	// SQL クエリを実行してテーブルを作成
}
*/
