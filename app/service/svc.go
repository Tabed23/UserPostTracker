package service

import (
	"context"

	"github.com/Tabed23/UserPostTracker/app/repo"
	"github.com/Tabed23/UserPostTracker/app/types"
)

type PostInterface interface {
	CreatePost(ctx context.Context, post *types.Post) error
	GetPostByID(ctx context.Context, id string) (*types.Post, error)
	GetPosts(ctx context.Context, skip, limit int64) ([]*types.Post, error)
	GetPostCount(ctx context.Context) (int64, error)
	AddComment(ctx context.Context, postID string, comment *types.Comment) error
	AddLike(ctx context.Context, postID, userID string) error
	RemoveLike(ctx context.Context, postID, userID string) error
}

type PostService struct {
	postRepo repo.PostRepo
}

func NewPostService(postRepo repo.PostRepo) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (p *PostService) CreatePost(ctx context.Context, post *types.Post, usrID string) (string, error) {
	return p.postRepo.CreatePost(ctx, post, usrID)
}

func (p *PostService) GetPostByID(ctx context.Context, id string) (*types.Post, error) {
	return p.postRepo.GetPostByID(ctx, id)
}

func (p *PostService) GetPosts(ctx context.Context, skip, limit int64) ([]*types.Post, error) {
	return p.postRepo.GetPosts(ctx, skip, limit)
}

func (p *PostService) GetPostCount(ctx context.Context) (int64, error) {
	return p.postRepo.GetPostCount(ctx)
}

func (p *PostService) AddComment(ctx context.Context, postID string, userID string, comment *types.Comment) (*types.Post, error)  {
	return p.postRepo.AddComment(ctx, postID, userID, comment)
}

func (p *PostService) AddLike(ctx context.Context, postID, userID string) error {
	return p.postRepo.LikePost(ctx, postID, userID)
}

func (p *PostService) RemoveLike(ctx context.Context, postID, userID string) error {
	return p.postRepo.UnlikePost(ctx, postID, userID)
}

func (p *PostService)UpdatePost(ctx context.Context, post *types.Post)(string, error) {
	return p.postRepo.UpdatePost(ctx, post)
}
type UserInterface interface {
	CreateUser(user *types.User) error
}

type UserService struct {
	userRepo repo.UserRepo
}

func NewUserService(usrRepo repo.UserRepo) *UserService {
	return &UserService{
		userRepo: usrRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *types.User) (string, error) {
	return s.userRepo.CreateUser(ctx, user)
}
