package customer

import (
	"errors"
	"time"

	pkg "github.com/obegendi/go-ddd/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastname"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Status    RegistrationStatus `bson:"status"`
	Birthday  time.Time          `bson:"birthday"`
}

func (c *Customer) IsLegal() bool {
	return c.Birthday.AddDate(18, 0, 0).Before(time.Now())
}

type RegistrationStatus struct {
	Status string `bson:"status"`
}

type status int

const (
	pending status = iota
	approved
	rejected
)

func (s status) String() string {
	switch s {
	case pending:
		return "pending"
	case approved:
		return "approved"
	case rejected:
		return "rejected"
	default:
		return "unknown"
	}
}

func (c *Customer) GetId() string {
	return c.Id.Hex()
}

func NewCustomer(firstName string, lastName string, email string, phone string) (customer *Customer, err error) {

	if firstName == "" || lastName == "" || email == "" || phone == "" {
		return nil, errors.New("Customer")
	}

	return &Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Status:    RegistrationStatus{Status: pending.String()},
	}, nil
}

func (c *Customer) UpdateStatus(status string) error {
	c.Status.Status = status
	return nil
}

type CustomerEvent struct {
	pkg.Event
	CustomerId string
}
