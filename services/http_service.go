package services

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"github.com/dispatchlabs/disgo/properties"
	"github.com/gorilla/mux"
)

var httpRouterInstance *mux.Router
var httpRouterOnce sync.Once

func GetHttpRouter() *mux.Router {
	httpRouterOnce.Do(func() {
		httpRouterInstance =  mux.NewRouter()
	})
	return httpRouterInstance
}

// HttpService
type HttpService struct {
	HostIp  string
	Port    int
	running bool
}

// NewHttpService
func NewHttpService() *HttpService {
	return &HttpService{
		properties.Properties.HttpHostIp,
		properties.Properties.HttpPort,
		false}
}

// Name
func (httpService *HttpService) Name() string {
	return "HttpService"
}

// IsRunning
func (httpService *HttpService) IsRunning() bool {
	return httpService.running
}

// Go
func (httpService *HttpService) Go(waitGroup *sync.WaitGroup) {
	httpService.running = true
	listen := httpService.HostIp + ":" + strconv.Itoa(httpService.Port)
	log.WithFields(log.Fields{
		"method": httpService.Name() + ".Go",
	}).Info("listening on http://" + listen)
	log.Fatal(http.ListenAndServe(listen, GetHttpRouter()))
}

// setHeaders
func (httpService *HttpService) setHeaders(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")

	// TODO: Add headers for cross domain access.
	/*
			rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	 */

}
