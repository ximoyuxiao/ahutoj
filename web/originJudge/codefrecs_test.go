package originjudge_test

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
	originjudge "ahutoj/web/originJudge"
	"context"
	"testing"
	"time"
)

func TestSubmitAndGetResult(t *testing.T) {
	cfJudge := originjudge.CodeForceJudge{}
	cfJudge.Submit = dao.Submit{
		Lang:       constanct.CPP11,
		SubmitTime: time.Now().UnixMilli(),
		Source: `#include<iostream>
		using namespace std;
		int main(){
			int a,b;
			cin>>a>>b;
			cout<<a+b<<endl;
		return 0;
		}`,
	}
	cfJudge.PID = "1003A"
	cfJudge.Judge(context.Background(), cfJudge.Submit, cfJudge.PID)
}

func TestAtcoderLogin(t *testing.T) {
	atcoderJudger := originjudge.AtCoderJudge{}
	atcoderJudger.Judge(context.Background(), dao.Submit{}, "1003A")
}
