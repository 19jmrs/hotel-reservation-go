package types

//capital letter initially to make it public
//same for the entities of the user
//bson and json notation is space sensitivy
type User struct{
	ID  string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
}
