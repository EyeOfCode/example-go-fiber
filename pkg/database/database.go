package database

import (
	"context"
	"go-fiber-api/internal/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func ConnectRedis(config *config.Config) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Redis!")
	return redisClient, nil
}
