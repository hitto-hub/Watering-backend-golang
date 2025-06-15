# Watering-backend-golang

- [Watering-Back-python](https://github.com/hitto-hub/Watering-back-python) をgolangに書き換える

PlantaTalk API Serverは、植物の自動水やりシステムを管理するためのGoで書かれたREST APIサーバーです。

センサーデータの管理、定期的な水やりスケジュール、手動による水やり指示を提供します。

## 特徴

## architecture

```tree
PlantaTalk-Backend
├── config/          # 設定管理
├── controllers/     # HTTPハンドラー
├── database/        # データベース接続・初期化
├── models/          # データモデル定義
├── repository/      # データアクセス層
├── routes/          # ルーティング設定
├── scheduler/       # バックグラウンドスケジューラー
└── services/        # ビジネスロジック
```

## インストール

### 設定ファイルの作成

```bash
cp config.yaml.example config.yaml
```

`config.yaml`を編集してサーバー設定を調整

```yaml
server:
  port: 8080
```
