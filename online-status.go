package OnlineStatus

import (
	"time"
	"github.com/minio/lsync"
)

type BoolSlice []bool

type OnlineStatus struct {
	states *lsync.LFrequentAccess
}

const onlineHealthCheckDelay = 250 * time.Millisecond

// NewOnlineStatus
func NewOnlineStatus(dim int, x interface{}, healthCheck func(index int) bool) *OnlineStatus {
	online := &OnlineStatus{states: lsync.NewLFrequentAccess(x)}

	// Start one monitoring routine per node or disk to check its health periodically
	for index := 0; index < dim; index++ {
		go online.HealthMonitorLoop(index, healthCheck)
	}

	return online
}

// IsOnline returns the online or offline status of a node or disk.
func (online *OnlineStatus) IsOnline(index int) bool {
	return online.states.ReadOnlyAccess().(BoolSlice)[index]
}

// HealthMonitorLoop
func (online *OnlineStatus) HealthMonitorLoop(index int, healthCheck func(index int) bool) {
	for {
		time.Sleep(onlineHealthCheckDelay)

		// Get the health of the node or disk
		state := healthCheck(index)
		if state != online.IsOnline(index) {
			online.modify(index, state)
		}
	}
}

// modify changes the online or offline state for one of the nodes or disks.
// It uses the copy-on-write paradigm of LFrequentAccess.
func (online *OnlineStatus) modify(index int, state bool) {
	current := online.states.LockBeforeSet().(BoolSlice)
	update := make(BoolSlice, len(current))
	copy(update, current)
	update[index] = state
	online.states.SetNewCopyAndUnlock(update)
}
