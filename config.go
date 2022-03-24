package ipLocationService

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/klovercloud-dev/get-ip-region/enums"
	"os"
)

var RedisClient redis.Client

func InitDb() {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv(string(enums.REDISCONNECTURL)) + ":" + os.Getenv(string(enums.REDISPORT)),
		Password: "",
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	RedisClient = *client
	LoadData(string(enums.IPBLOCKS))
}
