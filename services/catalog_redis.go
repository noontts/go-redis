package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"goredis/repositories"
	"time"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{productRepo, redisClient}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "services::GetProducts"

	//Redis Get
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	//Repository
	productDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range productDB {
		products = append(products, Product{
			ID:      p.ID,
			Name:    p.Name,
			Quatity: p.Quatity,
		})
	}

	//Redis SET
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("Database")
	return products, nil
}
