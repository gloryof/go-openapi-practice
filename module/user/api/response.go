package api

// UsersResponse ユーザ一覧のレスポンス
type UsersResponse struct {
	// Result 結果のリスト
	// データがない場合は空のリストが返る。
	Result []UserResponse `json:"result"`
	// Count 件数
	// データがない場合は0。
	Count int `json:"count" example:"3"`
}

// UserResponse ユーザのレスポンス
type UserResponse struct {
	// ID ユーザのID
	ID int `json:"id" example:"1"`
	// Name 名前
	Name string `json:"name" example:"テスト 太郎"`
	// Age 年齢
	Age int `json:"age" example:"24"`
}

// RegisteredResponse 登録時のレスポンス
type RegisteredResponse struct {
	ID int `json:"id" example:"1"`
}
