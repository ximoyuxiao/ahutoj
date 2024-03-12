package utils
import (
	"testing"
)

func TestSendmail(t *testing.T) {
	err:=EmailVerify("test","2648242688@qq.com","test","test")
	if err!=nil{
		t.Errorf("send mail error,err:%v",err)
	}
}