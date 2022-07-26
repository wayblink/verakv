package file

import (
	"sync"
)

// WalFile _
type WalFile struct {
	lock    *sync.RWMutex
	size    uint32
	writeAt uint32
}
