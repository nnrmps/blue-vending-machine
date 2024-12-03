package service

type HealthService interface {
	HealthCheck() error
}

type healthService struct {
}

func NewHealthService() HealthService {
	return &healthService{}
}

func (h *healthService) HealthCheck() error {
	return nil
}
