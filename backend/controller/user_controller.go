package controller

import (

	"github.com/labstack/echo/v4"
)

// controllerのインタフェースを定義
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// controllerの構造体を定義
type userController struct {
	// uu usecase.IUserUsecase
}

// コンストラクタ（usecaseとのDI）
// func NewUserController(uu usecase.IUserUsecase) IUserController {
// 	// 構造体の実体のポインタを返す
// 	return &userController{uu}
// }