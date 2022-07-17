package middlewares

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	sign    []byte
	ExpTime time.Duration
)

const JwtTokenCtxKey = "user"

type Permmision struct {
	Administrator   bool `json:"administrator"`
	Problem_edit    bool `json:"problem_edit"`
	Source_browser  bool `json:"source_browser"`
	Contest_creator bool `json:"contest_creator"`
}

type MyClaims struct {
	UserID string `json:"uid"`
	Permmision
	jwt.StandardClaims
}

func InitJwt() {
	config := utils.GetConfInstance()
	sign = []byte(config.Sign)
	ExpTime = 24 * time.Hour
}

func GetToken(ctx *gin.Context, userID string) (string, error) {
	// 创建一个我们自己的声明的数据
	permission, err := mysqldao.SelectPermissionByUid(ctx, userID)
	if err != nil {
		return "", err
	}
	c := MyClaims{
		userID,
		Permmision{
			Administrator:   permission.Administrator == "Y",
			Problem_edit:    permission.Problem_edit == "Y",
			Source_browser:  permission.Source_browser == "Y",
			Contest_creator: permission.Contest_creator == "Y",
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpTime).Unix(), // 过期时间
			Issuer:    "ahutoj",                       // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(sign)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var myclaims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myclaims, func(token *jwt.Token) (i interface{}, err error) {
		return sign, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return myclaims, nil
	}
	return nil, errors.New("invalid token")
}

// GetUid 从上下文对象中取出 uid
func GetUid(ctx *gin.Context) string {
	a, _ := ctx.Get(JwtTokenCtxKey)
	myClaims, ok := a.(*MyClaims)
	if !ok {
		return ""
	}
	return myClaims.UserID
}

// 验证token
func JwtVerify(c *gin.Context) {
	// 过滤是否验证token
	logger := utils.GetLogInstance()
	token := c.GetHeader("Authorization")
	if token == "" {
		logger.Errorf("token is empty")
		response.ResponseError(c, constanct.TokenInvalidCode)
		c.Abort()
		return
	}
	// 验证token，并存储在请求中
	claims, err := ParseToken(token)
	if err != nil {
		logger.Errorf("token parse error, token=%s, err = %s", token, err.Error())
		response.ResponseError(c, constanct.TokenInvalidCode)
		c.Abort()
	}

	c.Set(JwtTokenCtxKey, claims)
	c.Next()
}