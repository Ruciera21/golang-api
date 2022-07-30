package service

import (
	"goapi-nunu/domain"
	"goapi-nunu/dto"
	"goapi-nunu/errors"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errors.AppError)
	GetCustomerByID(string) (*dto.CustomerResponse, *errors.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errors.AppError) {
	// * add process here
	custs, err := s.repository.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response []dto.CustomerResponse
	for _, c := range custs {
		response = append(response, c.ToDTO())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errors.AppError) {
	// * add process here
	cust, err := s.repository.FindByID(customerID)
	if err != nil {
		return nil, err
	}
	response := cust.ToDTO()

	return &response, nil
}
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
