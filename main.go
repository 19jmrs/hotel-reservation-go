package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/19jmrs/hotel-reservation-go/api"
	"github.com/19jmrs/hotel-reservation-go/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"


func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	
	user := types.User{
		FirstName: "Jorge",
		LastName: "Sa",
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

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
