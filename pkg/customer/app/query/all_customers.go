package query

import (
	"github.com/obegendi/go-ddd/pkg/customer"
)

type CustomersHandler struct {
	repo customer.Repository
}

type CustomerReadModel struct {
}

func NewCustomersHandler(repo customer.Repository) CustomersHandler {

	return CustomersHandler{repo: repo}
}

func (h *CustomersHandler) Handle(id string) (*customer.Customer, error) {
	customer, err := h.repo.Get(id)

	if err != nil {
		return nil, err
	}
	return customer, nil
}
