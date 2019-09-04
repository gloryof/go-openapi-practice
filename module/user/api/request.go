package api

// UserUpdateRequest ユーザ更新リクエスト
type UserUpdateRequest struct {
	// Name 名前
	Name string `json:"name" maxLength:"30" example:"テスト 太郎"`
	// Age 年齢
	Age string `json:"age" example:"24"`
}
