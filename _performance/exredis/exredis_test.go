package exredis

import (
	"testing"

	"github.com/melodiez14/repokecil/util/conn"
)

func BenchmarkDo(b *testing.B) {
	pool := conn.InitRedis(conn.RedisConfig{
		ConnectionString: "devel-redis.tkpd:6379",
		MaxActive:        10,
		MaxIdle:          3,
		IdleTimeout:      10,
		Wait:             true,
	})

	b.ResetTimer()

	c := pool.Get()
	for i := 0; i < b.N; i++ {
		c.Do("PING")
	}
}

func BenchmarkSend(b *testing.B) {
	pool := conn.InitRedis(conn.RedisConfig{
		ConnectionString: "devel-redis.tkpd:6379",
		MaxActive:        10,
		MaxIdle:          3,
		IdleTimeout:      10,
		Wait:             true,
	})

	b.ResetTimer()

	c := pool.Get()
	for i := 0; i < b.N; i++ {
		c.Send("PING")
	}
}
