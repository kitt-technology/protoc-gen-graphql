package cases

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var fieldInits []func(...grpc.DialOption)

func Fields(opts ...grpc.DialOption) []*gql.Field {
	for _, fieldInit := range fieldInits {
		fieldInit(opts...)
	}
	return fields
}

var fields []*gql.Field

var GenreGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "Genre",
	Values: gql.EnumValueConfigMap{
		"Biography": &gql.EnumValueConfig{
			Value: Genre(1),
		},
		"Fiction": &gql.EnumValueConfig{
			Value: Genre(0),
		},
	},
})

var GenreGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "Genre",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(Genre).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var GetBooksRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "BooksRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"hardbackOnly": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*GetBooksRequest) == nil || p.Source.(*GetBooksRequest).HardbackOnly == nil {
					return nil, nil
				}
				return p.Source.(*GetBooksRequest).HardbackOnly.Value, nil
			},
		},
		"price": &gql.Field{
			Type: gql.Float,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*GetBooksRequest) == nil || p.Source.(*GetBooksRequest).Price == nil {
					return nil, nil
				}
				return p.Source.(*GetBooksRequest).Price.Value, nil
			},
		},
		"genres": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
		},
		"releasedAfter": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"pagination": &gql.Field{
			Type: PaginationOptionsGraphqlType,
		},
		"filters": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(FilterGraphqlType)),
		},
	},
})

var GetBooksRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BooksRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"hardbackOnly": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"price": &gql.InputObjectFieldConfig{
			Type: gql.Float,
		},
		"genres": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
		},
		"releasedAfter": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"pagination": &gql.InputObjectFieldConfig{
			Type: PaginationOptionsGraphqlInputType,
		},
		"filters": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(FilterGraphqlInputType)),
		},
	},
})

var GetBooksRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
	"hardbackOnly": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"price": &gql.ArgumentConfig{
		Type: gql.Float,
	},
	"genres": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GenreGraphqlEnum)),
	},
	"releasedAfter": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"pagination": &gql.ArgumentConfig{
		Type: PaginationOptionsGraphqlInputType,
	},
	"filters": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(FilterGraphqlInputType)),
	},
}

func GetBooksRequestFromArgs(args map[string]interface{}) *GetBooksRequest {
	return GetBooksRequestInstanceFromArgs(&GetBooksRequest{}, args)
}

func GetBooksRequestInstanceFromArgs(objectFromArgs *GetBooksRequest, args map[string]interface{}) *GetBooksRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		ids := make([]string, 0)

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	if args["hardbackOnly"] != nil {
		val := args["hardbackOnly"]
		objectFromArgs.HardbackOnly = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = wrapperspb.Float(float32(val.(float64)))
	}
	if args["genres"] != nil {
		genresInterfaceList := args["genres"].([]interface{})
		genres := make([]Genre, 0)

		for _, val := range genresInterfaceList {
			itemResolved := val.(Genre)
			genres = append(genres, itemResolved)
		}
		objectFromArgs.Genres = genres
	}
	if args["releasedAfter"] != nil {
		val := args["releasedAfter"]
		objectFromArgs.ReleasedAfter = pg.ToTimestamp(val)
	}
	if args["pagination"] != nil {
		val := args["pagination"]
		objectFromArgs.Pagination = PaginationOptionsFromArgs(val.(map[string]interface{}))
	}
	if args["filters"] != nil {
		filtersInterfaceList := args["filters"].([]interface{})
		filters := make([]*Filter, 0)

		for _, val := range filtersInterfaceList {
			itemResolved := FilterFromArgs(val.(map[string]interface{}))
			filters = append(filters, itemResolved)
		}
		objectFromArgs.Filters = filters
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksRequest) FromArgs(args map[string]interface{}) {
	GetBooksRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksRequest) XXX_GraphqlType() *gql.Object {
	return GetBooksRequestGraphqlType
}

func (msg *GetBooksRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksRequestGraphqlArgs
}

func (msg *GetBooksRequest) XXX_Package() string {
	return "books"
}

var PaginationOptionsGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "PaginationOptions",
	Fields: gql.Fields{
		"page": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"perPage": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var PaginationOptionsGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "PaginationOptionsInput",
	Fields: gql.InputObjectConfigFieldMap{
		"page": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"perPage": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var PaginationOptionsGraphqlArgs = gql.FieldConfigArgument{
	"page": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"perPage": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func PaginationOptionsFromArgs(args map[string]interface{}) *PaginationOptions {
	return PaginationOptionsInstanceFromArgs(&PaginationOptions{}, args)
}

func PaginationOptionsInstanceFromArgs(objectFromArgs *PaginationOptions, args map[string]interface{}) *PaginationOptions {
	if args["page"] != nil {
		val := args["page"]
		objectFromArgs.Page = int32(val.(int))
	}
	if args["perPage"] != nil {
		val := args["perPage"]
		objectFromArgs.PerPage = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *PaginationOptions) FromArgs(args map[string]interface{}) {
	PaginationOptionsInstanceFromArgs(objectFromArgs, args)
}

func (msg *PaginationOptions) XXX_GraphqlType() *gql.Object {
	return PaginationOptionsGraphqlType
}

func (msg *PaginationOptions) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return PaginationOptionsGraphqlArgs
}

func (msg *PaginationOptions) XXX_Package() string {
	return "books"
}

var FilterGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Filter",
	Fields: gql.Fields{
		"query": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var FilterGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "FilterInput",
	Fields: gql.InputObjectConfigFieldMap{
		"query": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var FilterGraphqlArgs = gql.FieldConfigArgument{
	"query": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func FilterFromArgs(args map[string]interface{}) *Filter {
	return FilterInstanceFromArgs(&Filter{}, args)
}

func FilterInstanceFromArgs(objectFromArgs *Filter, args map[string]interface{}) *Filter {
	if args["query"] != nil {
		val := args["query"]
		objectFromArgs.Query = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *Filter) FromArgs(args map[string]interface{}) {
	FilterInstanceFromArgs(objectFromArgs, args)
}

func (msg *Filter) XXX_GraphqlType() *gql.Object {
	return FilterGraphqlType
}

func (msg *Filter) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return FilterGraphqlArgs
}

func (msg *Filter) XXX_Package() string {
	return "books"
}

var GetBooksResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksResponse",
	Fields: gql.Fields{
		"books": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlType)),
		},
	},
})

var GetBooksResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"books": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
		},
	},
})

var GetBooksResponseGraphqlArgs = gql.FieldConfigArgument{
	"books": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
	},
}

func GetBooksResponseFromArgs(args map[string]interface{}) *GetBooksResponse {
	return GetBooksResponseInstanceFromArgs(&GetBooksResponse{}, args)
}

func GetBooksResponseInstanceFromArgs(objectFromArgs *GetBooksResponse, args map[string]interface{}) *GetBooksResponse {
	if args["books"] != nil {
		booksInterfaceList := args["books"].([]interface{})
		books := make([]*Book, 0)

		for _, val := range booksInterfaceList {
			itemResolved := BookFromArgs(val.(map[string]interface{}))
			books = append(books, itemResolved)
		}
		objectFromArgs.Books = books
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksResponse) FromArgs(args map[string]interface{}) {
	GetBooksResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksResponseGraphqlType
}

func (msg *GetBooksResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksResponseGraphqlArgs
}

func (msg *GetBooksResponse) XXX_Package() string {
	return "books"
}

var GetBooksBatchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksBatchRequest",
	Fields: gql.Fields{
		"reqs": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlType)),
		},
	},
})

var GetBooksBatchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksBatchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"reqs": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlInputType)),
		},
	},
})

var GetBooksBatchRequestGraphqlArgs = gql.FieldConfigArgument{
	"reqs": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GetBooksRequestGraphqlInputType)),
	},
}

func GetBooksBatchRequestFromArgs(args map[string]interface{}) *GetBooksBatchRequest {
	return GetBooksBatchRequestInstanceFromArgs(&GetBooksBatchRequest{}, args)
}

func GetBooksBatchRequestInstanceFromArgs(objectFromArgs *GetBooksBatchRequest, args map[string]interface{}) *GetBooksBatchRequest {
	if args["reqs"] != nil {
		reqsInterfaceList := args["reqs"].([]interface{})
		reqs := make([]*GetBooksRequest, 0)

		for _, val := range reqsInterfaceList {
			itemResolved := GetBooksRequestFromArgs(val.(map[string]interface{}))
			reqs = append(reqs, itemResolved)
		}
		objectFromArgs.Reqs = reqs
	}
	return objectFromArgs
}

func (objectFromArgs *GetBooksBatchRequest) FromArgs(args map[string]interface{}) {
	GetBooksBatchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksBatchRequest) XXX_GraphqlType() *gql.Object {
	return GetBooksBatchRequestGraphqlType
}

func (msg *GetBooksBatchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksBatchRequestGraphqlArgs
}

func (msg *GetBooksBatchRequest) XXX_Package() string {
	return "books"
}

var GetBooksBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var GetBooksBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetBooksBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetBooksBatchResponseFromArgs(args map[string]interface{}) *GetBooksBatchResponse {
	return GetBooksBatchResponseInstanceFromArgs(&GetBooksBatchResponse{}, args)
}

func GetBooksBatchResponseInstanceFromArgs(objectFromArgs *GetBooksBatchResponse, args map[string]interface{}) *GetBooksBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *GetBooksBatchResponse) FromArgs(args map[string]interface{}) {
	GetBooksBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksBatchResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksBatchResponseGraphqlType
}

func (msg *GetBooksBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksBatchResponseGraphqlArgs
}

func (msg *GetBooksBatchResponse) XXX_Package() string {
	return "books"
}

var GetBooksByAuthorResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetBooksByAuthorResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var GetBooksByAuthorResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetBooksByAuthorResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetBooksByAuthorResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetBooksByAuthorResponseFromArgs(args map[string]interface{}) *GetBooksByAuthorResponse {
	return GetBooksByAuthorResponseInstanceFromArgs(&GetBooksByAuthorResponse{}, args)
}

func GetBooksByAuthorResponseInstanceFromArgs(objectFromArgs *GetBooksByAuthorResponse, args map[string]interface{}) *GetBooksByAuthorResponse {
	return objectFromArgs
}

func (objectFromArgs *GetBooksByAuthorResponse) FromArgs(args map[string]interface{}) {
	GetBooksByAuthorResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetBooksByAuthorResponse) XXX_GraphqlType() *gql.Object {
	return GetBooksByAuthorResponseGraphqlType
}

func (msg *GetBooksByAuthorResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetBooksByAuthorResponseGraphqlArgs
}

func (msg *GetBooksByAuthorResponse) XXX_Package() string {
	return "books"
}

var BooksByAuthorGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "BooksByAuthor",
	Fields: gql.Fields{
		"results": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlType)),
		},
	},
})

var BooksByAuthorGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BooksByAuthorInput",
	Fields: gql.InputObjectConfigFieldMap{
		"results": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
		},
	},
})

var BooksByAuthorGraphqlArgs = gql.FieldConfigArgument{
	"results": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(BookGraphqlInputType)),
	},
}

func BooksByAuthorFromArgs(args map[string]interface{}) *BooksByAuthor {
	return BooksByAuthorInstanceFromArgs(&BooksByAuthor{}, args)
}

func BooksByAuthorInstanceFromArgs(objectFromArgs *BooksByAuthor, args map[string]interface{}) *BooksByAuthor {
	if args["results"] != nil {
		resultsInterfaceList := args["results"].([]interface{})
		results := make([]*Book, 0)

		for _, val := range resultsInterfaceList {
			itemResolved := BookFromArgs(val.(map[string]interface{}))
			results = append(results, itemResolved)
		}
		objectFromArgs.Results = results
	}
	return objectFromArgs
}

func (objectFromArgs *BooksByAuthor) FromArgs(args map[string]interface{}) {
	BooksByAuthorInstanceFromArgs(objectFromArgs, args)
}

func (msg *BooksByAuthor) XXX_GraphqlType() *gql.Object {
	return BooksByAuthorGraphqlType
}

func (msg *BooksByAuthor) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return BooksByAuthorGraphqlArgs
}

func (msg *BooksByAuthor) XXX_Package() string {
	return "books"
}

var BookGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Book",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"authorId": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"genre": &gql.Field{
			Type: GenreGraphqlEnum,
		},
		"releaseDate": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"price": &gql.Field{
			Type: common_example.MoneyGraphqlType,
		},
		"priceTwo": &gql.Field{
			Type: MoneyGraphqlType,
		},
		"isSigned": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*Book) == nil || p.Source.(*Book).IsSigned == nil {
					return nil, nil
				}
				return p.Source.(*Book).IsSigned.Value, nil
			},
		},
		"notes": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*Book) == nil || p.Source.(*Book).Notes == nil {
					return nil, nil
				}
				return p.Source.(*Book).Notes.Value, nil
			},
		},
		"historicPrices": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(common_example.MoneyGraphqlType)),
		},
		"pages": &gql.Field{
			Type: gql.Int,
		},
	},
})

var BookGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "BookInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"authorId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"genre": &gql.InputObjectFieldConfig{
			Type: GenreGraphqlEnum,
		},
		"releaseDate": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"price": &gql.InputObjectFieldConfig{
			Type: common_example.MoneyGraphqlInputType,
		},
		"priceTwo": &gql.InputObjectFieldConfig{
			Type: MoneyGraphqlInputType,
		},
		"isSigned": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"notes": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
		"historicPrices": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(common_example.MoneyGraphqlInputType)),
		},
		"pages": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
	},
})

var BookGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"authorId": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"genre": &gql.ArgumentConfig{
		Type: GenreGraphqlEnum,
	},
	"releaseDate": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"price": &gql.ArgumentConfig{
		Type: common_example.MoneyGraphqlInputType,
	},
	"priceTwo": &gql.ArgumentConfig{
		Type: MoneyGraphqlInputType,
	},
	"isSigned": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"notes": &gql.ArgumentConfig{
		Type: gql.String,
	},
	"historicPrices": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(common_example.MoneyGraphqlInputType)),
	},
	"pages": &gql.ArgumentConfig{
		Type: gql.Int,
	},
}

func BookFromArgs(args map[string]interface{}) *Book {
	return BookInstanceFromArgs(&Book{}, args)
}

func BookInstanceFromArgs(objectFromArgs *Book, args map[string]interface{}) *Book {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["name"] != nil {
		val := args["name"]
		objectFromArgs.Name = string(val.(string))
	}
	if args["authorId"] != nil {
		val := args["authorId"]
		objectFromArgs.AuthorId = string(val.(string))
	}
	if args["genre"] != nil {
		val := args["genre"]
		ptr := val.(Genre)
		objectFromArgs.Genre = &ptr
	}
	if args["releaseDate"] != nil {
		val := args["releaseDate"]
		objectFromArgs.ReleaseDate = pg.ToTimestamp(val)
	}
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = common_example.MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["priceTwo"] != nil {
		val := args["priceTwo"]
		objectFromArgs.PriceTwo = MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["isSigned"] != nil {
		val := args["isSigned"]
		objectFromArgs.IsSigned = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["notes"] != nil {
		val := args["notes"]
		objectFromArgs.Notes = wrapperspb.String(string(val.(string)))
	}
	if args["historicPrices"] != nil {
		historicPricesInterfaceList := args["historicPrices"].([]interface{})
		historicPrices := make([]*common_example.Money, 0)

		for _, val := range historicPricesInterfaceList {
			itemResolved := common_example.MoneyFromArgs(val.(map[string]interface{}))
			historicPrices = append(historicPrices, itemResolved)
		}
		objectFromArgs.HistoricPrices = historicPrices
	}
	if args["pages"] != nil {
		val := args["pages"]
		ptr := int32(val.(int))
		objectFromArgs.Pages = &ptr
	}
	return objectFromArgs
}

func (objectFromArgs *Book) FromArgs(args map[string]interface{}) {
	BookInstanceFromArgs(objectFromArgs, args)
}

func (msg *Book) XXX_GraphqlType() *gql.Object {
	return BookGraphqlType
}

func (msg *Book) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return BookGraphqlArgs
}

func (msg *Book) XXX_Package() string {
	return "books"
}

var MoneyGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Money",
	Fields: gql.Fields{
		"price": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var MoneyGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "MoneyInput",
	Fields: gql.InputObjectConfigFieldMap{
		"price": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
})

var MoneyGraphqlArgs = gql.FieldConfigArgument{
	"price": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
}

func MoneyFromArgs(args map[string]interface{}) *Money {
	return MoneyInstanceFromArgs(&Money{}, args)
}

func MoneyInstanceFromArgs(objectFromArgs *Money, args map[string]interface{}) *Money {
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *Money) FromArgs(args map[string]interface{}) {
	MoneyInstanceFromArgs(objectFromArgs, args)
}

func (msg *Money) XXX_GraphqlType() *gql.Object {
	return MoneyGraphqlType
}

func (msg *Money) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return MoneyGraphqlArgs
}

func (msg *Money) XXX_Package() string {
	return "books"
}

var BooksClientInstance BooksClient

func init() {
	fieldInits = append(fieldInits, func(opts ...grpc.DialOption) {
		BooksClientInstance = NewBooksClient(pg.GrpcConnection("localhost:50051", opts...))
	})
	fields = append(fields, &gql.Field{
		Name: "books_GetBooks",
		Type: GetBooksResponseGraphqlType,
		Args: GetBooksRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			return BooksClientInstance.GetBooks(p.Context, GetBooksRequestFromArgs(p.Args))
		},
	})

}

func WithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "GetBooksByAuthorLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			resp, err := BooksClientInstance.GetBooksByAuthor(ctx, &pg.BatchRequest{
				Keys: keys.Keys(),
			})

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *BooksByAuthor
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	ctx = context.WithValue(ctx, "GetBooksBatchLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var requests []*GetBooksRequest
			for _, key := range keys {
				requests = append(requests, key.(*GetBooksRequestKey).GetBooksRequest)
			}
			resp, err := BooksClientInstance.GetBooksBatch(ctx, &GetBooksBatchRequest{
				Reqs: requests,
			})

			if err != nil {
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else if err, ok := resp.Errors[key]; ok {
					results = append(results, &dataloader.Result{Error: errors.New(err)})
				} else {
					var empty *GetBooksResponse
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	return ctx
}

func GetBooksByAuthor(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksByAuthorLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*BooksByAuthor), nil
	}, nil
}

func GetBooksByAuthorMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksByAuthorLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksByAuthorLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*BooksByAuthor
		for _, res := range resSlice {
			results = append(results, res.(*BooksByAuthor))
		}

		return results, nil
	}, nil
}

type GetBooksRequestKey struct {
	*GetBooksRequest
}

func (key *GetBooksRequestKey) String() string {
	return pg.ProtoKey(key)
}

func (key *GetBooksRequestKey) Raw() interface{} {
	return key
}

func GetBooksBatch(p gql.ResolveParams, key *GetBooksRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetBooksBatchLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetBooksBatchLoader").(*dataloader.Loader)
	default:
		panic("Please call books.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, &GetBooksRequestKey{key})
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*GetBooksResponse), nil
	}, nil
}
