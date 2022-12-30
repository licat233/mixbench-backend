package jwtx

import (
	"errors"
	"fmt"
	"mixbench/common/errorx"
	"mixbench/common/utils"
	"net/http"
	"reflect"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	// uuid "github.com/satori/go.uuid"
)

var ReqHeaderKey = "Authorization"
var DefaultDuration int64 = 86400 //一天时间

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func CreateToken(claims jwt.MapClaims, secret string) (string, error) {
	// 创建一个新的令牌对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密码签名并获得完整的编码令牌作为字符串
	return token.SignedString([]byte(secret))
}

// 解析token, 返回原生error
func ParseToken(tokenString string, secret string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	} else {
		return nil, ErrTokenInvalid
	}
}

func GetToken(r *http.Request) string {
	return r.Header.Get(ReqHeaderKey)
}

// GetClaims 从请求中获取claims
func GetClaims(r *http.Request, secret string) (*jwt.MapClaims, error) {
	token := GetToken(r)
	claims, err := ParseToken(token, secret)
	if err != nil {
		err = fmt.Errorf("从request中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构,error:%s", err.Error())
		logx.WithContext(r.Context()).Error(err)
		return nil, errorx.InternalError(err)
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
