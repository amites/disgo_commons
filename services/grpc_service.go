package services

import (
	"net"
	"strconv"
	"sync"

	"github.com/dispatchlabs/disgo/properties"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var grpcServerInstance *grpc.Server
var grpcServerOnce sync.Once

func GetGrpcServer() *grpc.Server {
	grpcServerOnce.Do(func() {
		grpcServerInstance = grpc.NewServer()
	})
	return grpcServerInstance
}

type GrpcService struct {
	Port    int
	running bool
}

func NewGrpcService() *GrpcService {
	return &GrpcService{properties.Properties.GrpcPort, false}
}

func (grpcService *GrpcService) Name() string {
	return "GrpcService"
}

func (grpcService *GrpcService) IsRunning() bool {
	return grpcService.running
}

func (grpcService *GrpcService) Go(waitGroup *sync.WaitGroup) {

	grpcService.running = true
	listener, error := net.Listen("tcp", ":"+strconv.Itoa(grpcService.Port))
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}

	// Serve.
	log.WithFields(log.Fields{
		"method": grpcService.Name() + ".Go",
	}).Info("listening on " + strconv.Itoa(grpcService.Port))
	reflection.Register(GetGrpcServer())
	if error := GetGrpcServer().Serve(listener); error != nil {
		log.Fatalf("failed to serve: %v", error)
		grpcService.running = false
	}
}

