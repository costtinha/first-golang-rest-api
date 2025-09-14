package main

import (
	"log"

	"github.com/costtinha/first-golang-rest-api/internal/config"
	"github.com/costtinha/first-golang-rest-api/internal/http"
	"github.com/costtinha/first-golang-rest-api/internal/logger"
	"github.com/costtinha/first-golang-rest-api/internal/platform/database"
	"github.com/costtinha/first-golang-rest-api/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error : %v", err)
	}

	lg := logger.New(cfg)

	db, err := database.Connect(cfg, lg)
	if err != nil {
		lg.Fatal("db connect error", "err", err)
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		lg.Fatal("auto-migrate error", "err", err)
	}

	userRepo := user.NewGormRepository(db)
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc, lg)

	router := http.NewRouter(cfg, lg)
	router.RegisterHealth()
	router.RegisterUserRouter(userHandler)

	lg.Info("Starting http server", "port", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		lg.Fatal("server error", "err", err)
	}

}
