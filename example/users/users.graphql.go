package users

import (
	gql "github.com/graphql-go/graphql"
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"strings"
)

var UserTypeGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "UserType",
	Values: gql.EnumValueConfigMap{
		"ADMIN": &gql.EnumValueConfig{
			Value: UserType(2),
		},
		"CUSTOMER": &gql.EnumValueConfig{
			Value: UserType(0),
		},
		"SELLER": &gql.EnumValueConfig{
			Value: UserType(1),
		},
	},
})

var UserTypeGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "UserType",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(UserType).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var AddressTypeGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "AddressType",
	Values: gql.EnumValueConfigMap{
		"BILLING": &gql.EnumValueConfig{
			Value: AddressType(1),
		},
		"BOTH": &gql.EnumValueConfig{
			Value: AddressType(2),
		},
		"SHIPPING": &gql.EnumValueConfig{
			Value: AddressType(0),
		},
	},
})

var AddressTypeGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "AddressType",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(AddressType).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var GetUsersRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetUsersRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"type": &gql.Field{
			Type: UserTypeGraphqlEnum,
		},
	},
})

var GetUsersRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetUsersRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"type": &gql.InputObjectFieldConfig{
			Type: UserTypeGraphqlEnum,
		},
	},
})

var GetUsersRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(gql.String)),
	},
	"type": &gql.ArgumentConfig{
		Type: UserTypeGraphqlEnum,
	},
}

func GetUsersRequestFromArgs(args map[string]interface{}) *GetUsersRequest {
	return GetUsersRequestInstanceFromArgs(&GetUsersRequest{}, args)
}

func GetUsersRequestInstanceFromArgs(objectFromArgs *GetUsersRequest, args map[string]interface{}) *GetUsersRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		ids := make([]string, 0)

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	if args["type"] != nil {
		val := args["type"]
		objectFromArgs.Type = val.(UserType)
	}
	return objectFromArgs
}

func (objectFromArgs *GetUsersRequest) FromArgs(args map[string]interface{}) {
	GetUsersRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetUsersRequest) XXX_GraphqlType() *gql.Object {
	return GetUsersRequestGraphqlType
}

func (msg *GetUsersRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetUsersRequestGraphqlArgs
}

func (msg *GetUsersRequest) XXX_Package() string {
	return "users"
}

var GetUsersResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetUsersResponse",
	Fields: gql.Fields{
		"users": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(UserGraphqlType)),
		},
		"pageInfo": &gql.Field{
			Type: pg.PageInfoGraphqlType,
		},
	},
})

var GetUsersResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetUsersResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"users": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(UserGraphqlInputType)),
		},
		"pageInfo": &gql.InputObjectFieldConfig{
			Type: pg.PageInfoGraphqlInputType,
		},
	},
})

var GetUsersResponseGraphqlArgs = gql.FieldConfigArgument{
	"users": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(UserGraphqlInputType)),
	},
	"pageInfo": &gql.ArgumentConfig{
		Type: pg.PageInfoGraphqlInputType,
	},
}

func GetUsersResponseFromArgs(args map[string]interface{}) *GetUsersResponse {
	return GetUsersResponseInstanceFromArgs(&GetUsersResponse{}, args)
}

func GetUsersResponseInstanceFromArgs(objectFromArgs *GetUsersResponse, args map[string]interface{}) *GetUsersResponse {
	if args["users"] != nil {
		usersInterfaceList := args["users"].([]interface{})
		users := make([]*User, 0)

		for _, val := range usersInterfaceList {
			itemResolved := UserFromArgs(val.(map[string]interface{}))
			users = append(users, itemResolved)
		}
		objectFromArgs.Users = users
	}
	if args["pageInfo"] != nil {
		val := args["pageInfo"]
		objectFromArgs.PageInfo = pg.PageInfoFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *GetUsersResponse) FromArgs(args map[string]interface{}) {
	GetUsersResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetUsersResponse) XXX_GraphqlType() *gql.Object {
	return GetUsersResponseGraphqlType
}

func (msg *GetUsersResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetUsersResponseGraphqlArgs
}

func (msg *GetUsersResponse) XXX_Package() string {
	return "users"
}

var UsersBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "UsersBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var UsersBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "UsersBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var UsersBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func UsersBatchResponseFromArgs(args map[string]interface{}) *UsersBatchResponse {
	return UsersBatchResponseInstanceFromArgs(&UsersBatchResponse{}, args)
}

func UsersBatchResponseInstanceFromArgs(objectFromArgs *UsersBatchResponse, args map[string]interface{}) *UsersBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *UsersBatchResponse) FromArgs(args map[string]interface{}) {
	UsersBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *UsersBatchResponse) XXX_GraphqlType() *gql.Object {
	return UsersBatchResponseGraphqlType
}

func (msg *UsersBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return UsersBatchResponseGraphqlArgs
}

func (msg *UsersBatchResponse) XXX_Package() string {
	return "users"
}

var GetUserProfileRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetUserProfileRequest",
	Fields: gql.Fields{
		"userId": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var GetUserProfileRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetUserProfileRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"userId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var GetUserProfileRequestGraphqlArgs = gql.FieldConfigArgument{
	"userId": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func GetUserProfileRequestFromArgs(args map[string]interface{}) *GetUserProfileRequest {
	return GetUserProfileRequestInstanceFromArgs(&GetUserProfileRequest{}, args)
}

func GetUserProfileRequestInstanceFromArgs(objectFromArgs *GetUserProfileRequest, args map[string]interface{}) *GetUserProfileRequest {
	if args["userId"] != nil {
		val := args["userId"]
		objectFromArgs.UserId = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *GetUserProfileRequest) FromArgs(args map[string]interface{}) {
	GetUserProfileRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetUserProfileRequest) XXX_GraphqlType() *gql.Object {
	return GetUserProfileRequestGraphqlType
}

func (msg *GetUserProfileRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetUserProfileRequestGraphqlArgs
}

func (msg *GetUserProfileRequest) XXX_Package() string {
	return "users"
}

var UserGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Customer",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"email": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"firstName": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"lastName": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"type": &gql.Field{
			Type: UserTypeGraphqlEnum,
		},
		"createdAt": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"lastLogin": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"emailVerified": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*User) == nil || p.Source.(*User).EmailVerified == nil {
					return nil, nil
				}
				return p.Source.(*User).EmailVerified.Value, nil
			},
		},
		"phone": &gql.Field{
			Type: gql.String,
		},
	},
})

var UserGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "CustomerInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"email": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"firstName": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"lastName": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"type": &gql.InputObjectFieldConfig{
			Type: UserTypeGraphqlEnum,
		},
		"createdAt": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"lastLogin": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"emailVerified": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"phone": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
	},
})

var UserGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"email": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"firstName": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"lastName": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"type": &gql.ArgumentConfig{
		Type: UserTypeGraphqlEnum,
	},
	"createdAt": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"lastLogin": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"emailVerified": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"phone": &gql.ArgumentConfig{
		Type: gql.String,
	},
}

func UserFromArgs(args map[string]interface{}) *User {
	return UserInstanceFromArgs(&User{}, args)
}

func UserInstanceFromArgs(objectFromArgs *User, args map[string]interface{}) *User {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["email"] != nil {
		val := args["email"]
		objectFromArgs.Email = string(val.(string))
	}
	if args["firstName"] != nil {
		val := args["firstName"]
		objectFromArgs.FirstName = string(val.(string))
	}
	if args["lastName"] != nil {
		val := args["lastName"]
		objectFromArgs.LastName = string(val.(string))
	}
	if args["type"] != nil {
		val := args["type"]
		objectFromArgs.Type = val.(UserType)
	}
	if args["createdAt"] != nil {
		val := args["createdAt"]
		objectFromArgs.CreatedAt = pg.ToTimestamp(val)
	}
	if args["lastLogin"] != nil {
		val := args["lastLogin"]
		objectFromArgs.LastLogin = pg.ToTimestamp(val)
	}
	if args["emailVerified"] != nil {
		val := args["emailVerified"]
		objectFromArgs.EmailVerified = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["phone"] != nil {
		val := args["phone"]
		ptr := string(val.(string))
		objectFromArgs.Phone = &ptr
	}
	return objectFromArgs
}

func (objectFromArgs *User) FromArgs(args map[string]interface{}) {
	UserInstanceFromArgs(objectFromArgs, args)
}

func (msg *User) XXX_GraphqlType() *gql.Object {
	return UserGraphqlType
}

func (msg *User) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return UserGraphqlArgs
}

func (msg *User) XXX_Package() string {
	return "users"
}

var UserProfileGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "UserProfile",
	Fields: gql.Fields{
		"userId": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"addresses": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(AddressGraphqlType)),
		},
		"preferences": &gql.Field{
			Type: UserPreferencesGraphqlType,
		},
		"loyalty": &gql.Field{
			Type: LoyaltyInfoGraphqlType,
		},
		"totalOrders": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"memberSince": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
	},
})

var UserProfileGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "UserProfileInput",
	Fields: gql.InputObjectConfigFieldMap{
		"userId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"addresses": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(AddressGraphqlInputType)),
		},
		"preferences": &gql.InputObjectFieldConfig{
			Type: UserPreferencesGraphqlInputType,
		},
		"loyalty": &gql.InputObjectFieldConfig{
			Type: LoyaltyInfoGraphqlInputType,
		},
		"totalOrders": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"memberSince": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
	},
})

var UserProfileGraphqlArgs = gql.FieldConfigArgument{
	"userId": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"addresses": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(AddressGraphqlInputType)),
	},
	"preferences": &gql.ArgumentConfig{
		Type: UserPreferencesGraphqlInputType,
	},
	"loyalty": &gql.ArgumentConfig{
		Type: LoyaltyInfoGraphqlInputType,
	},
	"totalOrders": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"memberSince": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
}

func UserProfileFromArgs(args map[string]interface{}) *UserProfile {
	return UserProfileInstanceFromArgs(&UserProfile{}, args)
}

func UserProfileInstanceFromArgs(objectFromArgs *UserProfile, args map[string]interface{}) *UserProfile {
	if args["userId"] != nil {
		val := args["userId"]
		objectFromArgs.UserId = string(val.(string))
	}
	if args["addresses"] != nil {
		addressesInterfaceList := args["addresses"].([]interface{})
		addresses := make([]*Address, 0)

		for _, val := range addressesInterfaceList {
			itemResolved := AddressFromArgs(val.(map[string]interface{}))
			addresses = append(addresses, itemResolved)
		}
		objectFromArgs.Addresses = addresses
	}
	if args["preferences"] != nil {
		val := args["preferences"]
		objectFromArgs.Preferences = UserPreferencesFromArgs(val.(map[string]interface{}))
	}
	if args["loyalty"] != nil {
		val := args["loyalty"]
		objectFromArgs.Loyalty = LoyaltyInfoFromArgs(val.(map[string]interface{}))
	}
	if args["totalOrders"] != nil {
		val := args["totalOrders"]
		objectFromArgs.TotalOrders = int32(val.(int))
	}
	if args["memberSince"] != nil {
		val := args["memberSince"]
		objectFromArgs.MemberSince = pg.ToTimestamp(val)
	}
	return objectFromArgs
}

func (objectFromArgs *UserProfile) FromArgs(args map[string]interface{}) {
	UserProfileInstanceFromArgs(objectFromArgs, args)
}

func (msg *UserProfile) XXX_GraphqlType() *gql.Object {
	return UserProfileGraphqlType
}

func (msg *UserProfile) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return UserProfileGraphqlArgs
}

func (msg *UserProfile) XXX_Package() string {
	return "users"
}

var AddressGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Address",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"line1": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"line2": &gql.Field{
			Type: gql.String,
		},
		"city": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"stateProvince": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"postalCode": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"country": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"type": &gql.Field{
			Type: AddressTypeGraphqlEnum,
		},
		"isDefault": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*Address) == nil || p.Source.(*Address).IsDefault == nil {
					return nil, nil
				}
				return p.Source.(*Address).IsDefault.Value, nil
			},
		},
	},
})

var AddressGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "AddressInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"line1": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"line2": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
		"city": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"stateProvince": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"postalCode": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"country": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"type": &gql.InputObjectFieldConfig{
			Type: AddressTypeGraphqlEnum,
		},
		"isDefault": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var AddressGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"line1": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"line2": &gql.ArgumentConfig{
		Type: gql.String,
	},
	"city": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"stateProvince": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"postalCode": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"country": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"type": &gql.ArgumentConfig{
		Type: AddressTypeGraphqlEnum,
	},
	"isDefault": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func AddressFromArgs(args map[string]interface{}) *Address {
	return AddressInstanceFromArgs(&Address{}, args)
}

func AddressInstanceFromArgs(objectFromArgs *Address, args map[string]interface{}) *Address {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["line1"] != nil {
		val := args["line1"]
		objectFromArgs.Line1 = string(val.(string))
	}
	if args["line2"] != nil {
		val := args["line2"]
		objectFromArgs.Line2 = string(val.(string))
	}
	if args["city"] != nil {
		val := args["city"]
		objectFromArgs.City = string(val.(string))
	}
	if args["stateProvince"] != nil {
		val := args["stateProvince"]
		objectFromArgs.StateProvince = string(val.(string))
	}
	if args["postalCode"] != nil {
		val := args["postalCode"]
		objectFromArgs.PostalCode = string(val.(string))
	}
	if args["country"] != nil {
		val := args["country"]
		objectFromArgs.Country = string(val.(string))
	}
	if args["type"] != nil {
		val := args["type"]
		objectFromArgs.Type = val.(AddressType)
	}
	if args["isDefault"] != nil {
		val := args["isDefault"]
		objectFromArgs.IsDefault = wrapperspb.Bool(bool(val.(bool)))
	}
	return objectFromArgs
}

func (objectFromArgs *Address) FromArgs(args map[string]interface{}) {
	AddressInstanceFromArgs(objectFromArgs, args)
}

func (msg *Address) XXX_GraphqlType() *gql.Object {
	return AddressGraphqlType
}

func (msg *Address) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return AddressGraphqlArgs
}

func (msg *Address) XXX_Package() string {
	return "users"
}

var UserPreferencesGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "UserPreferences",
	Fields: gql.Fields{
		"marketingEmails": &gql.Field{
			Type: gql.NewNonNull(gql.Boolean),
		},
		"preferredLanguage": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"preferredCurrency": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"favoriteCategories": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var UserPreferencesGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "UserPreferencesInput",
	Fields: gql.InputObjectConfigFieldMap{
		"marketingEmails": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Boolean),
		},
		"preferredLanguage": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"preferredCurrency": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"favoriteCategories": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
	},
})

var UserPreferencesGraphqlArgs = gql.FieldConfigArgument{
	"marketingEmails": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Boolean),
	},
	"preferredLanguage": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"preferredCurrency": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"favoriteCategories": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
}

func UserPreferencesFromArgs(args map[string]interface{}) *UserPreferences {
	return UserPreferencesInstanceFromArgs(&UserPreferences{}, args)
}

func UserPreferencesInstanceFromArgs(objectFromArgs *UserPreferences, args map[string]interface{}) *UserPreferences {
	if args["marketingEmails"] != nil {
		val := args["marketingEmails"]
		objectFromArgs.MarketingEmails = bool(val.(bool))
	}
	if args["preferredLanguage"] != nil {
		val := args["preferredLanguage"]
		objectFromArgs.PreferredLanguage = string(val.(string))
	}
	if args["preferredCurrency"] != nil {
		val := args["preferredCurrency"]
		objectFromArgs.PreferredCurrency = string(val.(string))
	}
	if args["favoriteCategories"] != nil {
		favoriteCategoriesInterfaceList := args["favoriteCategories"].([]interface{})
		favoriteCategories := make([]string, 0)

		for _, val := range favoriteCategoriesInterfaceList {
			itemResolved := string(val.(string))
			favoriteCategories = append(favoriteCategories, itemResolved)
		}
		objectFromArgs.FavoriteCategories = favoriteCategories
	}
	return objectFromArgs
}

func (objectFromArgs *UserPreferences) FromArgs(args map[string]interface{}) {
	UserPreferencesInstanceFromArgs(objectFromArgs, args)
}

func (msg *UserPreferences) XXX_GraphqlType() *gql.Object {
	return UserPreferencesGraphqlType
}

func (msg *UserPreferences) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return UserPreferencesGraphqlArgs
}

func (msg *UserPreferences) XXX_Package() string {
	return "users"
}

var LoyaltyInfoGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "LoyaltyInfo",
	Fields: gql.Fields{
		"tier": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"points": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"discountPercentage": &gql.Field{
			Type: gql.NewNonNull(gql.Float),
		},
	},
})

var LoyaltyInfoGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "LoyaltyInfoInput",
	Fields: gql.InputObjectConfigFieldMap{
		"tier": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"points": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"discountPercentage": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Float),
		},
	},
})

var LoyaltyInfoGraphqlArgs = gql.FieldConfigArgument{
	"tier": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"points": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"discountPercentage": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Float),
	},
}

func LoyaltyInfoFromArgs(args map[string]interface{}) *LoyaltyInfo {
	return LoyaltyInfoInstanceFromArgs(&LoyaltyInfo{}, args)
}

func LoyaltyInfoInstanceFromArgs(objectFromArgs *LoyaltyInfo, args map[string]interface{}) *LoyaltyInfo {
	if args["tier"] != nil {
		val := args["tier"]
		objectFromArgs.Tier = string(val.(string))
	}
	if args["points"] != nil {
		val := args["points"]
		objectFromArgs.Points = int32(val.(int))
	}
	if args["discountPercentage"] != nil {
		val := args["discountPercentage"]
		objectFromArgs.DiscountPercentage = float32(val.(float64))
	}
	return objectFromArgs
}

func (objectFromArgs *LoyaltyInfo) FromArgs(args map[string]interface{}) {
	LoyaltyInfoInstanceFromArgs(objectFromArgs, args)
}

func (msg *LoyaltyInfo) XXX_GraphqlType() *gql.Object {
	return LoyaltyInfoGraphqlType
}

func (msg *LoyaltyInfo) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return LoyaltyInfoGraphqlArgs
}

func (msg *LoyaltyInfo) XXX_Package() string {
	return "users"
}
func LoadUsersBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadUsersLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadUsersLoader").(*dataloader.Loader)
	default:
		panic("Please call users.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*User), nil
	}, nil
}

func LoadUsersBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadUsersLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadUsersLoader").(*dataloader.Loader)
	default:
		panic("Please call users.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*User
		for _, res := range resSlice {
			results = append(results, res.(*User))
		}

		return results, nil
	}, nil
}

// allMessages contains all message types from this proto package
var allMessages = []pg.GraphqlMessage{
	&GetUsersRequest{},
	&GetUsersResponse{},
	&UsersBatchResponse{},
	&GetUserProfileRequest{},
	&User{},
	&UserProfile{},
	&Address{},
	&UserPreferences{},
	&LoyaltyInfo{},
}

// UsersModule implements the Module interface for the users package
type UsersModule struct {
	usersClient  UsersClient
	usersService UsersServer

	dialOpts pg.DialOptions
}

// UsersModuleOption configures the UsersModule
type UsersModuleOption func(*UsersModule)

// WithModuleUsersClient sets the gRPC client for the Users service
func WithModuleUsersClient(client UsersClient) UsersModuleOption {
	return func(m *UsersModule) {
		m.usersClient = client
	}
}

// WithModuleUsersService sets the direct service implementation for the Users service
func WithModuleUsersService(service UsersServer) UsersModuleOption {
	return func(m *UsersModule) {
		m.usersService = service
	}
}

// WithDialOptions sets dial options for lazy client creation
func WithDialOptions(opts pg.DialOptions) UsersModuleOption {
	return func(m *UsersModule) {
		m.dialOpts = opts
	}
}

// NewUsersModule creates a new module with optional service configurations
func NewUsersModule(opts ...UsersModuleOption) *UsersModule {
	m := &UsersModule{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// getUsersClient returns the client, creating it lazily if needed
func (m *UsersModule) getUsersClient() UsersClient {
	if m.usersClient == nil {
		host := os.Getenv("USERS_SERVICE_HOST")
		if host == "" {
			host = "localhost:50052"
		}
		m.usersClient = NewUsersClient(pg.GrpcConnection(host, m.dialOpts["Users"]...))
	}
	return m.usersClient
}

// Fields returns all GraphQL query/mutation fields from all services in this module
func (m *UsersModule) Fields() gql.Fields {
	fields := gql.Fields{}

	// Users service: GetUsers method
	fields["users_GetUsers"] = &gql.Field{
		Name: "users_GetUsers",
		Type: GetUsersResponseGraphqlType,
		Args: GetUsersRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			req := GetUsersRequestFromArgs(p.Args)
			if m.usersService != nil {
				return m.usersService.GetUsers(p.Context, req)
			}
			return m.getUsersClient().GetUsers(p.Context, req)
		},
	}

	// Users service: GetUserProfile method
	fields["users_GetUserProfile"] = &gql.Field{
		Name: "users_GetUserProfile",
		Type: UserProfileGraphqlType,
		Args: GetUserProfileRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			req := GetUserProfileRequestFromArgs(p.Args)
			if m.usersService != nil {
				return m.usersService.GetUserProfile(p.Context, req)
			}
			return m.getUsersClient().GetUserProfile(p.Context, req)
		},
	}

	return fields
}

// WithLoaders registers all dataloaders from all services into the context
func (m *UsersModule) WithLoaders(ctx context.Context) context.Context {
	// Users service: LoadUsers loader
	ctx = context.WithValue(ctx, "LoadUsersLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *UsersBatchResponse
			var err error

			req := &pg.BatchRequest{
				Keys: keys.Keys(),
			}
			if m.usersService != nil {
				resp, err = m.usersService.LoadUsers(ctx, req)
			} else {
				resp, err = m.getUsersClient().LoadUsers(ctx, req)
			}

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *User
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	return ctx
}

// Messages returns all message types from this package
func (m *UsersModule) Messages() []pg.GraphqlMessage {
	return allMessages
}

// PackageName returns the proto package name
func (m *UsersModule) PackageName() string {
	return "users"
}

// Type-safe field customization methods

// AddFieldToGetUsersRequest adds a custom field to the GetUsersRequest GraphQL type
func (m *UsersModule) AddFieldToGetUsersRequest(fieldName string, field *gql.Field) {
	GetUsersRequestGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToGetUsersResponse adds a custom field to the GetUsersResponse GraphQL type
func (m *UsersModule) AddFieldToGetUsersResponse(fieldName string, field *gql.Field) {
	GetUsersResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToUsersBatchResponse adds a custom field to the UsersBatchResponse GraphQL type
func (m *UsersModule) AddFieldToUsersBatchResponse(fieldName string, field *gql.Field) {
	UsersBatchResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToGetUserProfileRequest adds a custom field to the GetUserProfileRequest GraphQL type
func (m *UsersModule) AddFieldToGetUserProfileRequest(fieldName string, field *gql.Field) {
	GetUserProfileRequestGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToUser adds a custom field to the User GraphQL type
func (m *UsersModule) AddFieldToUser(fieldName string, field *gql.Field) {
	UserGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToUserProfile adds a custom field to the UserProfile GraphQL type
func (m *UsersModule) AddFieldToUserProfile(fieldName string, field *gql.Field) {
	UserProfileGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToAddress adds a custom field to the Address GraphQL type
func (m *UsersModule) AddFieldToAddress(fieldName string, field *gql.Field) {
	AddressGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToUserPreferences adds a custom field to the UserPreferences GraphQL type
func (m *UsersModule) AddFieldToUserPreferences(fieldName string, field *gql.Field) {
	UserPreferencesGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToLoyaltyInfo adds a custom field to the LoyaltyInfo GraphQL type
func (m *UsersModule) AddFieldToLoyaltyInfo(fieldName string, field *gql.Field) {
	LoyaltyInfoGraphqlType.AddFieldConfig(fieldName, field)
}

// GraphQL type accessors

// GetUsersRequestType returns the GraphQL type for GetUsersRequest
func (m *UsersModule) GetUsersRequestType() *gql.Object {
	return GetUsersRequestGraphqlType
}

// GetUsersResponseType returns the GraphQL type for GetUsersResponse
func (m *UsersModule) GetUsersResponseType() *gql.Object {
	return GetUsersResponseGraphqlType
}

// UsersBatchResponseType returns the GraphQL type for UsersBatchResponse
func (m *UsersModule) UsersBatchResponseType() *gql.Object {
	return UsersBatchResponseGraphqlType
}

// GetUserProfileRequestType returns the GraphQL type for GetUserProfileRequest
func (m *UsersModule) GetUserProfileRequestType() *gql.Object {
	return GetUserProfileRequestGraphqlType
}

// UserType returns the GraphQL type for User
func (m *UsersModule) UserType() *gql.Object {
	return UserGraphqlType
}

// UserProfileType returns the GraphQL type for UserProfile
func (m *UsersModule) UserProfileType() *gql.Object {
	return UserProfileGraphqlType
}

// AddressType returns the GraphQL type for Address
func (m *UsersModule) AddressType() *gql.Object {
	return AddressGraphqlType
}

// UserPreferencesType returns the GraphQL type for UserPreferences
func (m *UsersModule) UserPreferencesType() *gql.Object {
	return UserPreferencesGraphqlType
}

// LoyaltyInfoType returns the GraphQL type for LoyaltyInfo
func (m *UsersModule) LoyaltyInfoType() *gql.Object {
	return LoyaltyInfoGraphqlType
}

// DataLoader accessor methods

// UsersLoadUsers loads a single *User using the users service dataloader
func (m *UsersModule) UsersLoadUsers(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return LoadUsersBatch(p, key)
}

// UsersLoadUsersMany loads multiple *User using the users service dataloader
func (m *UsersModule) UsersLoadUsersMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return LoadUsersBatchMany(p, keys)
}

// Service instance accessors

// UsersInstance is a unified interface for calling Users methods
// It works with both gRPC clients and direct service implementations
type UsersInstance interface {
	GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error)
	GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*UserProfile, error)
	LoadUsers(ctx context.Context, req *pg.BatchRequest) (*UsersBatchResponse, error)
	LoadUsersBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error)
	LoadUsersBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error)
}

type usersServerAdapter struct {
	server UsersServer
}

func (a *usersServerAdapter) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	return a.server.GetUsers(ctx, req)
}

func (a *usersServerAdapter) GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*UserProfile, error) {
	return a.server.GetUserProfile(ctx, req)
}

func (a *usersServerAdapter) LoadUsers(ctx context.Context, req *pg.BatchRequest) (*UsersBatchResponse, error) {
	return a.server.LoadUsers(ctx, req)
}

func (a *usersServerAdapter) LoadUsersBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return LoadUsersBatch(p, key)
}

func (a *usersServerAdapter) LoadUsersBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return LoadUsersBatchMany(p, keys)
}

type usersClientAdapter struct {
	client UsersClient
}

func (a *usersClientAdapter) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	return a.client.GetUsers(ctx, req)
}

func (a *usersClientAdapter) GetUserProfile(ctx context.Context, req *GetUserProfileRequest) (*UserProfile, error) {
	return a.client.GetUserProfile(ctx, req)
}

func (a *usersClientAdapter) LoadUsers(ctx context.Context, req *pg.BatchRequest) (*UsersBatchResponse, error) {
	return a.client.LoadUsers(ctx, req)
}

func (a *usersClientAdapter) LoadUsersBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return LoadUsersBatch(p, key)
}

func (a *usersClientAdapter) LoadUsersBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return LoadUsersBatchMany(p, keys)
}

// Users returns a unified UsersInstance that works with both clients and services
// Returns nil if neither client nor service is configured
func (m *UsersModule) Users() UsersInstance {
	if m.usersClient != nil {
		return &usersClientAdapter{client: m.usersClient}
	}
	if m.usersService != nil {
		return &usersServerAdapter{server: m.usersService}
	}
	return nil
}

// Backward compatibility layer for v0.51.7 API
// All functions below are deprecated and will be removed in a future version.
// Please migrate to the module-based API using NewUsersModule()

var defaultModule *UsersModule

// UsersClientInstance provides a unified Users client interface
// Deprecated: Use NewUsersModule().Users() instead
var UsersClientInstance UsersInstance

// SetDefaultModule allows you to set a custom module instance as the default
// for use with deprecated package-level functions.
// This allows you to configure a module once and have all deprecated functions use it.
// Example:
//
//	module := NewUsersModule(WithDialOptions(...))
//	SetDefaultModule(module)
//	// Now all deprecated Init(), WithLoaders(), etc. will use your module
func SetDefaultModule(module *UsersModule) {
	defaultModule = module
}

func getDefaultModule() *UsersModule {
	if defaultModule == nil {
		defaultModule = NewUsersModule()
	}
	return defaultModule
}

func init() {
	// Initialize UsersClientInstance with lazy-loading adapter
	UsersClientInstance = &usersClientAdapter{client: nil}
}

// UsersInit initializes the Users service.
// Deprecated: Use NewUsersModule() and configure with WithModuleUsersClient() or WithModuleUsersService() instead.
func UsersInit(ctx context.Context, opts ...UsersModuleOption) (context.Context, []*gql.Field) {
	// Apply options to default module
	m := getDefaultModule()
	for _, opt := range opts {
		opt(m)
	}

	// Get fields from the module
	fields := m.Fields()

	// Register loaders in context
	ctx = m.WithLoaders(ctx)

	// Convert fields map to slice for this service only
	var serviceFields []*gql.Field
	servicePrefix := "users_"
	for name, field := range fields {
		if strings.HasPrefix(name, servicePrefix) {
			serviceFields = append(serviceFields, field)
		}
	}

	return ctx, serviceFields
}

// UsersWithLoaders registers dataloaders for the Users service into the context.
// Deprecated: Use NewUsersModule().WithLoaders(ctx) instead.
func UsersWithLoaders(ctx context.Context) context.Context {
	return getDefaultModule().WithLoaders(ctx)
}

// WithLoaders registers all dataloaders from all services into the context.
// Deprecated: Use NewUsersModule().WithLoaders(ctx) instead.
func WithLoaders(ctx context.Context) context.Context {
	return getDefaultModule().WithLoaders(ctx)
}

// Fields returns all GraphQL query/mutation fields from all services.
// Deprecated: Use NewUsersModule().Fields() instead.
func Fields() gql.Fields {
	return getDefaultModule().Fields()
}
