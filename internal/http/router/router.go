package router

import (
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/handler"
	"github.com/labstack/echo/v4"
)

const (
	Admin = "admin"
	User  = "user"
)

var (
	allRoles  = []string{Admin, User}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{User}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Registration,
		},
	}

}

func PrivateRoutes(userHandler *handler.UserHandler, jadwalHandler *handler.JadwalHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.CreateUser,
			Roles:   allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUser,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},

		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Roles:   allRoles,
		},

		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.POST,
			Path:    "/jadwals",
			Handler: jadwalHandler.CreateJadwal,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.GET,
			Path:    "/jadwals:id",
			Handler: jadwalHandler.GetJadwalByID,
			Roles:   allRoles,
		},

		{
			Method:  echo.PUT,
			Path:    "/jadwals/:id",
			Handler: jadwalHandler.UpdateJadwal,
			Roles:   onlyAdmin,
		},

		{
			Method:  echo.DELETE,
			Path:    "/jadwals/:id",
			Handler: jadwalHandler.DeleteJadwal,
			Roles:   onlyAdmin,
		},
	}
}
