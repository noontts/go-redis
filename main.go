package main

import (
	"fmt"
	"goredis/repositories"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	redis := initRedis()
	productRepo := repositories.NewProductRepositoryRedis(db, redis)

	products, err := productRepo.GetProduct()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)

	// app := fiber.New()
	//
	// app.Get("/hello", func(c fiber.Ctx) error {
	// 	time.Sleep(time.Millisecond * 10)
	// 	return c.SendString("Hello World!!")
	// })
	//
	// app.Listen(":8080")
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
