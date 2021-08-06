# API Docs

## Controller

Todo: API Blueprint or swaggerで整備

**JWTですね...**

フロント<->バックエンド部（Controllerに当たる部分）

- アジェンダ
- 会議
- タスク
- 人
- 組織

### アジェンダCRUD

アジェンダ登録 POST, /agendas/create

```json
{
	"title":"hogehoge",
	"meeting_id": "example_meeting_id",
	"place": "Example会議室",
	"agendas":[
		{
			"goal": 2
			"participants": ["user_id_1","user_id_2","user_id_3"]
			"category": 1
		}
	],
	"token":"example-exapmle-example", //OAuth2.0頑張りましょう！
}
```

アジェンダ登録（レスポンス）

```json
200
//
{
	"message": "hogehoge",
	"agenda_id": "hogehoge-hogehoge"
}

//失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

アジェンダ取得(GET)/agenda?agenda_id=<agenda_id>

```json
レスポンス
{
	"title":"hogehoge",
	"created_at":"2021-07-09T05:30:44+00:00",
	"updated_at":"2021-07-09T05:30:44+00:00", 
	"place": "Example会議室",
	"agendas":[
		{
			"goal": 2
			"participants": ["user_id_1","user_id_2","user_id_3"]
			"category": 1
		}
	],
	"token":"example-exapmle-example", //OAuth2.0頑張りましょう！
}
```

アジェンダ更新(POST) /agenda/update

```json
{
	"title":"hogehoge",
	"place": "Example会議室",
	"agendas":[
		{
			"goal": 2
			"participants": ["user_id_1","user_id_2","user_id_3"]
			"category": 1
		}
	],
	"token":"example-exapmle-example", //OAuth2.0頑張りましょう！
	"agenda_id": "example-agenda-id",
}
```

アジェンダ更新（レスポンス）

```json
200
// 成功時
{
	"message": "hogehoge",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

アジェンダ削除(POST) /agenda/delete

```json
{
	"token": "hogehoge",
	"agenda_id": "削除",
}
```

アジェンダ削除（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

---

---

### 会議CRUD

会議登録（POST)

```json
{
	"title":"example会議",
	"agenda_id": "example-agenda-id",
	"scheudled_at": "2021-07-19T05:30:44+00:00",
	"place": "example会議室",
}
```

会議登録（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
	"meeting_id" : "example-meeting-id",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

会議更新（POST）

```json
{
	"meeting_id": "example-meeting-id",
	"title":"example会議",
	"agenda_id": "example-agenda-id",
	"scheudled_at": "2021-07-19T05:30:44+00:00",
	"place": "example会議室",
}
```

会議更新（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

会議情報取得（GET）/meeting?meeting_id=<meeting_id>

```json
{
	"meeting_id": "example-meeting-id",
	"title":"example会議",
	"agenda_id": "example-agenda-id",
	"scheudled_at": "2021-07-19T05:30:44+00:00",
	"created_at": "2021-07-05T05:30:44+00:00",
	"updated_at": "2021-07-19T05:30:44+00:00",
	"place": "example会議室",
}
```

---

---

---

### タスクCRUD

タスク登録（POST）/agenda/tasks/create

```json
{
	"agenda_id": "hogehoge-hogehoge",
	"todos":[
			{
				"task":"hogehoge",
				"assign":["user_id_1"],
				"is_done":false
			}
		]
	},
}
```

タスク登録（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
	"task_id": "example-task-id",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

タスク取得（GET）/agenda/tasks?task_id=<task_id>

```json
{
	"agenda_id": "hogehoge-hogehoge",
	"todos":[
			{
				"task":"hogehoge",
				"assign":["user_id_1"],
				"is_done":false
			}
		]
	},
	"created_at": "2021-07-05T05:30:44+00:00",
	"updated_at": "2021-07-19T05:30:44+00:00",
}
```

---

タスク削除（POST）/agenda/tasks/delete

```json
{
	"task_id": "example_task_id",
}
```

タスク削除（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

---

タスク更新（POST）/agenda/tasks/update

```json
{
	"task_id": "example-task-id",
	"agenda_id": "hogehoge-hogehoge",
	"todos":[
			{
				"task":"hogehoge",
				"assign":["user_id_1"],
				"is_done":false
			}
		]
	},
}
```

タスク更新（レスポンス）

```json
// 成功時
{
	"message": "hogehoge",
	"updated_at": "2021-07-19T05:30:44+00:00",
}

// 失敗時
{
	"errors": [
    {
      "message": "Bad Authentication data",
      "code": 215  //これユーザーが知れるようにでも良いし、デバッグ用コードにしても良いかも（コケた箇所に応じてコードを入力する）
    }
  ]
}
```

### ユーザーCRUD (Teams連携）

- ユーザー
- 権限

### 組織CRUD