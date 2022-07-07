package manager

import "golang_lopei_grpc_server/service"

type ServiceManager interface {
	LopeiService() *service.LopeiService
}

type serviceManager struct {
	lopeiService *service.LopeiService
}

func (s *serviceManager) LopeiService() *service.LopeiService{
	return s.lopeiService
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		lopeiService: service.NewLopeiService(repoManager.LopeiRepository()),
	}
}