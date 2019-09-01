package main

import (
	"github.com/gloryof/go-openapi-practice/externals/server"
	_ "github.com/gloryof/go-openapi-practice/module/docs"
)

// @title OpenAPIドキュメントのお試し
// @version 1.0
//
// @description このアプリで試そうと思っていることは下記。
// @description - ドキュメント自体がどこまで手軽にかけるか
// @description - 認証周りは動作するか
// @description - 共通レスポンスなどが簡単にかけるか
// @description - ドキュメントのバージョン管理がうまくできるか
// @description - CIによるビルドができるかどうか
//
// @host localhost:8000
// @BasePath /api
//
// @tag.name User
// @tag.description ユーザに関するAPI
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	s := server.NewServer()

	s.Start()
}
