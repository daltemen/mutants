package datasources

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func ConnectRedis() *redis.Pool {
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	const maxConnections = 10

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", redisAddr)
	}, maxConnections)

	return redisPool
}
