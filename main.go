package main

import (
	"flag"

	"github.com/19jmrs/hotel-reservation-go/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The list address of the API server")
	flag.Parse()
	
	//initiate fiber
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)

	apiv1.Get("/user/:id", api.HandleGetUser)


	//define port
	app.Listen(*listenAddr)
}
