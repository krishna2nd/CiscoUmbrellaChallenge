package userver

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"uhandlers"
)

// UServer which handles the server base structure
type UServer struct {
	http.Server
	comChannel chan int
}

// New allows to create new uscan server
func New(com chan int) *UServer {
	return new(UServer).init(com)
}

// initializing server default configurations
func (server *UServer) init(com chan int) *UServer {
	svrConfig := config.Get()
	server.Addr = svrConfig.Port
	server.ReadTimeout = svrConfig.ReadTimeout
	server.WriteTimeout = svrConfig.WriteTimeout
	server.MaxHeaderBytes = svrConfig.MaxHeaderBytes
	server.comChannel = com
	server.addRouteMounts()
	return server
}

// Start starting the server
func (server *UServer) Start() {
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	fmt.Println("Server started in ", server.Addr)
}

// Stop the server
func (server *UServer) Stop() {
	server.comChannel <- 1
	fmt.Println("Server stopped")
}

// Add url mounts to server against the handlers
func (server *UServer) addRouteMounts() {
	http.HandleFunc("/", uhandlers.GlobalHandler)
}
