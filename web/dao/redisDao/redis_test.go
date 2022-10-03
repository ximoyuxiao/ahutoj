package redisdao_test

import (
	redisdao "ahutoj/web/dao/redisDao"
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
	rdbfd := redisdao.GetRedis()
	if rdbfd == -1 {
		t.Logf("Test connect Redis pool filed")
	}
	redisdao.SetKey(ctx, rdbfd, "str", "hello world")
	str := new(string)
	redisdao.GetKey(ctx, rdbfd, "str", str)
	if *str != "hello world" {
		t.Errorf("Test TestSetString failed,str=%s", *str)
	}
	redisdao.DelKey(ctx, rdbfd, "str")
}
func TestSetObj(t *testing.T) {
	ctx := context.Background()
	t.Logf("config:%+v", utils.Sdump(utils.GetConfInstance()))
	rdbfd := redisdao.GetRedis()
	cat := Cat{
		Name:  "xiaohua",
		Age:   3,
		Color: "red",
	}
	redisdao.SetKey(ctx, rdbfd, "cat", cat)
	ncats := new(Cat)
	var ll int64
	err := redisdao.GetKey(ctx, rdbfd, "ccc", &ll)
	if err != nil {
		t.Logf("err=%v", err.Error())
	}
	redisdao.GetKey(ctx, rdbfd, "cat", ncats)
	if *ncats != cat {
		t.Errorf("cat Set or get Obj failed, ncat:%v", utils.Sdump(ncats))
	}
	redisdao.DelKey(ctx, rdbfd, "cat")

}

func TestMain(m *testing.M) {
	utils.ConfigInit("../../../config.yaml")
	redisdao.InitRedisPool()
	m.Run()

}
