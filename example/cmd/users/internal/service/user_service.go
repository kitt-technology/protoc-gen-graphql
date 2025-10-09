package service

import (
	"context"

	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/users/internal/repository"
	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
)

// UserService implements the Users gRPC service
type UserService struct {
	users.UnimplementedUsersServer
	repo repository.UserRepository
}

// NewUserService creates a new UserService with the given repository
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// LoadUsers implements batch loading for users by ID
func (s *UserService) LoadUsers(ctx context.Context, request *graphql.BatchRequest) (*users.UsersBatchResponse, error) {
	results := make(map[string]*users.User)

	for _, key := range request.Keys {
		if user, ok := s.repo.GetByID(key); ok {
			results[key] = user
		}
	}

	return &users.UsersBatchResponse{Results: results}, nil
}

// GetUsers retrieves users based on the request criteria
func (s *UserService) GetUsers(ctx context.Context, request *users.GetUsersRequest) (*users.GetUsersResponse, error) {
	var userList []*users.User

	// Priority: Ids > All
	if len(request.Ids) > 0 {
		userList = s.repo.GetByIDs(request.Ids)
	} else {
		userList = s.repo.GetAll()
	}

	return &users.GetUsersResponse{
		Users: userList,
		PageInfo: &graphql.PageInfo{
			TotalCount:  int32(len(userList)),
			EndCursor:   "cursor_end",
			HasNextPage: false,
		},
	}, nil
}

// GetUserProfile retrieves a user's profile information
func (s *UserService) GetUserProfile(ctx context.Context, request *users.GetUserProfileRequest) (*users.UserProfile, error) {
	profile, ok := s.repo.GetProfile(request.UserId)
	if !ok {
		// Return empty profile if not found
		return &users.UserProfile{UserId: request.UserId}, nil
	}
	return profile, nil
}
