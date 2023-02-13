package models

type Library struct {
	ID string `json:"id" bson:"_id"` 
	Address string `json:"address" bson:"address"`
	Name string `json:"name" bson:"name"`
	Books []Books `json:"books" bson:"books"`
}