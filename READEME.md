# GoのOpenAPIのお試し場所
OpenAPIのドキュメントを構成する上でどうやって設計するかやどのライブラリを使うかを試す場所。

## APIの呼び出し

### ユーザ一覧
```
$ curl -H "Authorization:Bearer test" -H "X-API-VERSION:1" http://localhost:8000/users
```

### ユーザ取得
```
$ curl -H "Authorization:Bearer test" -H "X-API-VERSION:1" http://localhost:8000/users/1
```

### ユーザの登録
```
$ curl -X POST -H "Authorization:Bearer test" -H "X-API-VERSION:1" -H "Content-Type:application/json" -d '{"name":"テスト", "age":"25"}' http://localhost:8000/users
```
### ユーザの更新
```
$ curl -X PUT -H "Authorization:Bearer test" -H "X-API-VERSION:1" -H "Content-Type:application/json" -d '{"name":"テスト", "age":"25"}' http://localhost:8000/users/1
```
### ユーザの削除
```
$ curl -X DELETE -H "Authorization:Bearer test" -H "X-API-VERSION:1" http://localhost:8000/users/1
```