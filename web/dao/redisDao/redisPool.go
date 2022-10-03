package redisdao

import (
	"ahutoj/web/utils"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStat int

const (
	RDB_INIT  RedisStat = 0
	RDB_CLOSE RedisStat = 1
	RDB_FREE  RedisStat = 2
	RDB_BUSY  RedisStat = 3
)

type RedisItem struct {
	rdb      *redis.Client
	Status   RedisStat
	lastTime int64
}
type RedisPool struct {
	rdbs      []RedisItem
	Addr      string
	Password  string
	DB        int
	PoolSize  int16
	PoolLive  bool
	KeepAlive int64
	lock      sync.Mutex
}

var redisPool RedisPool

func CloseTimeoutRDB() {
	for {
		currentTime := time.Now().Unix()
		redisPool.lock.Lock()
		for idx := range redisPool.rdbs {
			rdb := &redisPool.rdbs[idx]
			if rdb.lastTime-currentTime < 0 && rdb.Status == RDB_FREE {
				rdb.lastTime = 0
				rdb.Status = RDB_CLOSE
				rdb.rdb.Close()
				rdb.rdb = nil
			}
		}
		redisPool.lock.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func InitRedisPool() error {
	logger := utils.GetLogInstance()
	config := utils.GetConfInstance().RedisConfig
	logger.Debugf("redis config:%+v", utils.Sdump(config))
	redisPool = RedisPool{
		rdbs:      make([]RedisItem, config.PoolSize),
		Addr:      config.Host + config.Port,
		Password:  config.Password,
		DB:        config.Db,
		PoolSize:  config.PoolSize,
		PoolLive:  true,
		KeepAlive: config.KeepAlive * int64(time.Second),
	}
	go CloseTimeoutRDB()
	return nil
}

func connectRDB(rdbfd int) error {
	rdb := &redisPool.rdbs[rdbfd]
	rdb.Status = RDB_BUSY
	rdb.lastTime = time.Now().Unix()
	rdb.rdb = redis.NewClient(&redis.Options{
		Addr:     redisPool.Addr,
		Password: redisPool.Password,
		DB:       redisPool.DB,
	})
	return nil
}

func findFreeRDB() int {
	waitOpenRdb := -1
	redisPool.lock.Lock()
	defer redisPool.lock.Unlock()
	for idx := range redisPool.rdbs {
		rdb := &redisPool.rdbs[idx]
		if rdb.Status == RDB_FREE {
			rdb.lastTime = time.Now().Unix()
			rdb.Status = RDB_BUSY

			return idx
		}
		if waitOpenRdb == -1 && rdb.Status <= RDB_FREE {
			waitOpenRdb = idx
		}
	}
	err := connectRDB(waitOpenRdb)
	if err != nil {
		return -1
	}
	return waitOpenRdb
}

func DestoryRedisPool() {
	redisPool.PoolLive = false
}

func GetRedis() int {
	idx := -1
	if redisPool.PoolLive {
		idx = findFreeRDB()
	}
	return idx
}

func GetRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisPool.Addr,
		Password: redisPool.Password,
		DB:       redisPool.DB,
	})
	return rdb
}
