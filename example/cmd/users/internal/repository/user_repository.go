package repository

import (
	"sync"
	"time"

	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	GetByID(id string) (*users.User, bool)
	GetByIDs(ids []string) []*users.User
	GetByType(userType users.UserType) []*users.User
	GetAll() []*users.User
	GetProfile(userID string) (*users.UserProfile, bool)
}

// InMemoryUserRepository is an in-memory implementation of UserRepository
type InMemoryUserRepository struct {
	mu       sync.RWMutex
	users    map[string]*users.User
	profiles map[string]*users.UserProfile
}

// NewInMemoryUserRepository creates a new in-memory user repository with sample data
func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		users:    make(map[string]*users.User),
		profiles: make(map[string]*users.UserProfile),
	}
	repo.seedData()
	return repo
}

// GetByID retrieves a user by their ID
func (r *InMemoryUserRepository) GetByID(id string) (*users.User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	return user, ok
}

// GetByIDs retrieves users by their IDs
func (r *InMemoryUserRepository) GetByIDs(ids []string) []*users.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*users.User, 0, len(ids))
	for _, id := range ids {
		if user, ok := r.users[id]; ok {
			result = append(result, user)
		}
	}
	return result
}

// GetByType retrieves all users of a specific type
func (r *InMemoryUserRepository) GetByType(userType users.UserType) []*users.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*users.User
	for _, user := range r.users {
		if user.Type == userType {
			result = append(result, user)
		}
	}
	return result
}

// GetAll retrieves all users
func (r *InMemoryUserRepository) GetAll() []*users.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*users.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}
	return result
}

// GetProfile retrieves a user's profile
func (r *InMemoryUserRepository) GetProfile(userID string) (*users.UserProfile, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	profile, ok := r.profiles[userID]
	return profile, ok
}

// seedData initializes the repository with sample data
func (r *InMemoryUserRepository) seedData() {
	now := timestamppb.New(time.Now())
	lastWeek := timestamppb.New(time.Now().AddDate(0, 0, -7))
	lastMonth := timestamppb.New(time.Now().AddDate(0, -1, 0))
	sixMonthsAgo := timestamppb.New(time.Now().AddDate(0, -6, 0))

	r.users = map[string]*users.User{
		"1": {
			Id:            "1",
			Email:         "alice@example.com",
			FirstName:     "Alice",
			LastName:      "Johnson",
			Type:          users.UserType_CUSTOMER,
			CreatedAt:     sixMonthsAgo,
			LastLogin:     now,
			EmailVerified: wrapperspb.Bool(true),
			Phone:         stringPtr("+1-555-0101"),
		},
		"2": {
			Id:            "2",
			Email:         "bob@example.com",
			FirstName:     "Bob",
			LastName:      "Smith",
			Type:          users.UserType_SELLER,
			CreatedAt:     lastMonth,
			LastLogin:     lastWeek,
			EmailVerified: wrapperspb.Bool(true),
			Phone:         stringPtr("+1-555-0102"),
		},
		"3": {
			Id:            "3",
			Email:         "charlie@example.com",
			FirstName:     "Charlie",
			LastName:      "Brown",
			Type:          users.UserType_CUSTOMER,
			CreatedAt:     lastMonth,
			LastLogin:     now,
			EmailVerified: wrapperspb.Bool(false),
		},
		"seller1": {
			Id:            "seller1",
			Email:         "techgadgets@example.com",
			FirstName:     "Tech",
			LastName:      "Gadgets Inc",
			Type:          users.UserType_SELLER,
			CreatedAt:     sixMonthsAgo,
			LastLogin:     now,
			EmailVerified: wrapperspb.Bool(true),
			Phone:         stringPtr("+1-555-TECH"),
		},
		"seller2": {
			Id:            "seller2",
			Email:         "sportsworld@example.com",
			FirstName:     "Sports",
			LastName:      "World LLC",
			Type:          users.UserType_SELLER,
			CreatedAt:     sixMonthsAgo,
			LastLogin:     lastWeek,
			EmailVerified: wrapperspb.Bool(true),
			Phone:         stringPtr("+1-555-SPORT"),
		},
		"seller3": {
			Id:            "seller3",
			Email:         "homeessentials@example.com",
			FirstName:     "Home",
			LastName:      "Essentials Co",
			Type:          users.UserType_SELLER,
			CreatedAt:     lastMonth,
			LastLogin:     now,
			EmailVerified: wrapperspb.Bool(true),
		},
	}

	r.profiles = map[string]*users.UserProfile{
		"1": {
			UserId: "1",
			Addresses: []*users.Address{
				{
					Id:            "addr1",
					Line1:         "123 Main St",
					Line2:         "Apt 4B",
					City:          "San Francisco",
					StateProvince: "CA",
					PostalCode:    "94102",
					Country:       "USA",
					Type:          users.AddressType_BOTH,
					IsDefault:     wrapperspb.Bool(true),
				},
				{
					Id:            "addr2",
					Line1:         "456 Work Ave",
					City:          "San Francisco",
					StateProvince: "CA",
					PostalCode:    "94103",
					Country:       "USA",
					Type:          users.AddressType_SHIPPING,
					IsDefault:     wrapperspb.Bool(false),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:    true,
				PreferredLanguage:  "en",
				PreferredCurrency:  "USD",
				FavoriteCategories: []string{"ELECTRONICS", "BOOKS"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Gold",
				Points:             2500,
				DiscountPercentage: 15.0,
			},
			TotalOrders: 47,
			MemberSince: sixMonthsAgo,
		},
		"2": {
			UserId: "2",
			Addresses: []*users.Address{
				{
					Id:            "addr3",
					Line1:         "789 Business Blvd",
					City:          "Seattle",
					StateProvince: "WA",
					PostalCode:    "98101",
					Country:       "USA",
					Type:          users.AddressType_BOTH,
					IsDefault:     wrapperspb.Bool(true),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:    false,
				PreferredLanguage:  "en",
				PreferredCurrency:  "USD",
				FavoriteCategories: []string{"SPORTS", "HOME_GARDEN"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Platinum",
				Points:             5000,
				DiscountPercentage: 20.0,
			},
			TotalOrders: 124,
			MemberSince: lastMonth,
		},
		"3": {
			UserId: "3",
			Addresses: []*users.Address{
				{
					Id:            "addr4",
					Line1:         "321 College Dr",
					City:          "Austin",
					StateProvince: "TX",
					PostalCode:    "78701",
					Country:       "USA",
					Type:          users.AddressType_SHIPPING,
					IsDefault:     wrapperspb.Bool(true),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:    true,
				PreferredLanguage:  "en",
				PreferredCurrency:  "USD",
				FavoriteCategories: []string{"CLOTHING", "TOYS"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Silver",
				Points:             750,
				DiscountPercentage: 10.0,
			},
			TotalOrders: 12,
			MemberSince: lastMonth,
		},
	}
}

// Helper function for creating string pointers
func stringPtr(s string) *string {
	return &s
}
