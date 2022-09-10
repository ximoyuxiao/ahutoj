package middlewares

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/models"
	"ahutoj/web/utils"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type VerfiyLevel uint8

var (
	sign    []byte
	ExpTime time.Duration
)

const (
	UNLOGIN         VerfiyLevel = 0
	CommomUser      VerfiyLevel = 1
	Problem_edit    VerfiyLevel = 2
	Source_browser  VerfiyLevel = 3
	Contest_creator VerfiyLevel = 4
	Administrator   VerfiyLevel = 0xff
)

var verifyMap = map[string]VerfiyLevel{
	"/api/logout/":                  CommomUser,
	"/api/user/edit/":               CommomUser,
	"/api/user/edit/pass/":          CommomUser,
	"/api/user/vjudgeBind":          CommomUser,
	"/api/admin/permission/edit/":   Administrator,
	"/api/admin/permission/delete/": Administrator,
	"/api/admin/permission/add/":    Administrator,
	"/api/admin/permission/list/":   Administrator,
	"/api/problem/add/":             Problem_edit,
	"/api/problem/edit/":            Problem_edit,
	"/api/problem/delete/":          Problem_edit,
	"/api/contest/add/ ":            Contest_creator,
	"/api/contest/edit/":            Contest_creator,
	"/api/contest/delete/":          Contest_creator,
	"/api/file/add/:pid":            Problem_edit,
	"/api/file/delete/:pid":         Problem_edit,
	"/api/file/unzip/:pid":          Problem_edit,
	"/api/submit/rejudge/":          Administrator,
	"/api/submit/commit/":           CommomUser,
	"/api/training/add/":            Administrator,
	"/api/training/edit/":           Administrator,
	"/api/training/delete/":         Administrator,
}

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
	ExpTime = 3 * 24 * time.Hour
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
	// 获得当前的url
	url := c.FullPath()
	// 判断该url需要的权限等级，如果为不需要登录则不需要token信息
	if GetVerifyUrl(url) <= UNLOGIN {
		c.Next()
		return
	}

	/*至少需要UNLOGIN等级才需要通过Token鉴权*/
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
	/*判断是否拥有访问权限*/
	if HasUrlVerify(c, url, claims) {
		c.Set(JwtTokenCtxKey, claims)
		c.Next()
		return
	}
	response.ResponseError(c, constanct.VerifyErrorCode)
	logger.Infof("the Url(%s) need Verify", url)
	c.Abort()
}

func HasUrlVerify(ctx *gin.Context, url string, claims *MyClaims) bool {
	/*此处需要大于普通用户的权限*/
	if GetVerifyUrl(url) > CommomUser {
		permission, err := models.GetPermission(ctx, claims.UserID)
		/*此处说明用户没有权限*/
		if err != nil {
			return false
		}
		switch GetVerifyUrl(url) {
		case Problem_edit:
			{
				return permission.Administrator == "Y" || permission.Problem_edit == "Y"
			}
		case Source_browser:
			{
				return permission.Administrator == "Y" || permission.Source_browser == "Y"
			}
		case Contest_creator:
			{
				return permission.Administrator == "Y" || permission.Contest_creator == "Y"
			}
		case Administrator:
			{
				return permission.Administrator == "Y"
			}
		default:
			return true
		}
	}
	return true
}

func GetVerifyUrl(url string) VerfiyLevel {
	level, ok := verifyMap[url]
	if !ok {
		return UNLOGIN
	}
	return level
}
