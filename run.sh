#!/bin/bash

# DBの初期化などのための環境リセット＆開発環境構築用スクリプト

# 全データの消去
docker-compose down --rmi all --volumes --remove-orphans

# ビルド
docker-compose build --no-cache

# 起動
docker-compose up -d

# ログの表示
docker-compose logs
