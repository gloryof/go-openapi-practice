package api

import "strconv"

// validateRequest 入力チェックを行う
func validateRequest(request UserUpdateRequest) []string {

	var msg []string

	n := request.Name
	if n == "" {
		msg = append(msg, "名前は必須です。")
	} else {

		if 30 < len(n) {

			msg = append(msg, "名前は30文字以内で入力してください。")
		}
	}

	a := request.Age
	if a == "" {
		msg = append(msg, "年齢は必須です。")
	} else {

		_, err := strconv.Atoi(a)
		if err != nil {

			msg = append(msg, "年齢には数値を入れてください。")
		}
	}

	return msg
}
