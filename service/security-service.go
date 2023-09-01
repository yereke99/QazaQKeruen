package service

import (
	"qkeruen/models"
	"qkeruen/repository"
	"time"
)

type SecurityService interface {
	Create(data models.Security) error
	Finish() *models.Security
}

type securityService struct {
	db repository.SecurityDB
}

func NewSecurityService(db repository.SecurityDB) *securityService {
	service := &securityService{
		db: db,
	}

	return service
}

func (s *securityService) Create(data models.Security) error {
	currentTime := time.Now()
	timeNow := currentTime.Format("2006-01-02 15:04:05")
	data.TimeStart = timeNow

	if err := s.db.Insert(data); err != nil {
		return err
	}

	return nil
}

func (s *securityService) Finish() *models.Security {
	var security models.Security

	return &security
}
