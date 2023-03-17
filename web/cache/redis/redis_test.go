package rediscache_test

import (
	rediscache "ahutoj/web/cache/redis"
	"ahutoj/web/utils"
	"context"
	"testing"
)

type Cat struct {
	Name  string
	Age   int64
	Color string
}

func init() {

}

func TestSetString(t *testing.T) {
	ctx := context.Background()
	t.Logf("config:%+v", utils.Sdump(utils.GetConfInstance()))
	rdbfd := rediscache.GetRedis()
	if rdbfd == -1 {
		t.Logf("Test connect Redis pool filed")
	}
	rediscache.SetKey(ctx, rdbfd, "str", "hello world")
	str := new(string)
	rediscache.GetKey(ctx, rdbfd, "str", str)
	if *str != "hello world" {
		t.Errorf("Test TestSetString failed,str=%s", *str)
	}
	rediscache.DelKey(ctx, rdbfd, "str")
}
func TestSetObj(t *testing.T) {
	ctx := context.Background()
	t.Logf("config:%+v", utils.Sdump(utils.GetConfInstance()))
	rdbfd := rediscache.GetRedis()
	cat := Cat{
		Name:  "xiaohua",
		Age:   3,
		Color: "red",
	}
	rediscache.SetKey(ctx, rdbfd, "cat", cat)
	ncats := new(Cat)
	var ll int64
	err := rediscache.GetKey(ctx, rdbfd, "ccc", &ll)
	if err != nil {
		t.Logf("err=%v", err.Error())
	}
	rediscache.GetKey(ctx, rdbfd, "cat", ncats)
	if *ncats != cat {
		t.Errorf("cat Set or get Obj failed, ncat:%v", utils.Sdump(ncats))
	}
	rediscache.DelKey(ctx, rdbfd, "cat")

}

func TestMain(m *testing.M) {
	utils.ConfigInit("../../../config.yaml")
	rediscache.InitRedisPool()
	m.Run()

}
