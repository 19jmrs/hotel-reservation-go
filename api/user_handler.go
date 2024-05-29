package api

import (
	"github.com/19jmrs/hotel-reservation-go/types"
	"github.com/gofiber/fiber/v2"
)

//to make function public to other packages by uppercasing the first letter.
func HandleGetUsers(c *fiber.Ctx) error{
	u := types.User{
		FirstName: "James",
		LastName: "Drinkwater",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error{
	return c.JSON("James")
}