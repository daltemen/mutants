package repositories

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"mutants/app/mutants"
	"strconv"
)

type redisReadRepository struct {
	redisPool *redis.Pool
}

func NewRedisReadRepository(redisPool *redis.Pool) ReadRepository {
	return &redisReadRepository{redisPool: redisPool}
}

func (r *redisReadRepository) GetStats(ctx context.Context) (mutants.HumanStats, error) {
	conn := r.redisPool.Get()
	defer conn.Close()

	mCount, err := redis.Bytes(conn.Do("GET", "mutants"))
	mutantsCount, _ := strconv.ParseInt(string(mCount), 10, 64)
	if err != nil {
		return mutants.HumanStats{}, err
	}

	hCount, err := redis.Bytes(conn.Do("GET", "humans"))
	humansCount, _ := strconv.ParseInt(string(hCount), 10, 64)
	if err != nil {
		return mutants.HumanStats{}, err
	}

	return mutants.HumanStats{Mutants: mutantsCount, Humans: humansCount}, nil
}
