package types

import (
	"net"
	"sync"
)

// IService
type IService interface {
	Name() string
	IsRunning() bool
	Register(listener *net.Listener)
	Go(waitGroup *sync.WaitGroup)
}
