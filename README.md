# yandere-server

## テストタスク登録

curl -i -H "Origin:http://localhost" localhost:8080/task/test/create
curl -i -H "Origin:http://localhost" localhost:8080/task/test/read
curl -i -H "Origin:http://localhost" localhost:8080/task/progress
curl -X POST -H "Content-Type: application/json" -H "Origin:http://localhost" -d '{"task":"example_task","desc":"hogehoge","task_id":"example_id"}' localhost:8080/task/create

## Docker滅びの呪文（SQL関連の更新が来たときはこれを実行することを推奨）

docker-compose down --rmi all --volumes --remove-orphans

## API仕様書について

### ビルド手順

1. npm install -g aglio
2. build.shを実行

apiディレクトリ直下にdocs.htmlが出力される。

## テスト方法

run.shを実行してください

- 登録用API：http://localhost:8080/task/create
- 更新用API：http://localhost:8080/task/update
- 進捗率計算用API：http://localhost:8080/task/progress
