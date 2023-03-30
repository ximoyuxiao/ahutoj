package main

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	pcodeforece "ahutoj/web/service/originproblem/p_codeforece"
	"ahutoj/web/utils"
	"fmt"
	"os"
	"sync"
)

var Wg sync.WaitGroup

func main() {
	config := "./config.yaml"
	if len(os.Args) >= 2 {
		config = os.Args[1]
	}
	err := utils.ConfigInit(config)
	if err != nil {
		fmt.Printf("call ConfigInit failed, config:%v\n", config)
	}
	utils.LogInit()
	err = mysqldao.InitMysql(nil)
	if err != nil {
		fmt.Printf("call ConfigInit failed, config:%v\n", utils.GetConfInstance().MySQLConfig)
	}
	Wg.Add(1)
	go func() {
		pcodeforece.DownLoadAllCodeForceProblem()
		Wg.Done()
	}()
	Wg.Wait()
	fmt.Println("hello world")
}
