# yandere-server

## テストタスク登録

curl -i -H "Origin:http://localhost" localhost:8080/task/test/create
curl -i -H "Origin:http://localhost" localhost:8080/task/test/read
  
curl -X POST -H "Content-Type: application/json" -H "Origin:http://localhost" -d '{"task":"example_task","desc":"hogehoge","task_id":"example_id"}' localhost:8080/task/create

## Docker滅びの呪文（SQL関連の更新が来たときはこれを実行することを推奨）

docker-compose down --rmi all --volumes --remove-orphans
