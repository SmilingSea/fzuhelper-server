package middleware

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	Identity      = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte("demo"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		TokenLookup:      "header:Authorization, query: token, cookie, jwt",
		TokenHeadName:    "Bearer",
		// 登录成功后的响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"message": "success",
			})
		},
		// 收到登录数据后的处理逻辑
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				id string `json:"id"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			id := loginStruct.id
			if !(id == "111") {
				return nil, errors.New("invalid username or password")
			}
			return id, nil
		},
		IdentityKey: Identity,
		//从token提取用户信息的函数
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			return nil
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}
