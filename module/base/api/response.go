package api

// ErrorResponse 共通エラーレスポンス
type ErrorResponse struct {
	// Summary 概要
	Summary string `json:"summary" example:"入力データに不備があります。"`
	// Details 詳細
	Details []string `json:"details" example:"名前は必須です。, 年齢は必須です。"`
}
