package users

import (
	gql "github.com/graphql-go/graphql"
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql/language/ast"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
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

var UsersClientInstance UsersClient
var UsersServiceInstance UsersServer
var UsersDialOpts []grpc.DialOption

type UsersOption func(*UsersConfig)

type UsersConfig struct {
	service  UsersServer
	client   UsersClient
	dialOpts []grpc.DialOption
}

// WithService sets the service implementation for direct calls (no gRPC)
func WithService(service UsersServer) UsersOption {
	return func(cfg *UsersConfig) {
		cfg.service = service
	}
}

// WithClient sets the gRPC client for remote calls
func WithClient(client UsersClient) UsersOption {
	return func(cfg *UsersConfig) {
		cfg.client = client
	}
}

// WithDialOptions sets the dial options for the gRPC client
func WithDialOptions(opts ...grpc.DialOption) UsersOption {
	return func(cfg *UsersConfig) {
		cfg.dialOpts = opts
	}
}

func Init(ctx context.Context, opts ...UsersOption) (context.Context, []*gql.Field) {
	cfg := &UsersConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	UsersServiceInstance = cfg.service
	UsersClientInstance = cfg.client
	UsersDialOpts = cfg.dialOpts

	var fields []*gql.Field
	fields = append(fields, &gql.Field{
		Name: "users_GetUsers",
		Type: GetUsersResponseGraphqlType,
		Args: GetUsersRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if UsersServiceInstance != nil {
				return UsersServiceInstance.GetUsers(p.Context, GetUsersRequestFromArgs(p.Args))
			}
			if UsersClientInstance == nil {
				UsersClientInstance = getUsersClient()
			}
			return UsersClientInstance.GetUsers(p.Context, GetUsersRequestFromArgs(p.Args))
		},
	})

	fields = append(fields, &gql.Field{
		Name: "users_GetUserProfile",
		Type: UserProfileGraphqlType,
		Args: GetUserProfileRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if UsersServiceInstance != nil {
				return UsersServiceInstance.GetUserProfile(p.Context, GetUserProfileRequestFromArgs(p.Args))
			}
			if UsersClientInstance == nil {
				UsersClientInstance = getUsersClient()
			}
			return UsersClientInstance.GetUserProfile(p.Context, GetUserProfileRequestFromArgs(p.Args))
		},
	})

	ctx = UsersWithLoaders(ctx)

	return ctx, fields
}

func getUsersClient() UsersClient {
	host := "localhost:50052"
	envHost := os.Getenv("SERVICE_HOST")
	if envHost != "" {
		host = envHost
	}
	return NewUsersClient(pg.GrpcConnection(host, UsersDialOpts...))
}

// SetUsersService sets the service implementation for direct calls (no gRPC)
func SetUsersService(service UsersServer) {
	UsersServiceInstance = service
}

// SetUsersClient sets the gRPC client for remote calls
func SetUsersClient(client UsersClient) {
	UsersClientInstance = client
}

func UsersWithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "LoadUsersLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *UsersBatchResponse
			var err error
			if UsersServiceInstance != nil {
				resp, err = UsersServiceInstance.LoadUsers(ctx, &pg.BatchRequest{
					Keys: keys.Keys(),
				})
			} else {
				if UsersClientInstance == nil {
					UsersClientInstance = getUsersClient()
				}
				resp, err = UsersClientInstance.LoadUsers(ctx, &pg.BatchRequest{
					Keys: keys.Keys(),
				})
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

func LoadUsers(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
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

func LoadUsersMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
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
