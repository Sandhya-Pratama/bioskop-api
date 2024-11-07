package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Sandhya-Pratama/bioskop-api/internal/builder"
	"github.com/Sandhya-Pratama/bioskop-api/internal/config"
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/binder"
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/server"
	"github.com/Sandhya-Pratama/bioskop-api/internal/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//membuat config ke env
	cfg, err := config.NewConfig(".env")
	checkError(err)

	//mengubah starter
	splash()

	//membuat koneksi ke postgres
	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	// //membuat koneksi ke redis
	// redisClient := buildRedis(cfg)

	// //koneksi midtrans
	// topUpHandler := builder.buildTopUp(db, cfg)

	//membuat routes
	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)

	//membuat binder
	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	//membuat server
	svr := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	//membuat start server
	runServer(svr, cfg.Port)

	//membuat shutdown
	waitForShutdown(svr)

}

// func untuk start server
func runServer(srv *server.Server, port string) {
	go func() {
		// Pastikan port diawali dengan ":"
		fmt.Printf("Starting server on port: %s\n", port)
		err := srv.Start(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Fatal(err)
		}
	}()
}

// func untuk shutdown
func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

// func untuk koneksi ke postgresql
func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ini buar ngecek gormnya
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// func untuk koneksi ke redis
// func buildRedis(cfg *config.Config) *redis.Client {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
// 		Password: cfg.Redis.Password,
// 		DB:       0,
// 	})
// 	return client
// }

// untuk membuat spalsh menggunakan patorjk.com
func splash() {
	colorReset := "\030[0m"

	splashText := `	
___.   .__              __                                        .__ 
\_ |__ |__| ____  _____|  | ______ ______           _____  ______ |__|
 | __ \|  |/  _ \/  ___/  |/ /  _ \\____ \   ______ \__  \ \____ \|  |
 | \_\ \  (  <_> )___ \|    <  <_> )  |_> > /_____/  / __ \|  |_> >  |
 |___  /__|\____/____  >__|_ \____/|   __/          (____  /   __/|__|
     \/              \/     \/     |__|                  \/|__|       
     
	`
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

// func untuk cek error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
