package types

import (
	"sync"
)

// IService
type IService interface {
	Name() string
	IsRunning() bool
	Go(waitGroup *sync.WaitGroup)
}
