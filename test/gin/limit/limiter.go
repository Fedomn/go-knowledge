package limit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

// 总体参考：https://www.zybuluo.com/kay2/note/949160

// 1、限制瞬时并发数
// 防止一次性建立过多连接 消耗大量资源 反而性能下降

// copy from https://github.com/gin-gonic/contrib gin-limit
func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}

// 2、限制时间窗最大请求数
// 可能出现流量不平滑情况，窗口内一小段流量占比特别大

// copy from https://medium.com/@salvatoregiordanoo/a-step-approach-to-rate-limiting-f2190dff9fd4
// 请求都是过来了的 限流只是不进行业务逻辑处理直接返回
// 想象成一条管道 左进右出。窗口是以左边(now)开始，向右边走(slidingWindow)距离
func NewRateLimiterMiddleware(redisClient *redis.Client, key string, limit int64, slidingWindow time.Duration) gin.HandlerFunc {

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key)

		// 移除有序集 key 中，所有 score 值介于 min 和 max 之间的成员
		// 移除窗口以外的元素
		redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		// 按照score从低到高的顺序返回范围内的元素
		// 返回窗口内的元素
		reqCount, _ := redisClient.ZCount(userCntKey, "-inf", "+inf").Result()

		if reqCount >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status":  http.StatusTooManyRequests,
				"message": "too many request",
			})
			return
		}

		c.Next()

		// 只添加新元素，不会更新已有元素
		// 向窗口内插入元素
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})

		// 重新设置TTL
		// 这里并不会对限速产生影响，它的存在是为了 自动删除这个key
		redisClient.Expire(userCntKey, slidingWindow)
	}
}
