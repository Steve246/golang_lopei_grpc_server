package delivery

import (
	"golang_lopei_grpc_server/manager"
	"golang_lopei_grpc_server/service"
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
}