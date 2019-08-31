package api

// UserUpdateRequest ユーザ更新リクエスト
type UserUpdateRequest struct {
	// Name 名前
	Name string `json:"name"`
	// Age 年齢
	Age string `json:"age"`
}
