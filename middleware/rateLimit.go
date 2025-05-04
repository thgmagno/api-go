package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RateLimiter(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rl:" + ip

		count, err := rdb.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Erro no rate limiter"})
			return
		}

		if count >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Limite de requisições excedido"})
			return
		}

		pipe := rdb.TxPipeline()
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, window)
		_, _ = pipe.Exec(ctx)

		c.Next()
	}
}
