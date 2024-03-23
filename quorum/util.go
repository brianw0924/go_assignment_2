package quorum

import (
	"math/rand"
	"time"
)

func HeartBeatInterval() <-chan time.Time {
	return time.After(time.Millisecond * time.Duration(30))
}

func ElectionTimeout() <-chan time.Time {
	return time.After(time.Millisecond * time.Duration(150+rand.Intn(151)))
}
