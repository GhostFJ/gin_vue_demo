package middleware

import (
	"fmt"
	"ginvue/response"

	"github.com/gin-gonic/gin"
)

// 拦截panic的错误信息，并返回的中间件
// 自定义recovery() 函数的实现
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(c, nil, fmt.Sprint(err))
				return
			}
		}()
		c.Next()
	}
}
