package models

type Books struct {
	Author string `json:"author" bson:"author"`
	Id string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Name string `json:"name" bson:"name"`
	ISBN string `json:"isbn" bson:"isbn"`
	
}