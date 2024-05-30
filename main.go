package main

import (
	"context"
	"flag"
	"log"

	"github.com/19jmrs/hotel-reservation-go/api"
	"github.com/19jmrs/hotel-reservation-go/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"


var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error{
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {

	
	listenAddr := flag.String("listenAddr", ":5000", "The list address of the API server")
	flag.Parse()
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	//handler initiated
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	
	//initiate fiber
	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")


	apiv1.Get("/user", userHandler.HandleGetUsers)

	apiv1.Get("/user/:id", userHandler.HandleGetUser)


	//define port
	app.Listen(*listenAddr)
}
