package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

var (
	limiterMap = make(map[string]*rate.Limiter)
	mu         sync.Mutex
)

// 限制访问频率，防止恶意刷接口

func RateLimit(rps float64, burst int) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//获取客户端IP
		clientIP := c.ClientIP()
		mu.Lock()
		limiter, exists := limiterMap[clientIP]
		if !exists {
			// 创建一个新的速率限制器，限制每秒rps次请求，允许burst次突发请求
			limiter = rate.NewLimiter(rate.Limit(rps), burst)
			limiterMap[clientIP] = limiter
		}
		mu.Unlock()

		// 检查速率
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, utils.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests, please wait a moment",
			})
			c.Abort()
			return
		}
	}
}
