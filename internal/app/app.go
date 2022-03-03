package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"rest-api-go-lang/internal/app/database"
	"rest-api-go-lang/internal/app/domain"
	"rest-api-go-lang/internal/app/domain/product"
	"rest-api-go-lang/internal/app/repository"
	"rest-api-go-lang/internal/app/web/router"
	"rest-api-go-lang/internal/app/web/validator"
	"sync"
	"time"
)

type app struct {
	sync.Mutex
	running  bool
	server   *echo.Echo
	services *domain.Services
	db       *sql.DB
}

var instance = new(app)

func Start() {
	instance.Lock()

	if instance.running {
		instance.Unlock()
		return
	}

	instance.running = true

	err := configureDatabase()
	if err != nil {
		instance.running = false
		instance.Unlock()
		return
	}

	configureServices()
	configureServer()

	instance.Unlock()
	err = instance.server.Start(":8282")
	if err != nil {
		instance.running = false
	}
}

func Stop() {
	instance.Lock()
	defer instance.Unlock()

	if !instance.running {
		return
	}

	_ = instance.server.Shutdown(context.Background())
	_ = instance.db.Close()

	instance.running = false
}

func configureDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("getting the env values")
	}
	fmt.Println(os.Getenv("POSTGRES_HOST"))
	fmt.Println(os.Getenv("POSTGRES_PORT"))
	fmt.Println(os.Getenv("POSTGRES_DATABASE"))
	fmt.Println(os.Getenv("POSTGRES_USERNAME"))
	fmt.Println(os.Getenv("POSTGRES_PASSWORD"))
	db, err := database.NewPostgres(&database.Config{
		Host:                  os.Getenv("POSTGRES_HOST"),
		Port:                  os.Getenv("POSTGRES_PORT"),
		Database:              os.Getenv("POSTGRES_DATABASE"),
		Username:              os.Getenv("POSTGRES_USERNAME"),
		Password:              os.Getenv("POSTGRES_PASSWORD"),
		MinConnections:        20,
		MaxConnections:        30,
		ConnectionMaxLifetime: 15 * time.Minute,
		ConnectionMaxIdleTime: 5 * time.Minute,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	instance.db = db
	return nil
}

func configureServices() {
	instance.services = &domain.Services{
		ProductService: product.NewService(repository.NewProductRepository(instance.db)),
	}
}

func configureServer() {
	instance.running = true
	instance.server = echo.New()
	instance.server.Use(middleware.Recover())
	instance.server.Validator = validator.New()

	router.Configure(instance.server, instance.services)
}
