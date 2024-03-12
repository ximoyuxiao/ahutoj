package parsejwt

import (
	"ahutoj/web/io/constanct"
	"ahutoj/web/io/response"
	"ahutoj/web/mapping"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var (
	sign    []byte
	ExpTime time.Duration
)

// 需要的权限等级
var VerifyMap = map[string]constanct.VerfiyLevel{}

const JwtTokenCtxKey = "user"

type Permmision struct {
	PermissionMap int `json:"permission_map"`
}

type MyClaims struct {
	UserID string `json:"uid"`
	Permmision
	jwt.StandardClaims
}

func InitJwt(Sign string) {
	sign = []byte(Sign)
	ExpTime = 3 * 24 * time.Hour
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
func JwtVerify(c *gin.Context) (interface{}, bool) {
	// 获得当前的url
	url := c.FullPath()
	token := c.GetHeader("Authorization")
	if token == "" {
		if GetVerifyUrl(url) == constanct.UNLOGIN {
			return nil, true
		}
		return response.CreateResponse(constanct.AUTH_Token_EmptyCode), false
	}
	// 验证token，并存储在请求中
	claims, err := ParseToken(token)
	if err != nil {
		return response.CreateResponse(constanct.AUTH_Token_InvalidCode), false
	}
	c.Set(JwtTokenCtxKey, claims)
	/*判断是否拥有访问权限*/
	if HasUrlVerify(c, url, claims) {
		return nil, true
	}
	return response.CreateResponse(constanct.AUTH_Token_URLVerifyCode), false
}

func HasUrlVerify(ctx *gin.Context, url string, claims *MyClaims) bool {
	/*此处需要大于普通用户的权限*/
	UserPermission := mapping.UNLOGINBit
	NeedPermission := GetVerifyUrl(url)
	// fmt.Println(utils.Sdump(claims))
	if claims != nil {
		mapping.AddPermissionBit(&UserPermission, mapping.CommomUserBit)
		mapping.AddPermissionBit(&UserPermission, mapping.PermissionBit(claims.PermissionMap))
	}
	// fmt.Printf("need permission:%v, UserPermission:%v\n", NeedPermission, UserPermission)
	return (NeedPermission & constanct.VerfiyLevel(UserPermission)) != 0
}

func GetVerifyUrl(url string) constanct.VerfiyLevel {
	level, ok := VerifyMap[url]
	if !ok {
		return constanct.UNLOGIN
	}
	return level
}
