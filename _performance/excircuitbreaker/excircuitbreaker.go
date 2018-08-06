package excircuitbreaker

import (
	"time"
)

func getUserInfo() {
	time.Sleep(10 * time.Millisecond)
}
