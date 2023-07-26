package jwtx

import (
	"errors"
	"fmt"
	"mixbench/common/utils"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
	// uuid "github.com/satori/go.uuid"
)

var (
	JwtUidKey      = "uid" //不要使用id，因为无法在ctx中获取id值
	JwtIpKey       = "ip"
	JwtPlatformKey = "platform"
)

var ReqHeaderTokenKey = "Authorization"
var DefaultDuration int64 = 86400 //一天时间

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("invalid token")
	ErrNotExist         = errors.New("not exist")
)

func NewClaims() jwt.MapClaims {
	return make(jwt.MapClaims)
}

func CreateToken(claims jwt.MapClaims, secret string) (string, error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	return token.SignedString([]byte(secret))
}

// 解析token, 返回原生error
func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	if secret = strings.TrimSpace(secret); secret == "" {
		secret = os.Getenv("JWT_SECRET")
		if secret = strings.TrimSpace(secret); secret == "" {
			return nil, errors.New("the environment variable JWT_SECRET equal nil")
		}
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		// 返回用于验证签名的密钥
		return []byte(secret), nil // 替换为实际的密钥
	})

	if err != nil {
		//Token 解析失败
		return nil, err
	}

	if token == nil {
		return nil, ErrTokenInvalid
	}

	// 验证 token 是否有效
	if !token.Valid {
		//无效的 Token
		return nil, ErrTokenInvalid
	}

	// 提取 claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrTokenInvalid
	}
	return claims, nil
}

func GetToken(r *http.Request) string {
	return r.Header.Get(ReqHeaderTokenKey)
}

// GetClaims 从请求中获取claims
func GetClaims(r *http.Request, secret string) (jwt.MapClaims, error) {
	token := GetToken(r)
	if token == "" {
		return nil, errors.New("token not exist")
	}
	claims, err := ParseToken(token, secret)
	if err != nil {
		err = fmt.Errorf("从request中获取从jwt解析信息失败, 请检查请求头是否存在%s且claims是否为规定结构,error:%s", ReqHeaderTokenKey, err.Error())
		logx.Error(err)
		return nil, err
	}
	return claims, nil
}

// GetClaimsValue 从token对象里获得参数(key)对应的值
// 示例 GetClaimsValue("username", token.claims) 其中token是已经解密的token
func GetClaimsValue(key string, claims jwt.MapClaims) string {
	v := reflect.ValueOf(claims)
	//确保是map类型
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}

// 将claims转化为map[string]string类型
func ClaimsToMap(claims jwt.MapClaims) map[string]string {
	res := make(map[string]string)
	v := reflect.ValueOf(claims)
	//确保是map类型
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			res[fmt.Sprintf("%s", k.Interface())] = fmt.Sprintf("%v", value.Interface())
		}
	}
	return res
}

func ExpiredTime(claims jwt.MapClaims) int64 {
	s := GetClaimsValue("exp", claims)
	if len(s) == 0 {
		return 0
	}
	res, err := utils.ExpToNumber(s)
	if err != nil {
		return 0
	}
	return res
}

func ValidateToken(token string, secret string) (jwt.MapClaims, bool) {
	claims, err := ParseToken(token, secret)
	if err != nil {
		return nil, false
	}
	return claims, !IsExpired(claims)
}

func IsExpired(claims jwt.MapClaims) bool {
	// 获取当前时间
	currentTime := time.Now().Local()

	// 判断是否已过期
	if exp, ok := claims["exp"].(float64); ok {
		expirationTime := time.Unix(int64(exp), 0).Local()

		if currentTime.After(expirationTime) {
			//Token 已过期
			return true
		} else {
			//Token 未过期
			return false
		}
	} else {
		//Token 没有设置过期时间，不会过期
		return false
	}
}

func GetId(claims jwt.MapClaims) (int64, bool) {
	// 获取 id 值
	id, ok := claims[JwtUidKey].(float64)
	if !ok {
		// id 不存在或类型不正确
		return 0, false
	}
	return int64(id), true
}

func GetInt64(claims jwt.MapClaims, key string) (res int64, exist bool) {
	v, ok := claims[key].(float64)
	exist = ok
	if !ok {
		//不存在或类型不正确
		return 0, false
	}
	res = int64(v)
	return res, exist
}

func GetString(claims jwt.MapClaims, key string) (res string, exist bool) {
	res, exist = claims[key].(string)
	return res, exist
}

func GetIdByRequest(r *http.Request, secret string) (int64, error) {
	if secret = strings.TrimSpace(secret); secret == "" {
		secret = os.Getenv("JWT_SECRET")
		if secret = strings.TrimSpace(secret); secret == "" {
			return 0, errors.New("the environment variable JWT_SECRET equal nil")
		}
	}
	claims, err := GetClaims(r, secret)
	if err != nil {
		return 0, err
	}
	id, ok := GetId(claims)
	if !ok {
		return 0, errors.New("ID does not exist in token")
	}
	return id, nil
}

func GetIdByToken(token string, secret string) (int64, error) {
	if token == "" {
		return 0, errors.New("token not exist")
	}
	if secret = strings.TrimSpace(secret); secret == "" {
		secret := os.Getenv("JWT_SECRET")
		if secret = strings.TrimSpace(secret); secret == "" {
			return 0, errors.New("the environment variable JWT_SECRET equal nil")
		}
	}
	claims, err := ParseToken(token, secret)
	if err != nil {
		logx.Error(err)
		return 0, err
	}
	id, ok := GetId(claims)
	if !ok {
		return 0, errors.New("ID does not exist in token")
	}
	return id, nil
}
