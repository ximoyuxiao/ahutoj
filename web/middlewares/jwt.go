package middlewares

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/mapping"
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

// 需要的权限等级
var VerifyMap = map[string]constanct.VerfiyLevel{
	"/api/auth/logout/":             constanct.CommomUser,
	"/api/user/edit/":               constanct.CommomUser,
	"/api/user/edit/pass/":          constanct.CommomUser,
	"/api/user/vjudgeBind":          constanct.CommomUser,
	"/api/admin/permission/edit/":   constanct.SuperAdmin,
	"/api/admin/permission/delete/": constanct.SuperAdmin,
	"/api/admin/permission/add/":    constanct.SuperAdmin,
	"/api/admin/permission/list/":   constanct.SuperAdmin,
	"/api/admin/permission/:id":     constanct.CommomUser,
	"/api/problem/add/":             constanct.ProblemAdmin,
	"/api/problem/edit/":            constanct.ProblemAdmin,
	"/api/problem/delete/":          constanct.ProblemAdmin,
	"/api/contest/add/ ":            constanct.ContestAdmin,
	"/api/contest/edit/":            constanct.ContestAdmin,
	"/api/contest/delete/":          constanct.ContestAdmin,
	"/api/file/add/:pid":            constanct.CommomUser,
	"/api/file/delete/:pid":         constanct.ProblemAdmin,
	"/api/file/unzip/:pid":          constanct.ProblemAdmin,
	"/api/submit/rejudge/":          constanct.SuperAdmin,
	"/api/submit/commit/":           constanct.CommomUser,
	"/api/training/add/":            constanct.ListAdmin,
	"/api/training/edit/":           constanct.ListAdmin,
	"/api/training/delete/":         constanct.ListAdmin,
	"/api/training/user/":           constanct.CommomUser,
	"/api/admin/users":              constanct.SuperAdmin,
	"/api/submit/:id":               constanct.CommomUser,
	"/api/user/editHead/":           constanct.CommomUser,
	"/api/file/image/":              constanct.ProblemAdmin,
}

const JwtTokenCtxKey = "user"

type Permmision struct {
	PermissionMap int `json:"permission_map"`
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
	permission, err := mysqldao.SelectPermissionByUID(ctx, userID)
	if err != nil {
		return "", err
	}
	c := MyClaims{
		userID,
		Permmision{
			PermissionMap: mapping.PermissionToBitMap(permission),
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

func GetAdmin(ctx *gin.Context) mapping.PermissionBit {
	a, _ := ctx.Get(JwtTokenCtxKey)
	MyClaims, ok := a.(*MyClaims)
	if !ok {
		return mapping.UNLOGINBit
	}
	return mapping.PermissionBit(MyClaims.PermissionMap)
}

func CheckUserHasPermission(ctx *gin.Context, needVerfiyLevel constanct.VerfiyLevel) bool {
	return constanct.VerfiyLevel(GetAdmin(ctx))&needVerfiyLevel != 0
}

// 验证token
func JwtVerify(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return
	}
	// 验证token，并存储在请求中
	claims, _ := ParseToken(token)
	c.Set(JwtTokenCtxKey, claims)
	c.Next()
}

func GetVerifyUrl(url string) constanct.VerfiyLevel {
	level, ok := VerifyMap[url]
	if !ok {
		return constanct.UNLOGIN
	}
	return level
}
