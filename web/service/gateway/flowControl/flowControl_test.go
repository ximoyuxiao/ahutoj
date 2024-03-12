package flowcontrol

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Mwork struct {
	Idx int
}

func (mwork *Mwork) Run(ctx context.Context) {
	fmt.Println(mwork.Idx)
}

var i = 0

func getTask() {
	rd := rand.Int31()%1000 + 1
	time.Sleep(time.Duration(rd * int32(time.Millisecond)))
}

func ExecuteTask(t *testing.T) {
	t.Errorf("%v", i)
	i++
}
func TestTokenBucket(t *testing.T) {
	MAXTOKEN := int64(1024)
	TOKEN_SPEED := int64(10)
	tb := InitTokenBucket(MAXTOKEN, TOKEN_SPEED) //初始化一个漏桶
	var token int64 = 0
	for i := 0; i < 50; i++ {
		// 接受一个任务
		getTask()
		var ti int64 = 1
		if token <= 0 {
			token = tb.FetchToken(ti)
		}
		if token <= 0 {
			t.Error("have no token")
			t.Logf("have no token\n")
			continue
		}
		ExecuteTask(t)
	}
}
