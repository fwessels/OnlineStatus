package OnlineStatus

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestOnlineStatus(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	state := make(BoolSlice, 16)
	for i := range state {
		state[i] = true
	}

	online := NewOnlineStatus(len(state), state, HealthCheck)

	for i := 0; i < 100; i++ {

		for index := 0; index < len(state); index++ {
			if !online.IsOnline(index) {
				fmt.Println(index, "is offline")
			}
		}

		time.Sleep(1500 * time.Millisecond)
		fmt.Println("---------------")
	}
}