package api

// UsersResponse ユーザ一覧のレスポンス
type UsersResponse struct {
	// Result 結果のリスト
	Result []UserResponse `json:"result"`
	// Count 件数
	Count int `json:"count"`
}

// UserResponse ユーザのレスポンス
type UserResponse struct {
	// ID ユーザのID
	ID int `json:"id"`
	// Name 名前
	Name string `json:"name"`
	// Age 年齢
	Age int `json:"age"`
}

// RegisteredResponse 登録時のレスポンス
type RegisteredResponse struct {
	ID int `json:"id"`
}
