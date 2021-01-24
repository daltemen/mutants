package repositories

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/suite"
	"mutants/app/mutants"
	"testing"
)

type redisEventualRepoSuite struct {
	suite.Suite
	repository EventualWriteRepository
	ctx        context.Context
	miniRedis  *miniredis.Miniredis
}

func (suite *redisEventualRepoSuite) SetupTest() {
	suite.ctx = context.Background()
	miniRedis, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	suite.miniRedis = miniRedis
	redisPool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", miniRedis.Addr())
	}, 11)
	suite.repository = NewRedisEventualRepository(redisPool)
}

func (suite *redisEventualRepoSuite) TearDownSuite() {
	suite.miniRedis.Close()
}

func TestRedisEventualSuite(t *testing.T) {
	suite.Run(t, &redisEventualRepoSuite{})
}

func (suite *redisEventualRepoSuite) TestRedisEventualWriteRepository_IncrementCount() {
	err1 := suite.repository.IncrementCount(suite.ctx, mutants.HumanType)
	err2 := suite.repository.IncrementCount(suite.ctx, mutants.MutantType)

	humansCount, _ := suite.miniRedis.Get("humans")
	mutantsCount, _ := suite.miniRedis.Get("mutants")

	suite.NoError(err1)
	suite.NoError(err2)
	suite.Equal("1", humansCount)
	suite.Equal("1", mutantsCount)
}
