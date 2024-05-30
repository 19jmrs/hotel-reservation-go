package api

import (
	"github.com/19jmrs/hotel-reservation-go/db"
	"github.com/19jmrs/hotel-reservation-go/types"
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

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := c.BodyParser(&params); err != nil{
		return err
	}
	if err := params.Validate(); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	
	if err != nil{
		return err
	}

	return c.JSON(insertedUser)

}

//get user by id
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error{
	var(
		id = c.Params("id")		
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)

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

