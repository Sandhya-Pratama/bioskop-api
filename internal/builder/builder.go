package builder

import (
	"github.com/Sandhya-Pratama/bioskop-api/internal/config"
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/handler"
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/router"
	"github.com/Sandhya-Pratama/bioskop-api/internal/repository"
	"github.com/Sandhya-Pratama/bioskop-api/internal/service"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)

	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)
	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	jadwalRepository := repository.NewJadwalTayangRepository(db)
	jadwalService := service.NewJadwalTayangService(jadwalRepository)
	jadwalHandler := handler.NewJadwalHandler(jadwalService)
	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler, jadwalHandler)
}
