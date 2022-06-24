package command

import (
	"context"
	"log"

	"github.com/obegendi/go-ddd/pkg/customer"
)

type RegisterCustomer struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RegisterCustomerHandler struct {
	repo customer.Repository
}

func NewRegisterCustomerHandler(repo customer.Repository) RegisterCustomerHandler {
	return RegisterCustomerHandler{repo: repo}
}

func (h *RegisterCustomerHandler) Handle(ctx context.Context, cmd *RegisterCustomer) error {
	defer func() {
		log.Printf("RegisterCustomerHandler.Handle: %+v", cmd)
	}()

	customer, err := customer.NewCustomer(cmd.FirstName, cmd.LastName, cmd.Email, cmd.Phone)
	if err != nil {
		return err
	}

	err = h.repo.Create(customer)
	if err != nil {
		return err
	}

	return nil
}
