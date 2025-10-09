package repository

import (
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/example/users"
)

func TestInMemoryUserRepository_GetByID(t *testing.T) {
	repo := NewInMemoryUserRepository()

	tests := []struct {
		name      string
		id        string
		wantFound bool
		wantEmail string
	}{
		{
			name:      "existing user",
			id:        "1",
			wantFound: true,
			wantEmail: "alice@example.com",
		},
		{
			name:      "existing seller",
			id:        "seller1",
			wantFound: true,
			wantEmail: "techgadgets@example.com",
		},
		{
			name:      "non-existing user",
			id:        "999",
			wantFound: false,
		},
		{
			name:      "empty id",
			id:        "",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, found := repo.GetByID(tt.id)
			if found != tt.wantFound {
				t.Errorf("GetByID() found = %v, want %v", found, tt.wantFound)
			}
			if found && user.Email != tt.wantEmail {
				t.Errorf("GetByID() email = %v, want %v", user.Email, tt.wantEmail)
			}
		})
	}
}

func TestInMemoryUserRepository_GetByIDs(t *testing.T) {
	repo := NewInMemoryUserRepository()

	tests := []struct {
		name      string
		ids       []string
		wantCount int
	}{
		{
			name:      "multiple existing users",
			ids:       []string{"1", "2", "3"},
			wantCount: 3,
		},
		{
			name:      "some existing, some not",
			ids:       []string{"1", "999", "2"},
			wantCount: 2,
		},
		{
			name:      "no existing users",
			ids:       []string{"999", "998"},
			wantCount: 0,
		},
		{
			name:      "empty ids",
			ids:       []string{},
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetByIDs(tt.ids)
			if len(result) != tt.wantCount {
				t.Errorf("GetByIDs() count = %d, want %d", len(result), tt.wantCount)
			}
		})
	}
}

func TestInMemoryUserRepository_GetByType(t *testing.T) {
	repo := NewInMemoryUserRepository()

	tests := []struct {
		name      string
		userType  users.UserType
		wantCount int
	}{
		{
			name:      "get customers",
			userType:  users.UserType_CUSTOMER,
			wantCount: 2, // alice and charlie
		},
		{
			name:      "get sellers",
			userType:  users.UserType_SELLER,
			wantCount: 4, // bob + 3 seller accounts
		},
		{
			name:      "get admins",
			userType:  users.UserType_ADMIN,
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetByType(tt.userType)
			if len(result) != tt.wantCount {
				t.Errorf("GetByType() count = %d, want %d", len(result), tt.wantCount)
			}
			// Verify all returned users are of the requested type
			for _, user := range result {
				if user.Type != tt.userType {
					t.Errorf("GetByType() returned user with type %v, want %v", user.Type, tt.userType)
				}
			}
		})
	}
}

func TestInMemoryUserRepository_GetAll(t *testing.T) {
	repo := NewInMemoryUserRepository()

	result := repo.GetAll()

	// We seeded 6 users
	expectedCount := 6
	if len(result) != expectedCount {
		t.Errorf("GetAll() count = %d, want %d", len(result), expectedCount)
	}
}

func TestInMemoryUserRepository_GetProfile(t *testing.T) {
	repo := NewInMemoryUserRepository()

	tests := []struct {
		name      string
		userID    string
		wantFound bool
		wantTier  string
	}{
		{
			name:      "existing profile",
			userID:    "1",
			wantFound: true,
			wantTier:  "Gold",
		},
		{
			name:      "another existing profile",
			userID:    "2",
			wantFound: true,
			wantTier:  "Platinum",
		},
		{
			name:      "non-existing profile",
			userID:    "999",
			wantFound: false,
		},
		{
			name:      "user without profile",
			userID:    "seller1",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, found := repo.GetProfile(tt.userID)
			if found != tt.wantFound {
				t.Errorf("GetProfile() found = %v, want %v", found, tt.wantFound)
			}
			if found && profile.Loyalty != nil && profile.Loyalty.Tier != tt.wantTier {
				t.Errorf("GetProfile() tier = %v, want %v", profile.Loyalty.Tier, tt.wantTier)
			}
		})
	}
}

func TestInMemoryUserRepository_Concurrency(t *testing.T) {
	repo := NewInMemoryUserRepository()

	// Test concurrent reads
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			_ = repo.GetAll()
			_, _ = repo.GetByID("1")
			_ = repo.GetByType(users.UserType_CUSTOMER)
			_, _ = repo.GetProfile("1")
			done <- true
		}()
	}

	// Wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestInMemoryUserRepository_ProfileData(t *testing.T) {
	repo := NewInMemoryUserRepository()

	// Test that profile data is correctly structured
	profile, found := repo.GetProfile("1")
	if !found {
		t.Fatal("GetProfile() could not find profile for user 1")
	}

	// Test addresses
	if len(profile.Addresses) != 2 {
		t.Errorf("Profile has %d addresses, want 2", len(profile.Addresses))
	}

	// Test address fields
	if profile.Addresses[0].City != "San Francisco" {
		t.Errorf("Address city = %v, want San Francisco", profile.Addresses[0].City)
	}

	// Test preferences
	if profile.Preferences == nil {
		t.Error("Profile preferences is nil")
	}
	if profile.Preferences != nil && profile.Preferences.PreferredLanguage != "en" {
		t.Errorf("Preferred language = %v, want en", profile.Preferences.PreferredLanguage)
	}

	// Test loyalty
	if profile.Loyalty == nil {
		t.Error("Profile loyalty is nil")
	}
	if profile.Loyalty != nil && profile.Loyalty.Points != 2500 {
		t.Errorf("Loyalty points = %d, want 2500", profile.Loyalty.Points)
	}
}
