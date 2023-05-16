package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

type Post struct {
	ID        primitive.ObjectID   `json:"_id" bson:"_id"`
	UserID    primitive.ObjectID   `json:"userId" bson:"userId"`
	Title     string               `json:"title" bson:"title"`
	Body      string               `json:"body" bson:"body"`
	Media     string               `json:"media" bson:"media"`
	Likes     []primitive.ObjectID `json:"likes,omitempty" bson:"likes,omitempty"`
	Comments  []Comment            `json:"comments" bson:"comments"`
	CreatedAt primitive.DateTime   `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime   `json:"updatedAt" bson:"updatedAt"`
}

type Comment struct {
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Comment   string             `json:"comment" bson:"comment"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}

type Like struct {
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	PostID    primitive.ObjectID `json:"postId" bson:"postId"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
