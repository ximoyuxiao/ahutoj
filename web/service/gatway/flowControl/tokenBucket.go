package flowcontrol

import (
	"context"
	"sync"
	"time"
)

type TokenBucketWork interface {
	run(ctx context.Context)
}
type TokenBucket struct {
	TokenNumber int64
	MAXTOKEN    int64
	TokenSpeed  int64 //每秒钟 放多少个令牌
	Live        bool
	mtx         sync.Mutex
}

func (tb *TokenBucket) addToken(token int64) int64 {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	if tb.TokenNumber+token > tb.MAXTOKEN {
		ret := tb.MAXTOKEN - tb.TokenNumber
		tb.TokenNumber = tb.MAXTOKEN
		return ret
	}
	tb.TokenNumber += token
	return token
}

func (tb *TokenBucket) FetchToken(token int64) int64 {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	if token < 0 {
		return -1
	}
	if tb.TokenNumber < token {
		return 0 // token不够用,等下一次
	}
	tb.TokenNumber -= token
	return token
}

func (tb *TokenBucket) ReturnToken(token int64) int64 {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	size := tb.addToken(token)
	if size < 0 {
		return size
	}
	return token
}
func TokenThread(tb *TokenBucket) {
	for tb.Live {
		time.Sleep(time.Second)
		tb.addToken(tb.TokenSpeed)
	}
}
func InitTokenBucket(maxToken, tokenSpeed int64) *TokenBucket {
	tb := TokenBucket{}
	tb.MAXTOKEN = maxToken
	tb.TokenNumber = 0
	tb.TokenSpeed = tokenSpeed
	tb.Live = true
	go TokenThread(&tb)
	return &tb
}

func DestoryToken(tb *TokenBucket) {
	tb.Live = false
}
