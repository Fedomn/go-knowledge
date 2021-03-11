package limit_test

import (
	. "github.com/fedomn/go-knowledge/test/gin/limit"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"testing"
	"time"
)

// SortedSet命令 http://redisdoc.com/sorted_set/index.html
// ZRANGE 127.0.0.1:test 0 -1 WITHSCORES
// ZREMRANGEBYSCORE 127.0.0.1:test 0 1

func TestRedisLimiter(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	slidingWindow := time.Second

	r.Use(NewRateLimiterMiddleware(redis.NewClient(
		&redis.Options{
			DB:       0,
			Password: "",
			Addr:     "127.0.0.1:6379",
		},
	), "test", 20, slidingWindow))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	go r.Run(":9999")

	for i := 0; i < 100; i++ {
		c := &http.Client{}

		resp, e := c.Get("http://127.0.0.1:9999")
		if e != nil {
			t.Error("Error during requests ", e.Error())
			return
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			//fmt.Printf("too many requests: %v\n", resp.StatusCode)
		}

		if i == 50 {
			// 等待目的：为了从now开始 向前推slidingWindow 窗口内的请求数减少，以便可以继续发请求
			time.Sleep(time.Second * 2)
		}
	}
}
