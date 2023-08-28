package service

import (
	"log"
	"qkeruen/dto"
	"qkeruen/models"
	"qkeruen/repository"
)

type OfferUserService interface {
	GetByID(id int64) (*models.DriverModelForUser, error)
	Create(id int, offer dto.OfferRequest) error
	MyOffer(id int64) ([]*dto.OfferResponseUser, error)
	FindAllOffers() ([]*dto.OfferResponseUser, error)
	Search(from, to, type_ string) ([]*models.OfferUserModel, error)
	DeleteOffer(offerId int64) error
}

type offerUserService struct {
	db repository.OfferUserDB
}

func NewOfferuserService(ds repository.OfferUserDB) *offerUserService {
	return &offerUserService{
		db: ds,
	}
}

func (s *offerUserService) GetByID(id int64) (*models.DriverModelForUser, error) {
	res, err := s.db.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *offerUserService) Create(id int, offer dto.OfferRequest) error {
	return s.db.Create(id, offer)
}

func (s *offerUserService) MyOffer(id int64) ([]*dto.OfferResponseUser, error) {
	res, err := s.db.MyOffer(id)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) FindAllOffers() ([]*dto.OfferResponseUser, error) {
	res, err := s.db.FindAllOffers()

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) Search(from, to, type_ string) ([]*models.OfferUserModel, error) {
	if to == "" {
		log.Println("working this point!")
		res, err := s.db.SearchA(from, type_)
		if err != nil {
			return nil, err
		}

		return res, err
	}
	res, err := s.db.Search(from, to, type_)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *offerUserService) DeleteOffer(offerId int64) error {
	return s.db.Delete(offerId)
}
