package repositories

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"mutants/app/mutants"
)

type redisEventualWriteRepository struct {
	redisPool *redis.Pool
}

func NewRedisEventualRepository(redisPool *redis.Pool) EventualWriteRepository {
	return &redisEventualWriteRepository{redisPool: redisPool}
}

func (r *redisEventualWriteRepository) IncrementCount(ctx context.Context, humanType mutants.HumanTypes) error {
	conn := r.redisPool.Get()
	defer conn.Close()

	counterKey := "mutants"
	if humanType == mutants.HumanType {
		counterKey = "humans"
	}
	_, err := redis.Int(conn.Do("INCR", counterKey))
	return err // TODO: handle err
}
