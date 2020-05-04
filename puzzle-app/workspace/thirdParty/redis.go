package thirdParty
import (
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func GetRedisClient() *redis.Client {

	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "redis.puzzle-space:6379", //"redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return client
}