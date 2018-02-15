package types

import (
	"sync"
	"google.golang.org/grpc"
)

// IService
type IService interface {
	Name() string
	IsRunning() bool
	RegisterGrpc(grpcServer *grpc.Server)
	Go(waitGroup *sync.WaitGroup)
}
