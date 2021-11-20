package ipLocationService

import (
"fmt"
"github.com/go-redis/redis"
	"os"
)

var RedisClient redis.Client

func InitDb() {

	client := redis.NewClient(&redis.Options{
		Addr:   os.Getenv("REDIS_CONNECT_URL") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	RedisClient = *client
	LoadData("ipblocks")
}


