package redisdao

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

func SetKey(ctx context.Context, rdbfd int, key string, value interface{}) error {
	rdb := redisPool.rdbs[rdbfd].rdb
	if rdb != nil {
		jvalue, _ := json.Marshal(value)
		rdb.Set(ctx, key, string(jvalue), time.Hour)
	}
	return nil
}

func GetKey(ctx context.Context, rdbfd int, key string, ret interface{}) error {
	var str string
	rdb := redisPool.rdbs[rdbfd].rdb
	if rdb != nil {
		cmd := rdb.Get(ctx, key)
		err := cmd.Scan(&str)
		json.Unmarshal([]byte(str), ret)
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
