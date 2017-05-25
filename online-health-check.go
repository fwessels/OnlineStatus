package OnlineStatus

import (
	"math/rand"
)

// HealthCheck
//
// This function should do a health check on the designated item
// It will be called periodically
func HealthCheck(index int) bool {

	return rand.Int63n(3) != 0
}
