package app

import (
	"github.com/obegendi/go-ddd/pkg/customer/app/command"
	"github.com/obegendi/go-ddd/pkg/customer/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	RegisterCustomerHandler command.RegisterCustomerHandler
}

type Queries struct {
	Customer query.CustomersHandler
}
