package services

import (
	"net"
	"strconv"
	"sync"

	"github.com/dispatchlabs/disgo/party"
	"github.com/dispatchlabs/disgo/properties"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var instance *grpc.Server
var once sync.Once

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
	// Register disgoGrpc.
	log.WithFields(log.Fields{
		"method": grpcService.Name() + ".Go",
	}).Info("registering Disgover...")
	party.RegisterPartyServer(GetCGrpcServer(), party.NewParty())

	// Register Disgover.
	log.WithFields(log.Fields{
		"method": grpcService.Name() + ".Go",
	}).Info("registering Disgover...")
	//disgover.RegisterDisgoverRPCServer(GetCGrpcServer(), disgover.GetInstance())

	// Serve.
	reflection.Register(GetCGrpcServer())
	log.WithFields(log.Fields{
		"method": grpcService.Name() + ".Go",
	}).Info("listening on " + strconv.Itoa(grpcService.Port))
	if error := GetCGrpcServer().Serve(listener); error != nil {
		log.Fatalf("failed to serve: %v", error)
		grpcService.running = false
	}
}


func GetCGrpcServer() *grpc.Server {
	once.Do(func() {
		instance = grpc.NewServer()
	})
	return instance
}

