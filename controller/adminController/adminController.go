package adminController

import (
	"net/http"

	"github.com/KaungHtetMon29/BreakPoint_Backend/api_gen/admin"
	"github.com/labstack/echo/v4"
)

type Admin struct {
}

func NewAdminCtrler() *Admin {
	return &Admin{}
}

func (pc *Admin) InsertAdmin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Insert Admin")
}

func (pc *Admin) AdminLogin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Admin Login")
}

func (pc *Admin) AdminLogout(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Admin Logout")
}

func (pc *Admin) GetUsers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Users")
}

func (pc *Admin) GetUserDetails(ctx echo.Context, id admin.Id) error {
	return ctx.JSON(http.StatusOK, "Get User Details")
}

func (pc *Admin) GetUserStatus(ctx echo.Context, id admin.Id) error {
	return ctx.JSON(http.StatusOK, "Get User Status")
}

func (pc *Admin) DeleteAdmin(ctx echo.Context, adminId admin.AdminId) error {
	return ctx.JSON(http.StatusOK, "Delete Admin")
}

func (pc *Admin) GetAdminDetail(ctx echo.Context, adminId admin.AdminId) error {
	return ctx.JSON(http.StatusOK, "Get Admin Detail")
}

func (pc *Admin) UpdateAdminDetail(ctx echo.Context, adminId admin.AdminId) error {
	return ctx.JSON(http.StatusOK, "Update Admin Detail")
}
