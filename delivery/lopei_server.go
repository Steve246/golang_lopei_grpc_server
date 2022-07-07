package delivery

import (
	"golang_lopei_grpc_server/config"
	"golang_lopei_grpc_server/manager"
	"golang_lopei_grpc_server/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LopeiGrpcServer struct {
	netListen net.Listener
	server *grpc.Server
	serviceManager manager.ServiceManager

}

func(lgs *LopeiGrpcServer) Run() {
	service.RegisterLopeiPaymentServer(lgs.server, lgs.serviceManager.LopeiService())

	log.Println("Server Run", lgs.netListen.Addr().String())

	err := lgs.server.Serve(lgs.netListen)

	if err != nil {
		log.Fatalln("Failed to Serve..", err)
	}
}


func Server() *LopeiGrpcServer{
	lopeiGrpcServer := new(LopeiGrpcServer)
	c := config.NewConfig()

	listen, err := net.Listen("tcp", c.Url)
	if err != nil {
		log.Fatalln("Failed to Listen", err)
	}

	grpcServer := grpc.NewServer()
	repoManager := manager.NewRepositoryManager()
	lopeiGrpcServer.serviceManager = manager.NewServiceManager(repoManager)

	lopeiGrpcServer.netListen = listen 
	lopeiGrpcServer.server = grpcServer
	return lopeiGrpcServer
}