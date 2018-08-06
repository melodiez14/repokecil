package excircuitbreaker

import (
	"testing"

	"github.com/afex/hystrix-go/hystrix"
)

func BenchmarkCB(b *testing.B) {
	cmd := map[string]hystrix.CommandConfig{
		"USER_INFO": hystrix.CommandConfig{
			ErrorPercentThreshold: 50,
			MaxConcurrentRequests: 10,
			Timeout:               10,
		},
	}
	hystrix.Configure(cmd)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go hystrix.Do("USER_INFO", func() error {
			getUserInfo()
			return nil
		}, func(err error) error {
			return err
		})
	}
}

func BenchmarkWithoutCB(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go getUserInfo()
	}
}
