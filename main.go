package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redis := initRedis()
	_ = redis

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogServiceRedis(productRepo, redis)
	productHandlers := handlers.NewCatalogHandler(productService)

	app := fiber.New()

	app.Get("/products", productHandlers.GetProducts)

	app.Listen(":8080")
}

func initDatabase() *gorm.DB {
	dsn := "host=localhost user=sqladmin password=password dbname=goredis port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
