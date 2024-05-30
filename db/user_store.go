package db

import (
	"context"

	"github.com/19jmrs/hotel-reservation-go/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

//interface allows to implement different functions for it
type UserStore interface {
	GetUserByID (context.Context,string) (*types.User, error)
	GetUsers (context.Context) ([]*types.User, error)
}

//implementation of the user store
type MongoUserStore struct{
	client *mongo.Client
	coll *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore{
	
	return &MongoUserStore{
		client: client,
		coll: client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error){
	var user types.User
	if err := s.coll.FindOne(ctx, bson.M{"_id": ToObjectID(id)}).Decode(&user); err != nil{
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers (ctx context.Context) ([]*types.User, error){
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil,err
	}
	var users []*types.User

	if err := cur.All(ctx, &users); err != nil{
		return nil, err
	}

	return users, nil
}