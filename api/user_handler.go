package api

import (
	"context"

	"github.com/19jmrs/hotel-reservation-go/db"
	"github.com/gofiber/fiber/v2"
)

//everything that implements the user store can be used here
type UserHandler struct{
	userStore db.UserStore
}

//kind of constructor
func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

//get user by id
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error{
	var(
		id = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}
//to make function public to other packages by uppercasing the first letter.
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error{
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}
