package rediscache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const nDuration = 5 * 60 * 60 * time.Second

func SetKey(ctx context.Context, rdbfd int, key string, value interface{}) error {
	rdb := redisPool.rdbs[rdbfd].rdb
	if rdb != nil {
		jvalue, _ := json.Marshal(value)
		cmd := rdb.Set(ctx, key, string(jvalue), nDuration)
		res, _ := cmd.Result()
		fmt.Println(res)
		return cmd.Err()
	}
	return errors.New("rdbpool is not exits")
}

func GetKey(ctx context.Context, rdbfd int, key string, ret interface{}) error {
	var str string
	rdb := redisPool.rdbs[rdbfd].rdb
	if rdb != nil {
		cmd := rdb.Get(ctx, key)
		err := cmd.Scan(&str)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(str), ret)
		return err
	}
	return errors.New("rdbpool is not exits")
}

func DelKey(ctx context.Context, rdbfd int, key string) error {
	rdb := redisPool.rdbs[rdbfd].rdb
	if rdb != nil {
		err := rdb.Del(ctx, key).Err()
		return err
	}
	return errors.New("rdbpool is not exits")
}
