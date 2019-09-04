package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gloryof/go-openapi-practice/module/base/api"
	"github.com/labstack/echo/v4"
)

// ListUsers ユーザの一覧を取得する
// @router /users [GET]
// @summary ユーザの一覧を取得する
// @description 全てのユーザの一覧を取得する。
// @tags User,Reference
// @produce appliction/json
// @security ApiKeyAuth
//
// @param X-API-VERSION header string true "APIバージョンヘッダ"
//
// @success 200 {object} api.UsersResponse "正常に処理された場合"
// @failure 400 {object} api.ErrorResponse "共通ヘッダーなどが設定されていない場合"
func ListUsers(c echo.Context) error {

	ls := UsersResponse{
		Result: getAllUsers(),
		Count:  3,
	}

	return c.JSON(http.StatusOK, ls)
}

// GetUser ユーザを取得する
// @router /users/{id} [GET]
// @summary ユーザを取得する
// @description IDで指定したユーザを取得する。
// @tags User,Reference
//
// @produce appliction/json
// @security ApiKeyAuth
//
// @param X-API-VERSION header string true "APIバージョンヘッダ"
// @param id path int true "ユーザID"
//
// @success 200 {object} api.UsersResponse "正常に処理された場合"
// @failure 400 {object} api.ErrorResponse "共通ヘッダーなどが設定されていない場合"
// @failure 404 {object} api.ErrorResponse "データが存在しない場合"
func GetUser(c echo.Context) error {

	res, err := checkAndFind(c)
	if !c.Response().Committed {

		return c.JSON(http.StatusOK, res)
	}

	return err
}

// UpdatetUser ユーザを更新する
// @router /users/{id} [PUT]
// @summary ユーザを更新する
// @description 実際はチェックのみで更新はしない
//
// @tags User,Update
// @produce appliction/json
// @security ApiKeyAuth
//
// @param X-API-VERSION header string true "APIバージョンヘッダ"
// @param id path int true "ユーザID"
// @param param body api.UserUpdateRequest true "ユーザ更新リクエスト"
//
// @success 200 {object} api.RegisteredResponse "正常に更新できた場合"
// @failure 400 {object} api.ErrorResponse "共通ヘッダーなどが設定されていない場合"
// @failure 404 {object} api.ErrorResponse "データが存在しない場合"
func UpdatetUser(c echo.Context) error {

	_, err := checkAndFind(c)
	if !c.Response().Committed {

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

	return err
}

// DeletetUser ユーザを削除する
// @router /users/{id} [DELETE]
// @summary ユーザを削除する
// @description 実際はチェックのみで削除はしない
//
// @tags User,Update
// @produce appliction/json
// @security ApiKeyAuth
//
// @param X-API-VERSION header string true "APIバージョンヘッダ"
// @param id path int true "ユーザID"
//
// @success 204 "正常に更新できた場合"
// @failure 400 {object} api.ErrorResponse "共通ヘッダーなどが設定されていない場合"
// @failure 404 {object} api.ErrorResponse "データが存在しない場合"
func DeletetUser(c echo.Context) error {

	_, err := checkAndFind(c)
	if !c.Response().Committed {

		return c.NoContent(http.StatusNoContent)
	}

	return err
}

// RegisterUser ユーザの登録
// @router /users/{id} [POST]
// @summary ユーザを登録する
// @description 実際はチェックのみで更新はしない
//
// @tags User,Update
// @produce appliction/json
// @security ApiKeyAuth
//
// @param X-API-VERSION header string true "APIバージョンヘッダ"
// @param id path int true "ユーザID"
// @param param body api.UserUpdateRequest true "ユーザ更新リクエスト"
//
// @success 200 {object} api.RegisteredResponse "正常に更新できた場合"
// @failure 400 {object} api.ErrorResponse "共通ヘッダーなどが設定されていない場合"
// @failure 404 {object} api.ErrorResponse "データが存在しない場合"
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
