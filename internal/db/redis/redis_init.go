package redis

import (
	"KatodForAnod/siteGallery/internal/db"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type redisClient struct {
	conn *redis.Client
}

func ConnectRedis(ctx context.Context) db.DatabaseUserAuth {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pong)

	return redisClient{conn: client}
}
