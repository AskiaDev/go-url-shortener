package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)


type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour


func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing redis: %v", err))		
	}

	fmt.Printf("\nRedis started successfully: pong message = %s\n", pong)

	storeService.redisClient = redisClient
	return storeService
}


func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration)

	if err != nil {
		panic(fmt.Sprintf("Failed to save url mapping: %v", err))
	}
}


func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf("Failed to retrieve initial url: %v", err))
	}

	return result
}

