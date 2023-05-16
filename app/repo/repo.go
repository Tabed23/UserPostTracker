package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/Tabed23/UserPostTracker/app/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepo struct {
	db *mongo.Collection
}

func NewPostRepo(db *mongo.Collection) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(ctx context.Context, post *types.Post, usrId string) (string, error) {
	post.ID = primitive.NewObjectID()
	post.UserID, _ = primitive.ObjectIDFromHex(usrId)
	post.Comments = []types.Comment{}

	_, err := r.db.InsertOne(ctx, post)
	if err != nil {
		return "", errors.Wrap(err, "failed to create post")
	}

	filter := bson.M{"_id": post.ID}
	update := bson.M{
		"$set": bson.M{"userId": usrId},
	}
	_, err = r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return post.ID.Hex(), nil
}

func (r *PostRepo) GetPostByID(ctx context.Context, id string) (*types.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse post ID")
	}

	var post types.Post
	err = r.db.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "failed to get post")
	}
	return &post, nil
}

func (r *PostRepo) GetPosts(ctx context.Context, skip, limit int64) ([]*types.Post, error) {
	matchStage := bson.D{{"$match", bson.D{}}}
	sortStage := bson.D{{"$sort", bson.D{{"createdAt", -1}}}}
	limitStage := bson.D{{"$limit", limit}}

	pipeline := mongo.Pipeline{matchStage, sortStage, limitStage}

	cursor, err := r.db.Aggregate(ctx, pipeline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(ctx)

	posts := []*types.Post{}
	for cursor.Next(ctx) {
		post := &types.Post{}
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepo) UpdatePost(ctx context.Context, post *types.Post) (string, error) {
	filter := bson.M{"_id": post.ID}
	update := bson.M{"$set": bson.M{
		"title":     post.Title,
		"body":      post.Body,
		"media":     post.Media,
		"updatedAt": time.Now(),
	}}
	_, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return "post updated", nil
}
func (r *PostRepo) LikePost(ctx context.Context, postID string, userID string) error {

	postObjID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}
	fmt.Println(userID)

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjID}
	update := bson.M{
		"$addToSet": bson.M{"likes": userObjID},
		"$inc":      bson.M{"likeCount": 1},
		"$set":      bson.M{"userId": userObjID},
	}

	_, err = r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
func (r *PostRepo) UnlikePost(ctx context.Context, postID string, userID string) error {
	postObjID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": postObjID, "likes": userID}
	update := bson.M{
		"$pull": bson.M{"likes": userID},
		"$inc":  bson.M{"likeCount": -1},
	}

	_, err = r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil

}
func (r *PostRepo) AddComment(ctx context.Context, postID string, userID string, comment *types.Comment) (*types.Post, error) {
	postObjID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": postObjID}
	update := bson.M{
		"$push": bson.M{"comments": comment},
		"$inc":  bson.M{"commentCount": 1},
	}
	_, err = r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	updatedPost := &types.Post{}
	err = r.db.FindOne(ctx, bson.M{"_id": postObjID}).Decode(updatedPost)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (r *PostRepo) GetPostCount(ctx context.Context) (int64, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return count, nil
}

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Collection) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *types.User) (string, error) {
	user.ID = primitive.NewObjectID()

	_, err := u.db.InsertOne(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "failed to create user")
	}

	return user.ID.Hex(), nil
}
