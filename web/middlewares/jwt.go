package middlewares

import (
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"
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

// map所需要的的权限 第x为为1表示  如果用户权限为第X位则具备权限
const (
	UNLOGIN       VerfiyLevel = 0xff // 1111 1111
	CommomUser    VerfiyLevel = 0xfe // 1111 1110
	Administrator VerfiyLevel = 0xfc // 1111 1100
	ProblemAdmin  VerfiyLevel = 0x88 // 1000 1000
	ContestAdmin  VerfiyLevel = 0x90 // 1001 0000
	SourceBorwser VerfiyLevel = 0xa0 // 1010 0000
	ListAdmin     VerfiyLevel = 0xc0 // 1100 0000
	SuperAdmin    VerfiyLevel = 0x80 // 1000 0000
)

// 需要的权限等级
var verifyMap = map[string]VerfiyLevel{
	"/api/logout/":                  CommomUser,
	"/api/user/edit/":               CommomUser,
	"/api/user/edit/pass/":          CommomUser,
	"/api/user/vjudgeBind":          CommomUser,
	"/api/admin/permission/edit/":   SuperAdmin,
	"/api/admin/permission/delete/": SuperAdmin,
	"/api/admin/permission/add/":    SuperAdmin,
	"/api/admin/permission/list/":   SuperAdmin,
	"/api/admin/permission/:id":     CommomUser,
	"/api/problem/add/":             ProblemAdmin,
	"/api/problem/edit/":            ProblemAdmin,
	"/api/problem/delete/":          ProblemAdmin,
	"/api/contest/add/ ":            ContestAdmin,
	"/api/contest/edit/":            ContestAdmin,
	"/api/contest/delete/":          ContestAdmin,
	"/api/file/add/:pid":            CommomUser,
	"/api/file/delete/:pid":         ProblemAdmin,
	"/api/file/unzip/:pid":          ProblemAdmin,
	"/api/submit/rejudge/":          SuperAdmin,
	"/api/submit/commit/":           CommomUser,
	"/api/training/add/":            ListAdmin,
	"/api/training/edit/":           ListAdmin,
	"/api/training/delete/":         ListAdmin,
	"/api/admin/users":              SuperAdmin,
	"/api/submit/:id":               CommomUser,
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
func CheckUserHasPermission(ctx *gin.Context, needVerfiyLevel VerfiyLevel) bool {
	return VerfiyLevel(GetAdmin(ctx))&needVerfiyLevel != 0
}

// 验证token
func JwtVerify(c *gin.Context) {
	logger := utils.GetLogInstance()
	// 获得当前的url
	url := c.FullPath()

	token := c.GetHeader("Authorization")
	if token == "" {
		if GetVerifyUrl(url) == UNLOGIN {
			c.Next()
			return
		}
		logger.Errorf("token is empty")
		response.ResponseError(c, constanct.AUTH_Token_EmptyCode)
		c.Abort()
		return
	}
	// 验证token，并存储在请求中
	claims, err := ParseToken(token)
	if err != nil {
		logger.Errorf("token parse error, token=%s, err = %s", token, err.Error())
		response.ResponseError(c, constanct.AUTH_Token_InvalidCode)
		c.Abort()
		return
	}
	/*判断是否拥有访问权限*/
	if HasUrlVerify(c, url, claims) {
		c.Set(JwtTokenCtxKey, claims)
		c.Next()
		return
	}
	response.ResponseError(c, constanct.AUTH_Token_URLVerifyCode)
	logger.Infof("the Url(%s) need Verify", url)
	c.Abort()
}

func HasUrlVerify(ctx *gin.Context, url string, claims *MyClaims) bool {
	/*此处需要大于普通用户的权限*/
	UserPermission := mapping.UNLOGINBit
	NeedPermission := GetVerifyUrl(url)
	if claims != nil {
		mapping.AddPermissionBit(&UserPermission, mapping.CommomUserBit)
		permission, err := models.GetPermission(ctx, claims.UserID)
		if err == nil {
			mapping.AddPermissionBit(&UserPermission, mapping.PermissionBit(mapping.PermissionToBitMap(permission)))
		}
	}
	return (NeedPermission & VerfiyLevel(UserPermission)) != 0
}

func GetVerifyUrl(url string) VerfiyLevel {
	level, ok := verifyMap[url]
	if !ok {
		return UNLOGIN
	}
	return level
}
