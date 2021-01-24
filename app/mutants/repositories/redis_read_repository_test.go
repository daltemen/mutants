package repositories

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/suite"
	"testing"
)

type redisReadRepoSuite struct {
	suite.Suite
	repository ReadRepository
	ctx        context.Context
	miniRedis  *miniredis.Miniredis
}

func (suite *redisReadRepoSuite) SetupTest() {
	suite.ctx = context.Background()
	miniRedis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	suite.miniRedis = miniRedis
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", miniRedis.Addr())
	}, 11)
	suite.repository = NewRedisReadRepository(redisPool)
}

func (suite *redisReadRepoSuite) TearDownSuite() {
	suite.miniRedis.Close()
}

func TestRedisReadSuite(t *testing.T) {
	suite.Run(t, &redisReadRepoSuite{})
}

func (suite *redisReadRepoSuite) Test_redisReadRepository_GetStats_Success() {
	_, _ = suite.miniRedis.Incr("humans", 1)
	_, _ = suite.miniRedis.Incr("humans", 1)
	_, _ = suite.miniRedis.Incr("mutants", 1)
	_, _ = suite.miniRedis.Incr("mutants", 1)

	h, err := suite.repository.GetStats(suite.ctx)

	suite.NoError(err)
	suite.Equal(int64(2), h.Humans)
	suite.Equal(int64(2), h.Mutants)
}

func (suite *redisReadRepoSuite) Test_redisReadRepository_GetStats_Failed() {
	h, err := suite.repository.GetStats(suite.ctx)

	suite.NoError(err)
	suite.Equal(int64(2), h.Humans)
	suite.Equal(int64(2), h.Mutants)
}
