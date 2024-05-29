package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//initiate fiber
	app := fiber.New()
	apiv1 := app.Group("/api/v1")


	//create foo route and attach function
	app.Get("/foo", handleFoo)
	apiv1.Get("/user", handleUser)
	//define port
	app.Listen(":5000")
}

//handle functin to pass message
func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user" : "James Foo"})
}