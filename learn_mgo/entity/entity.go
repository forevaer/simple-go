package entity

type Person struct {
	ID     string `bson:"_id"`
	Name   string `bson:"name"`
	Age    string `bson:"age"`
	Gender string `bson:"gender"`
}
