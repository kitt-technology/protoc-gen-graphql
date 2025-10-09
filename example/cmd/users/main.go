package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kitt-technology/protoc-gen-graphql/example/users"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	users.RegisterUsersServer(s, &UserService{})
	reflection.Register(s)

	fmt.Println("========================================")
	fmt.Println("Users gRPC Service")
	fmt.Println("========================================")
	fmt.Println("Listening on: localhost:50052")
	fmt.Println("\nThis service provides user/customer data")
	fmt.Println("Available sample users:")
	fmt.Println("  - alice@example.com (Gold tier customer)")
	fmt.Println("  - bob@example.com (Platinum tier seller)")
	fmt.Println("  - charlie@example.com (Silver tier customer)")
	fmt.Println("\nReady to accept gRPC requests...")
	fmt.Println("========================================\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type UserService struct {
	users.UnimplementedUsersServer
}

func (s UserService) LoadUsers(ctx context.Context, request *graphql.BatchRequest) (*users.UsersBatchResponse, error) {
	results := make(map[string]*users.User)

	for _, key := range request.Keys {
		if user, ok := usersDb[key]; ok {
			results[key] = user
		}
	}

	return &users.UsersBatchResponse{Results: results}, nil
}

func (s UserService) GetUsers(ctx context.Context, request *users.GetUsersRequest) (*users.GetUsersResponse, error) {
	var userList []*users.User

	if len(request.Ids) > 0 {
		for _, id := range request.Ids {
			if user, ok := usersDb[id]; ok {
				userList = append(userList, user)
			}
		}
	} else {
		for _, user := range usersDb {
			userList = append(userList, user)
		}
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

func (s UserService) GetUserProfile(ctx context.Context, request *users.GetUserProfileRequest) (*users.UserProfile, error) {
	profile, ok := profilesDb[request.UserId]
	if !ok {
		return &users.UserProfile{UserId: request.UserId}, nil
	}
	return profile, nil
}

var usersDb map[string]*users.User
var profilesDb map[string]*users.UserProfile

func init() {
	now := timestamppb.New(time.Now())
	lastWeek := timestamppb.New(time.Now().AddDate(0, 0, -7))
	lastMonth := timestamppb.New(time.Now().AddDate(0, -1, 0))
	sixMonthsAgo := timestamppb.New(time.Now().AddDate(0, -6, 0))

	usersDb = map[string]*users.User{
		"1": {
			Id:            "1",
			Email:         "alice@example.com",
			FirstName:     "Alice",
			LastName:      "Johnson",
			Type:          users.UserType_CUSTOMER,
			CreatedAt:     sixMonthsAgo,
			LastLogin:     now,
			EmailVerified: wrapperspb.Bool(true),
			Phone:         proto.String("+1-555-0101"),
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
			Phone:         proto.String("+1-555-0102"),
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
			Phone:         proto.String("+1-555-TECH"),
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
			Phone:         proto.String("+1-555-SPORT"),
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

	profilesDb = map[string]*users.UserProfile{
		"1": {
			UserId: "1",
			Addresses: []*users.Address{
				{
					Id:         "addr1",
					Line1:      "123 Main St",
					Line2:      "Apt 4B",
					City:       "San Francisco",
					StateProvince: "CA",
					PostalCode: "94102",
					Country:    "USA",
					Type:       users.AddressType_BOTH,
					IsDefault:  wrapperspb.Bool(true),
				},
				{
					Id:         "addr2",
					Line1:      "456 Work Ave",
					City:       "San Francisco",
					StateProvince: "CA",
					PostalCode: "94103",
					Country:    "USA",
					Type:       users.AddressType_SHIPPING,
					IsDefault:  wrapperspb.Bool(false),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:     true,
				PreferredLanguage:   "en",
				PreferredCurrency:   "USD",
				FavoriteCategories:  []string{"ELECTRONICS", "BOOKS"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Gold",
				Points:             2500,
				DiscountPercentage: 15.0,
			},
			TotalOrders:  47,
			MemberSince:  sixMonthsAgo,
		},
		"2": {
			UserId: "2",
			Addresses: []*users.Address{
				{
					Id:         "addr3",
					Line1:      "789 Business Blvd",
					City:       "Seattle",
					StateProvince: "WA",
					PostalCode: "98101",
					Country:    "USA",
					Type:       users.AddressType_BOTH,
					IsDefault:  wrapperspb.Bool(true),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:     false,
				PreferredLanguage:   "en",
				PreferredCurrency:   "USD",
				FavoriteCategories:  []string{"SPORTS", "HOME_GARDEN"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Platinum",
				Points:             5000,
				DiscountPercentage: 20.0,
			},
			TotalOrders:  124,
			MemberSince:  lastMonth,
		},
		"3": {
			UserId: "3",
			Addresses: []*users.Address{
				{
					Id:         "addr4",
					Line1:      "321 College Dr",
					City:       "Austin",
					StateProvince: "TX",
					PostalCode: "78701",
					Country:    "USA",
					Type:       users.AddressType_SHIPPING,
					IsDefault:  wrapperspb.Bool(true),
				},
			},
			Preferences: &users.UserPreferences{
				MarketingEmails:     true,
				PreferredLanguage:   "en",
				PreferredCurrency:   "USD",
				FavoriteCategories:  []string{"CLOTHING", "TOYS"},
			},
			Loyalty: &users.LoyaltyInfo{
				Tier:               "Silver",
				Points:             750,
				DiscountPercentage: 10.0,
			},
			TotalOrders:  12,
			MemberSince:  lastMonth,
		},
	}
}

// proto is a helper package for creating proto optional values
var proto = struct {
	String func(string) *string
}{
	String: func(s string) *string { return &s },
}