package service

import (
	"context"
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// mockUserRepository is a mock implementation of UserRepository for testing
type mockUserRepository struct {
	users    map[string]*users.User
	profiles map[string]*users.UserProfile
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: map[string]*users.User{
			"1": {
				Id:            "1",
				Email:         "user1@example.com",
				FirstName:     "User",
				LastName:      "One",
				Type:          users.UserType_CUSTOMER,
				EmailVerified: wrapperspb.Bool(true),
			},
			"2": {
				Id:            "2",
				Email:         "user2@example.com",
				FirstName:     "User",
				LastName:      "Two",
				Type:          users.UserType_SELLER,
				EmailVerified: wrapperspb.Bool(false),
			},
			"3": {
				Id:        "3",
				Email:     "user3@example.com",
				FirstName: "User",
				LastName:  "Three",
				Type:      users.UserType_CUSTOMER,
			},
		},
		profiles: map[string]*users.UserProfile{
			"1": {
				UserId: "1",
				Loyalty: &users.LoyaltyInfo{
					Tier:   "Gold",
					Points: 1000,
				},
				TotalOrders: 10,
			},
			"2": {
				UserId: "2",
				Loyalty: &users.LoyaltyInfo{
					Tier:   "Silver",
					Points: 500,
				},
				TotalOrders: 5,
			},
		},
	}
}

func (m *mockUserRepository) GetByID(id string) (*users.User, bool) {
	user, ok := m.users[id]
	return user, ok
}

func (m *mockUserRepository) GetByIDs(ids []string) []*users.User {
	result := make([]*users.User, 0, len(ids))
	for _, id := range ids {
		if user, ok := m.users[id]; ok {
			result = append(result, user)
		}
	}
	return result
}

func (m *mockUserRepository) GetByType(userType users.UserType) []*users.User {
	var result []*users.User
	for _, user := range m.users {
		if user.Type == userType {
			result = append(result, user)
		}
	}
	return result
}

func (m *mockUserRepository) GetAll() []*users.User {
	result := make([]*users.User, 0, len(m.users))
	for _, user := range m.users {
		result = append(result, user)
	}
	return result
}

func (m *mockUserRepository) GetProfile(userID string) (*users.UserProfile, bool) {
	profile, ok := m.profiles[userID]
	return profile, ok
}

func TestUserService_GetUsers(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)
	ctx := context.Background()

	tests := []struct {
		name      string
		request   *users.GetUsersRequest
		wantCount int
		wantErr   bool
	}{
		{
			name:      "get all users",
			request:   &users.GetUsersRequest{},
			wantCount: 3,
			wantErr:   false,
		},
		{
			name: "get by IDs",
			request: &users.GetUsersRequest{
				Ids: []string{"1", "2"},
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "get by single ID",
			request: &users.GetUsersRequest{
				Ids: []string{"1"},
			},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "get by non-existent ID",
			request: &users.GetUsersRequest{
				Ids: []string{"999"},
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name: "get by multiple IDs with some non-existent",
			request: &users.GetUsersRequest{
				Ids: []string{"1", "999", "2"},
			},
			wantCount: 2,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetUsers(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Users) != tt.wantCount {
				t.Errorf("GetUsers() count = %d, want %d", len(resp.Users), tt.wantCount)
			}
			if !tt.wantErr && resp.PageInfo == nil {
				t.Error("GetUsers() PageInfo is nil")
			}
			if !tt.wantErr && resp.PageInfo != nil && resp.PageInfo.TotalCount != int32(tt.wantCount) {
				t.Errorf("GetUsers() PageInfo.TotalCount = %d, want %d", resp.PageInfo.TotalCount, tt.wantCount)
			}
		})
	}
}

func TestUserService_LoadUsers(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)
	ctx := context.Background()

	tests := []struct {
		name          string
		request       *graphql.BatchRequest
		wantResultLen int
		wantErr       bool
	}{
		{
			name: "load single user",
			request: &graphql.BatchRequest{
				Keys: []string{"1"},
			},
			wantResultLen: 1,
			wantErr:       false,
		},
		{
			name: "load multiple users",
			request: &graphql.BatchRequest{
				Keys: []string{"1", "2", "3"},
			},
			wantResultLen: 3,
			wantErr:       false,
		},
		{
			name: "load with some non-existent users",
			request: &graphql.BatchRequest{
				Keys: []string{"1", "999", "2"},
			},
			wantResultLen: 2,
			wantErr:       false,
		},
		{
			name: "load non-existent users only",
			request: &graphql.BatchRequest{
				Keys: []string{"999", "998"},
			},
			wantResultLen: 0,
			wantErr:       false,
		},
		{
			name: "empty keys",
			request: &graphql.BatchRequest{
				Keys: []string{},
			},
			wantResultLen: 0,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.LoadUsers(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Results) != tt.wantResultLen {
				t.Errorf("LoadUsers() result count = %d, want %d", len(resp.Results), tt.wantResultLen)
			}
		})
	}
}

func TestUserService_GetUserProfile(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)
	ctx := context.Background()

	tests := []struct {
		name       string
		request    *users.GetUserProfileRequest
		wantTier   string
		wantOrders int32
		wantErr    bool
	}{
		{
			name: "get existing profile",
			request: &users.GetUserProfileRequest{
				UserId: "1",
			},
			wantTier:   "Gold",
			wantOrders: 10,
			wantErr:    false,
		},
		{
			name: "get another existing profile",
			request: &users.GetUserProfileRequest{
				UserId: "2",
			},
			wantTier:   "Silver",
			wantOrders: 5,
			wantErr:    false,
		},
		{
			name: "get non-existent profile",
			request: &users.GetUserProfileRequest{
				UserId: "999",
			},
			wantTier:   "",
			wantOrders: 0,
			wantErr:    false,
		},
		{
			name: "get profile for user without profile",
			request: &users.GetUserProfileRequest{
				UserId: "3",
			},
			wantTier:   "",
			wantOrders: 0,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetUserProfile(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if resp.UserId != tt.request.UserId {
					t.Errorf("GetUserProfile() UserId = %v, want %v", resp.UserId, tt.request.UserId)
				}
				if tt.wantTier != "" && (resp.Loyalty == nil || resp.Loyalty.Tier != tt.wantTier) {
					tier := ""
					if resp.Loyalty != nil {
						tier = resp.Loyalty.Tier
					}
					t.Errorf("GetUserProfile() Loyalty.Tier = %v, want %v", tier, tt.wantTier)
				}
				if resp.TotalOrders != tt.wantOrders {
					t.Errorf("GetUserProfile() TotalOrders = %d, want %d", resp.TotalOrders, tt.wantOrders)
				}
			}
		})
	}
}

func TestUserService_LoadUsers_ReturnCorrectUsers(t *testing.T) {
	repo := newMockUserRepository()
	service := NewUserService(repo)
	ctx := context.Background()

	request := &graphql.BatchRequest{
		Keys: []string{"1", "2"},
	}

	resp, err := service.LoadUsers(ctx, request)
	if err != nil {
		t.Fatalf("LoadUsers() error = %v", err)
	}

	// Verify correct users are returned
	if user, ok := resp.Results["1"]; !ok {
		t.Error("LoadUsers() missing user 1")
	} else if user.Email != "user1@example.com" {
		t.Errorf("LoadUsers() user 1 email = %v, want user1@example.com", user.Email)
	}

	if user, ok := resp.Results["2"]; !ok {
		t.Error("LoadUsers() missing user 2")
	} else if user.Email != "user2@example.com" {
		t.Errorf("LoadUsers() user 2 email = %v, want user2@example.com", user.Email)
	}
}
