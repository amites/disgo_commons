package types

import (
	"sync"
	"google.golang.org/grpc"
)

// IService
type IService interface {
	Init()
	Name() string
	IsRunning() bool
	RegisterGrpc(grpcServer *grpc.Server)
	Go(waitGroup *sync.WaitGroup)
}
