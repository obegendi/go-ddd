package customer

import (
	"errors"

	pkg "github.com/obegendi/go-ddd/pkg"
)

type Service interface {
	Register(customer *Customer) error
}

type service struct {
	repo      Repository
	publisher pkg.EventPublisher
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(customer *Customer) error {
	if !customer.IsLegal() {
		return errors.New("Customer is not legal")
	}

	err := s.repo.Create(customer)

	if err != nil {
		return err
	}

	s.publisher.Notify(CustomerEvent{
		CustomerId: customer.GetId(),
	})
	return nil
}
