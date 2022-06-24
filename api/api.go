package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/obegendi/go-ddd/api/routes"
	"github.com/obegendi/go-ddd/config"
	inf "github.com/obegendi/go-ddd/infrastructure"
	"github.com/obegendi/go-ddd/pkg/customer"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(config *config.Config) {

	db := new(config)

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// For invalid credentials
			//return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
			return nil
		}
	})

	customerRepo := customer.NewCustomerRepository(db.Collection("customers"))
	customerService := customer.NewService(customerRepo)

	routes.CustomerRouter(e, customerService)
	log.Fatalf("%v", e.Start(":8080"))
}

func new(config *config.Config) *mongo.Database {
	db, cancel, err := inf.Connect(config)
	defer cancel()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	return db

	// var cApp = customerApp.Application{
	// 	Commands: customerApp.Commands{
	// 		RegisterCustomerHandler: command.NewRegisterCustomerHandler(customerRepo),
	// 	},
	// 	Queries: customerApp.Queries{
	// 		Customer: query.NewCustomersHandler(customerRepo),
	// 	},
	// }
}
