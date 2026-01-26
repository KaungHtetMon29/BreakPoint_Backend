package authController

import (
	"fmt"
	"net/http"

	authentication "github.com/KaungHtetMon29/BreakPoint_Backend/internal/auth"

	"github.com/labstack/echo/v4"
)

type Auth struct {
	auth *authentication.Oauth
}

func NewAuthCtrler(auth *authentication.Oauth) *Auth {
	return &Auth{
		auth: auth,
	}
}

func (pc *Auth) Login(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"redirect_url": pc.auth.AuthCodeUrl,
	})
}

func (pc *Auth) Callback(ctx echo.Context) error {
	code := ctx.Request().FormValue("code")
	fmt.Println(ctx.Request().FormValue("code"))
	userInfo, err := pc.auth.GetGoogleUserInfo(code)
	if err != nil {
		return err
	}
	token, err := authentication.CreateJWTToken(userInfo)
	if err != nil {
		return err
	}
	fmt.Println(*token)
	cookie := http.Cookie{
		Name:     "test",
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	ctx.SetCookie(&cookie)
	return ctx.Redirect(http.StatusPermanentRedirect, "http://localhost:3000")
}

func (pc *Auth) Logout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "logout")
}

func (pc *Auth) SignUp(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "signup")
}

func (pc *Auth) GetProfile(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "get profile")
}
