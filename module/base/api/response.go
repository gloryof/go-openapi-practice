package api

// ErrorResponse 共通エラーレスポンス
type ErrorResponse struct {
	// Summary 概要
	Summary string `json:"summary"`
	// Details 詳細
	Details []string `json:"details"`
}
