package service

import (
	"fmt"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/domain"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/port"
)

type saveLogService struct {
	logRepository port.LogRepository
}

func NewSaveLogService(logRepository port.LogRepository) port.SaveLogServicePort {
	return &saveLogService{
		logRepository: logRepository,
	}
}

func (s *saveLogService) Exec(logDomain *domain.LogDomain) error {
	err := s.logRepository.Save(logDomain)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
