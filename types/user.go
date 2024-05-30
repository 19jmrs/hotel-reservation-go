package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type CreateUserParams struct {
	FirstName string
	LastName string
	Email     string
	Password  string
}

// capital letter initially to make it public
// same for the entities of the user
// bson and json notation is space sensitivy
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}