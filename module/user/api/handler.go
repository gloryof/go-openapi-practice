package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gloryof/go-openapi-practice/module/base/api"
	"github.com/labstack/echo/v4"
)

// ListUsers ユーザの一覧を取得する
func ListUsers(c echo.Context) error {

	ls := UsersResponse{
		Result: getAllUsers(),
		Count:  3,
	}

	return c.JSON(http.StatusOK, ls)
}

// GetUser ユーザを取得する
func GetUser(c echo.Context) error {

	res, err := checkAndFind(c)
	if err != nil {

		return err
	}

	return c.JSON(http.StatusOK, res)
}

// UpdatetUser ユーザを更新する
// 実際はチェックのみで更新はしなし
func UpdatetUser(c echo.Context) error {

	_, err := checkAndFind(c)
	if err != nil {

		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// DeletetUser ユーザを削除する
// 実際はチェックのみで削除はしなし
func DeletetUser(c echo.Context) error {

	_, err := checkAndFind(c)
	if err != nil {

		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// RegisterUser ユーザの登録
// 実際はチェックのみで登録はしない
func RegisterUser(c echo.Context) error {

	req := UserUpdateRequest{}

	c.Bind(&req)

	msg := validateRequest(req)

	if 0 < len(msg) {

		return c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Summary: "入力データに不正がありました。",
			Details: msg,
		})
	}

	return c.JSON(http.StatusOK, RegisteredResponse{
		ID: 4,
	})
}

// getAllUsers 全てのユーザを取得する
func getAllUsers() []UserResponse {

	return []UserResponse{
		{
			ID:   1,
			Name: "テスト1",
			Age:  20,
		},
		{
			ID:   2,
			Name: "テスト2",
			Age:  31,
		},
		{
			ID:   3,
			Name: "テスト3",
			Age:  43,
		},
	}
}

// finldByID IDをキーに検索する
func finldByID(id int) (UserResponse, error) {

	for _, v := range getAllUsers() {

		if v.ID == id {

			return v, nil
		}
	}

	return UserResponse{}, errors.New("対象のIDのユーザが見つかりません。")
}

// checkAndFind 入力チェックと検索を行う
func checkAndFind(c echo.Context) (UserResponse, error) {

	id := c.Param("id")
	if id == "" {

		return UserResponse{}, c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Summary: "入力データに不正がありました。",
			Details: []string{
				"IDが指定されていません。",
			},
		})
	}

	cid, err := strconv.Atoi(id)

	if err != nil {

		return UserResponse{}, c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Summary: "入力データに不正がありました。",
			Details: []string{
				"IDには数値を設定してください。",
			},
		})
	}

	res, rerr := finldByID(cid)

	if rerr != nil {

		return UserResponse{}, c.JSON(http.StatusNotFound, api.ErrorResponse{
			Summary: "対象のデータ見つかりません。",
			Details: []string{
				"対象のIDのユーザ見つかりません。",
			},
		})
	}

	return res, nil
}
