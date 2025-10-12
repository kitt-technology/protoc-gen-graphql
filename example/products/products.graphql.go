package products

import (
	gql "github.com/graphql-go/graphql"
	"context"
	"github.com/graph-gophers/dataloader"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	"google.golang.org/protobuf/types/known/wrapperspb"
	pg "github.com/kitt-technology/protoc-gen-graphql/graphql"
	"strings"
)

var CategoryGraphqlEnum = gql.NewEnum(gql.EnumConfig{
	Name: "Category",
	Values: gql.EnumValueConfigMap{
		"BOOKS": &gql.EnumValueConfig{
			Value: Category(2),
		},
		"CLOTHING": &gql.EnumValueConfig{
			Value: Category(1),
		},
		"ELECTRONICS": &gql.EnumValueConfig{
			Value: Category(0),
		},
		"HOME_GARDEN": &gql.EnumValueConfig{
			Value: Category(3),
		},
		"SPORTS": &gql.EnumValueConfig{
			Value: Category(4),
		},
		"TOYS": &gql.EnumValueConfig{
			Value: Category(5),
		},
	},
})

var CategoryGraphqlType = gql.NewScalar(gql.ScalarConfig{
	Name: "Category",
	ParseValue: func(value interface{}) interface{} {
		return nil
	},
	Serialize: func(value interface{}) interface{} {
		return value.(Category).String()
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return nil
	},
})

var GetProductsRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "ProductsRequest",
	Fields: gql.Fields{
		"ids": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"categories": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(CategoryGraphqlEnum)),
		},
		"inStockOnly": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*GetProductsRequest) == nil || p.Source.(*GetProductsRequest).InStockOnly == nil {
					return nil, nil
				}
				return p.Source.(*GetProductsRequest).InStockOnly.Value, nil
			},
		},
		"availableAfter": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
	},
})
var GetProductsRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "ProductsRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"ids": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(gql.String)),
		},
		"categories": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(CategoryGraphqlEnum)),
		},
		"inStockOnly": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"availableAfter": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
	},
})

var GetProductsRequestGraphqlArgs = gql.FieldConfigArgument{
	"ids": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(gql.String)),
	},
	"categories": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(CategoryGraphqlEnum)),
	},
	"inStockOnly": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"availableAfter": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
}

func GetProductsRequestFromArgs(args map[string]interface{}) *GetProductsRequest {
	return GetProductsRequestInstanceFromArgs(&GetProductsRequest{}, args)
}

func GetProductsRequestInstanceFromArgs(objectFromArgs *GetProductsRequest, args map[string]interface{}) *GetProductsRequest {
	if args["ids"] != nil {
		idsInterfaceList := args["ids"].([]interface{})
		ids := make([]string, 0)

		for _, val := range idsInterfaceList {
			itemResolved := string(val.(string))
			ids = append(ids, itemResolved)
		}
		objectFromArgs.Ids = ids
	}
	if args["categories"] != nil {
		categoriesInterfaceList := args["categories"].([]interface{})
		categories := make([]Category, 0)

		for _, val := range categoriesInterfaceList {
			itemResolved := val.(Category)
			categories = append(categories, itemResolved)
		}
		objectFromArgs.Categories = categories
	}
	if args["inStockOnly"] != nil {
		val := args["inStockOnly"]
		objectFromArgs.InStockOnly = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["availableAfter"] != nil {
		val := args["availableAfter"]
		objectFromArgs.AvailableAfter = pg.ToTimestamp(val)
	}
	return objectFromArgs
}

func (objectFromArgs *GetProductsRequest) FromArgs(args map[string]interface{}) {
	GetProductsRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetProductsRequest) XXX_GraphqlType() *gql.Object {
	return GetProductsRequestGraphqlType
}

func (msg *GetProductsRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetProductsRequestGraphqlArgs
}

func (msg *GetProductsRequest) XXX_Package() string {
	return "products"
}

var GetProductsResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetProductsResponse",
	Fields: gql.Fields{
		"products": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlType)),
		},
		"pageInfo": &gql.Field{
			Type: pg.PageInfoGraphqlType,
		},
	},
})
var GetProductsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetProductsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"products": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
		},
		"pageInfo": &gql.InputObjectFieldConfig{
			Type: pg.PageInfoGraphqlInputType,
		},
	},
})

var GetProductsResponseGraphqlArgs = gql.FieldConfigArgument{
	"products": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
	},
	"pageInfo": &gql.ArgumentConfig{
		Type: pg.PageInfoGraphqlInputType,
	},
}

func GetProductsResponseFromArgs(args map[string]interface{}) *GetProductsResponse {
	return GetProductsResponseInstanceFromArgs(&GetProductsResponse{}, args)
}

func GetProductsResponseInstanceFromArgs(objectFromArgs *GetProductsResponse, args map[string]interface{}) *GetProductsResponse {
	if args["products"] != nil {
		productsInterfaceList := args["products"].([]interface{})
		products := make([]*Product, 0)

		for _, val := range productsInterfaceList {
			itemResolved := ProductFromArgs(val.(map[string]interface{}))
			products = append(products, itemResolved)
		}
		objectFromArgs.Products = products
	}
	if args["pageInfo"] != nil {
		val := args["pageInfo"]
		objectFromArgs.PageInfo = pg.PageInfoFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *GetProductsResponse) FromArgs(args map[string]interface{}) {
	GetProductsResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetProductsResponse) XXX_GraphqlType() *gql.Object {
	return GetProductsResponseGraphqlType
}

func (msg *GetProductsResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetProductsResponseGraphqlArgs
}

func (msg *GetProductsResponse) XXX_Package() string {
	return "products"
}

var GetProductsByCategoryResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetProductsByCategoryResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})
var GetProductsByCategoryResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetProductsByCategoryResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetProductsByCategoryResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetProductsByCategoryResponseFromArgs(args map[string]interface{}) *GetProductsByCategoryResponse {
	return GetProductsByCategoryResponseInstanceFromArgs(&GetProductsByCategoryResponse{}, args)
}

func GetProductsByCategoryResponseInstanceFromArgs(objectFromArgs *GetProductsByCategoryResponse, args map[string]interface{}) *GetProductsByCategoryResponse {
	return objectFromArgs
}

func (objectFromArgs *GetProductsByCategoryResponse) FromArgs(args map[string]interface{}) {
	GetProductsByCategoryResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetProductsByCategoryResponse) XXX_GraphqlType() *gql.Object {
	return GetProductsByCategoryResponseGraphqlType
}

func (msg *GetProductsByCategoryResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetProductsByCategoryResponseGraphqlArgs
}

func (msg *GetProductsByCategoryResponse) XXX_Package() string {
	return "products"
}

var LoadProductsRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "LoadProductsRequest",
	Fields: gql.Fields{
		"reqs": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlType)),
		},
	},
})
var LoadProductsRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "LoadProductsRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"reqs": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlInputType)),
		},
	},
})

var LoadProductsRequestGraphqlArgs = gql.FieldConfigArgument{
	"reqs": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlInputType)),
	},
}

func LoadProductsRequestFromArgs(args map[string]interface{}) *LoadProductsRequest {
	return LoadProductsRequestInstanceFromArgs(&LoadProductsRequest{}, args)
}

func LoadProductsRequestInstanceFromArgs(objectFromArgs *LoadProductsRequest, args map[string]interface{}) *LoadProductsRequest {
	if args["reqs"] != nil {
		reqsInterfaceList := args["reqs"].([]interface{})
		reqs := make([]*GetProductsRequest, 0)

		for _, val := range reqsInterfaceList {
			itemResolved := GetProductsRequestFromArgs(val.(map[string]interface{}))
			reqs = append(reqs, itemResolved)
		}
		objectFromArgs.Reqs = reqs
	}
	return objectFromArgs
}

func (objectFromArgs *LoadProductsRequest) FromArgs(args map[string]interface{}) {
	LoadProductsRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *LoadProductsRequest) XXX_GraphqlType() *gql.Object {
	return LoadProductsRequestGraphqlType
}

func (msg *LoadProductsRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return LoadProductsRequestGraphqlArgs
}

func (msg *LoadProductsRequest) XXX_Package() string {
	return "products"
}

var LoadProductsResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "LoadProductsResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})
var LoadProductsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "LoadProductsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var LoadProductsResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func LoadProductsResponseFromArgs(args map[string]interface{}) *LoadProductsResponse {
	return LoadProductsResponseInstanceFromArgs(&LoadProductsResponse{}, args)
}

func LoadProductsResponseInstanceFromArgs(objectFromArgs *LoadProductsResponse, args map[string]interface{}) *LoadProductsResponse {
	return objectFromArgs
}

func (objectFromArgs *LoadProductsResponse) FromArgs(args map[string]interface{}) {
	LoadProductsResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *LoadProductsResponse) XXX_GraphqlType() *gql.Object {
	return LoadProductsResponseGraphqlType
}

func (msg *LoadProductsResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return LoadProductsResponseGraphqlArgs
}

func (msg *LoadProductsResponse) XXX_Package() string {
	return "products"
}

var SearchProductsRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SearchProductsRequest",
	Fields: gql.Fields{
		"query": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"limit": &gql.Field{
			Type: gql.Int,
		},
		"cursor": &gql.Field{
			Type: gql.String,
		},
	},
})
var SearchProductsRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SearchProductsRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"query": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"limit": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
		"cursor": &gql.InputObjectFieldConfig{
			Type: gql.String,
		},
	},
})

var SearchProductsRequestGraphqlArgs = gql.FieldConfigArgument{
	"query": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"limit": &gql.ArgumentConfig{
		Type: gql.Int,
	},
	"cursor": &gql.ArgumentConfig{
		Type: gql.String,
	},
}

func SearchProductsRequestFromArgs(args map[string]interface{}) *SearchProductsRequest {
	return SearchProductsRequestInstanceFromArgs(&SearchProductsRequest{}, args)
}

func SearchProductsRequestInstanceFromArgs(objectFromArgs *SearchProductsRequest, args map[string]interface{}) *SearchProductsRequest {
	if args["query"] != nil {
		val := args["query"]
		objectFromArgs.Query = string(val.(string))
	}
	if args["limit"] != nil {
		val := args["limit"]
		objectFromArgs.Limit = int32(val.(int))
	}
	if args["cursor"] != nil {
		val := args["cursor"]
		objectFromArgs.Cursor = string(val.(string))
	}
	return objectFromArgs
}

func (objectFromArgs *SearchProductsRequest) FromArgs(args map[string]interface{}) {
	SearchProductsRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *SearchProductsRequest) XXX_GraphqlType() *gql.Object {
	return SearchProductsRequestGraphqlType
}

func (msg *SearchProductsRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SearchProductsRequestGraphqlArgs
}

func (msg *SearchProductsRequest) XXX_Package() string {
	return "products"
}

var SearchProductsResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "SearchProductsResponse",
	Fields: gql.Fields{
		"products": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlType)),
		},
		"pageInfo": &gql.Field{
			Type: pg.PageInfoGraphqlType,
		},
	},
})
var SearchProductsResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "SearchProductsResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"products": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
		},
		"pageInfo": &gql.InputObjectFieldConfig{
			Type: pg.PageInfoGraphqlInputType,
		},
	},
})

var SearchProductsResponseGraphqlArgs = gql.FieldConfigArgument{
	"products": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
	},
	"pageInfo": &gql.ArgumentConfig{
		Type: pg.PageInfoGraphqlInputType,
	},
}

func SearchProductsResponseFromArgs(args map[string]interface{}) *SearchProductsResponse {
	return SearchProductsResponseInstanceFromArgs(&SearchProductsResponse{}, args)
}

func SearchProductsResponseInstanceFromArgs(objectFromArgs *SearchProductsResponse, args map[string]interface{}) *SearchProductsResponse {
	if args["products"] != nil {
		productsInterfaceList := args["products"].([]interface{})
		products := make([]*Product, 0)

		for _, val := range productsInterfaceList {
			itemResolved := ProductFromArgs(val.(map[string]interface{}))
			products = append(products, itemResolved)
		}
		objectFromArgs.Products = products
	}
	if args["pageInfo"] != nil {
		val := args["pageInfo"]
		objectFromArgs.PageInfo = pg.PageInfoFromArgs(val.(map[string]interface{}))
	}
	return objectFromArgs
}

func (objectFromArgs *SearchProductsResponse) FromArgs(args map[string]interface{}) {
	SearchProductsResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *SearchProductsResponse) XXX_GraphqlType() *gql.Object {
	return SearchProductsResponseGraphqlType
}

func (msg *SearchProductsResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return SearchProductsResponseGraphqlArgs
}

func (msg *SearchProductsResponse) XXX_Package() string {
	return "products"
}

var ProductsByCategoryGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "ProductsByCategory",
	Fields: gql.Fields{
		"results": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlType)),
		},
	},
})
var ProductsByCategoryGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "ProductsByCategoryInput",
	Fields: gql.InputObjectConfigFieldMap{
		"results": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
		},
	},
})

var ProductsByCategoryGraphqlArgs = gql.FieldConfigArgument{
	"results": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(ProductGraphqlInputType)),
	},
}

func ProductsByCategoryFromArgs(args map[string]interface{}) *ProductsByCategory {
	return ProductsByCategoryInstanceFromArgs(&ProductsByCategory{}, args)
}

func ProductsByCategoryInstanceFromArgs(objectFromArgs *ProductsByCategory, args map[string]interface{}) *ProductsByCategory {
	if args["results"] != nil {
		resultsInterfaceList := args["results"].([]interface{})
		results := make([]*Product, 0)

		for _, val := range resultsInterfaceList {
			itemResolved := ProductFromArgs(val.(map[string]interface{}))
			results = append(results, itemResolved)
		}
		objectFromArgs.Results = results
	}
	return objectFromArgs
}

func (objectFromArgs *ProductsByCategory) FromArgs(args map[string]interface{}) {
	ProductsByCategoryInstanceFromArgs(objectFromArgs, args)
}

func (msg *ProductsByCategory) XXX_GraphqlType() *gql.Object {
	return ProductsByCategoryGraphqlType
}

func (msg *ProductsByCategory) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return ProductsByCategoryGraphqlArgs
}

func (msg *ProductsByCategory) XXX_Package() string {
	return "products"
}

var ProductGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Product",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"description": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"category": &gql.Field{
			Type: CategoryGraphqlEnum,
		},
		"price": &gql.Field{
			Type: common_example.MoneyGraphqlType,
		},
		"inventory": &gql.Field{
			Type: InventoryGraphqlType,
		},
		"variants": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(ProductVariantGraphqlType)),
		},
		"createdAt": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"updatedAt": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
		"featured": &gql.Field{
			Type: gql.Boolean,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if p.Source.(*Product) == nil || p.Source.(*Product).Featured == nil {
					return nil, nil
				}
				return p.Source.(*Product).Featured.Value, nil
			},
		},
		"sellerId": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"imageUrls": &gql.Field{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"rating": &gql.Field{
			Type: gql.Float,
		},
		"reviewCount": &gql.Field{
			Type: gql.Int,
		},
	},
})
var ProductGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "ProductInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"description": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"category": &gql.InputObjectFieldConfig{
			Type: CategoryGraphqlEnum,
		},
		"price": &gql.InputObjectFieldConfig{
			Type: common_example.MoneyGraphqlInputType,
		},
		"inventory": &gql.InputObjectFieldConfig{
			Type: InventoryGraphqlInputType,
		},
		"variants": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(ProductVariantGraphqlInputType)),
		},
		"createdAt": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"updatedAt": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
		"featured": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
		"sellerId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"imageUrls": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
		},
		"rating": &gql.InputObjectFieldConfig{
			Type: gql.Float,
		},
		"reviewCount": &gql.InputObjectFieldConfig{
			Type: gql.Int,
		},
	},
})

var ProductGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"description": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"category": &gql.ArgumentConfig{
		Type: CategoryGraphqlEnum,
	},
	"price": &gql.ArgumentConfig{
		Type: common_example.MoneyGraphqlInputType,
	},
	"inventory": &gql.ArgumentConfig{
		Type: InventoryGraphqlInputType,
	},
	"variants": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(ProductVariantGraphqlInputType)),
	},
	"createdAt": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"updatedAt": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
	"featured": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
	"sellerId": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"imageUrls": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(gql.String))),
	},
	"rating": &gql.ArgumentConfig{
		Type: gql.Float,
	},
	"reviewCount": &gql.ArgumentConfig{
		Type: gql.Int,
	},
}

func ProductFromArgs(args map[string]interface{}) *Product {
	return ProductInstanceFromArgs(&Product{}, args)
}

func ProductInstanceFromArgs(objectFromArgs *Product, args map[string]interface{}) *Product {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["name"] != nil {
		val := args["name"]
		objectFromArgs.Name = string(val.(string))
	}
	if args["description"] != nil {
		val := args["description"]
		objectFromArgs.Description = string(val.(string))
	}
	if args["category"] != nil {
		val := args["category"]
		objectFromArgs.Category = val.(Category)
	}
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = common_example.MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["inventory"] != nil {
		val := args["inventory"]
		objectFromArgs.Inventory = InventoryFromArgs(val.(map[string]interface{}))
	}
	if args["variants"] != nil {
		variantsInterfaceList := args["variants"].([]interface{})
		variants := make([]*ProductVariant, 0)

		for _, val := range variantsInterfaceList {
			itemResolved := ProductVariantFromArgs(val.(map[string]interface{}))
			variants = append(variants, itemResolved)
		}
		objectFromArgs.Variants = variants
	}
	if args["createdAt"] != nil {
		val := args["createdAt"]
		objectFromArgs.CreatedAt = pg.ToTimestamp(val)
	}
	if args["updatedAt"] != nil {
		val := args["updatedAt"]
		objectFromArgs.UpdatedAt = pg.ToTimestamp(val)
	}
	if args["featured"] != nil {
		val := args["featured"]
		objectFromArgs.Featured = wrapperspb.Bool(bool(val.(bool)))
	}
	if args["sellerId"] != nil {
		val := args["sellerId"]
		objectFromArgs.SellerId = string(val.(string))
	}
	if args["imageUrls"] != nil {
		imageUrlsInterfaceList := args["imageUrls"].([]interface{})
		imageUrls := make([]string, 0)

		for _, val := range imageUrlsInterfaceList {
			itemResolved := string(val.(string))
			imageUrls = append(imageUrls, itemResolved)
		}
		objectFromArgs.ImageUrls = imageUrls
	}
	if args["rating"] != nil {
		val := args["rating"]
		ptr := float32(val.(float64))
		objectFromArgs.Rating = &ptr
	}
	if args["reviewCount"] != nil {
		val := args["reviewCount"]
		ptr := int32(val.(int))
		objectFromArgs.ReviewCount = &ptr
	}
	return objectFromArgs
}

func (objectFromArgs *Product) FromArgs(args map[string]interface{}) {
	ProductInstanceFromArgs(objectFromArgs, args)
}

func (msg *Product) XXX_GraphqlType() *gql.Object {
	return ProductGraphqlType
}

func (msg *Product) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return ProductGraphqlArgs
}

func (msg *Product) XXX_Package() string {
	return "products"
}

var ProductVariantGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "ProductVariant",
	Fields: gql.Fields{
		"id": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"sku": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"price": &gql.Field{
			Type: common_example.MoneyGraphqlType,
		},
		"stockQuantity": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})
var ProductVariantGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "ProductVariantInput",
	Fields: gql.InputObjectConfigFieldMap{
		"id": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"sku": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"price": &gql.InputObjectFieldConfig{
			Type: common_example.MoneyGraphqlInputType,
		},
		"stockQuantity": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
	},
})

var ProductVariantGraphqlArgs = gql.FieldConfigArgument{
	"id": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"name": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"sku": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"price": &gql.ArgumentConfig{
		Type: common_example.MoneyGraphqlInputType,
	},
	"stockQuantity": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
}

func ProductVariantFromArgs(args map[string]interface{}) *ProductVariant {
	return ProductVariantInstanceFromArgs(&ProductVariant{}, args)
}

func ProductVariantInstanceFromArgs(objectFromArgs *ProductVariant, args map[string]interface{}) *ProductVariant {
	if args["id"] != nil {
		val := args["id"]
		objectFromArgs.Id = string(val.(string))
	}
	if args["name"] != nil {
		val := args["name"]
		objectFromArgs.Name = string(val.(string))
	}
	if args["sku"] != nil {
		val := args["sku"]
		objectFromArgs.Sku = string(val.(string))
	}
	if args["price"] != nil {
		val := args["price"]
		objectFromArgs.Price = common_example.MoneyFromArgs(val.(map[string]interface{}))
	}
	if args["stockQuantity"] != nil {
		val := args["stockQuantity"]
		objectFromArgs.StockQuantity = int32(val.(int))
	}
	return objectFromArgs
}

func (objectFromArgs *ProductVariant) FromArgs(args map[string]interface{}) {
	ProductVariantInstanceFromArgs(objectFromArgs, args)
}

func (msg *ProductVariant) XXX_GraphqlType() *gql.Object {
	return ProductVariantGraphqlType
}

func (msg *ProductVariant) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return ProductVariantGraphqlArgs
}

func (msg *ProductVariant) XXX_Package() string {
	return "products"
}

var InventoryGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Inventory",
	Fields: gql.Fields{
		"quantity": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"reserved": &gql.Field{
			Type: gql.NewNonNull(gql.Int),
		},
		"warehouseLocation": &gql.Field{
			Type: gql.NewNonNull(gql.String),
		},
		"lastRestocked": &gql.Field{
			Type: pg.TimestampGraphqlType,
		},
	},
})
var InventoryGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "InventoryInput",
	Fields: gql.InputObjectConfigFieldMap{
		"quantity": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"reserved": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"warehouseLocation": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"lastRestocked": &gql.InputObjectFieldConfig{
			Type: pg.TimestampGraphqlInputType,
		},
	},
})

var InventoryGraphqlArgs = gql.FieldConfigArgument{
	"quantity": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"reserved": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.Int),
	},
	"warehouseLocation": &gql.ArgumentConfig{
		Type: gql.NewNonNull(gql.String),
	},
	"lastRestocked": &gql.ArgumentConfig{
		Type: pg.TimestampGraphqlInputType,
	},
}

func InventoryFromArgs(args map[string]interface{}) *Inventory {
	return InventoryInstanceFromArgs(&Inventory{}, args)
}

func InventoryInstanceFromArgs(objectFromArgs *Inventory, args map[string]interface{}) *Inventory {
	if args["quantity"] != nil {
		val := args["quantity"]
		objectFromArgs.Quantity = int32(val.(int))
	}
	if args["reserved"] != nil {
		val := args["reserved"]
		objectFromArgs.Reserved = int32(val.(int))
	}
	if args["warehouseLocation"] != nil {
		val := args["warehouseLocation"]
		objectFromArgs.WarehouseLocation = string(val.(string))
	}
	if args["lastRestocked"] != nil {
		val := args["lastRestocked"]
		objectFromArgs.LastRestocked = pg.ToTimestamp(val)
	}
	return objectFromArgs
}

func (objectFromArgs *Inventory) FromArgs(args map[string]interface{}) {
	InventoryInstanceFromArgs(objectFromArgs, args)
}

func (msg *Inventory) XXX_GraphqlType() *gql.Object {
	return InventoryGraphqlType
}

func (msg *Inventory) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return InventoryGraphqlArgs
}

func (msg *Inventory) XXX_Package() string {
	return "products"
}
func GetProductsByCategory(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetProductsByCategoryLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetProductsByCategoryLoader").(*dataloader.Loader)
	default:
		panic("Please call products.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, dataloader.StringKey(key))
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*ProductsByCategory), nil
	}, nil
}

func GetProductsByCategoryMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetProductsByCategoryLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetProductsByCategoryLoader").(*dataloader.Loader)
	default:
		panic("Please call products.WithLoaders with the current context first")
	}

	thunk := loader.LoadMany(p.Context, dataloader.NewKeysFromStrings(keys))
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*ProductsByCategory
		for _, res := range resSlice {
			results = append(results, res.(*ProductsByCategory))
		}

		return results, nil
	}, nil
}

type GetProductsRequestKey struct {
	*GetProductsRequest
}

func (key *GetProductsRequestKey) String() string {
	return pg.ProtoKey(key)
}

func (key *GetProductsRequestKey) Raw() interface{} {
	return key
}

func LoadProducts(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadProductsLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadProductsLoader").(*dataloader.Loader)
	default:
		panic("Please call products.WithLoaders with the current context first")
	}

	thunk := loader.Load(p.Context, &GetProductsRequestKey{key})
	return func() (interface{}, error) {
		res, err := thunk()
		if err != nil {
			return nil, err
		}
		return res.(*GetProductsResponse), nil
	}, nil
}

func LoadProductsMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("LoadProductsLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("LoadProductsLoader").(*dataloader.Loader)
	default:
		panic("Please call products.WithLoaders with the current context first")
	}

	loaderKeys := make(dataloader.Keys, len(keys))
	for ix := range keys {
		loaderKeys[ix] = &GetProductsRequestKey{keys[ix]}
	}

	thunk := loader.LoadMany(p.Context, loaderKeys)
	return func() (interface{}, error) {
		resSlice, errSlice := thunk()

		for _, err := range errSlice {
			if err != nil {
				return nil, err
			}
		}

		var results []*GetProductsResponse
		for _, res := range resSlice {
			results = append(results, res.(*GetProductsResponse))
		}

		return results, nil
	}, nil
}

// allMessages contains all message types from this proto package
var allMessages = []pg.GraphqlMessage{
	&GetProductsRequest{},
	&GetProductsResponse{},
	&GetProductsByCategoryResponse{},
	&LoadProductsRequest{},
	&LoadProductsResponse{},
	&SearchProductsRequest{},
	&SearchProductsResponse{},
	&ProductsByCategory{},
	&Product{},
	&ProductVariant{},
	&Inventory{},
}

// ProductsModule implements the Module interface for the products package
type ProductsModule struct {
	productsClient  ProductsClient
	productsService ProductsServer

	dialOpts pg.DialOptions
}

// ProductsModuleOption configures the ProductsModule
type ProductsModuleOption func(*ProductsModule)

// WithModuleProductsClient sets the gRPC client for the Products service
func WithModuleProductsClient(client ProductsClient) ProductsModuleOption {
	return func(m *ProductsModule) {
		m.productsClient = client
	}
}

// WithModuleProductsService sets the direct service implementation for the Products service
func WithModuleProductsService(service ProductsServer) ProductsModuleOption {
	return func(m *ProductsModule) {
		m.productsService = service
	}
}

// WithDialOptions sets dial options for lazy client creation
func WithDialOptions(opts pg.DialOptions) ProductsModuleOption {
	return func(m *ProductsModule) {
		m.dialOpts = opts
	}
}

// NewProductsModule creates a new module with optional service configurations
func NewProductsModule(opts ...ProductsModuleOption) *ProductsModule {
	m := &ProductsModule{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// getProductsClient returns the client, creating it lazily if needed
func (m *ProductsModule) getProductsClient() ProductsClient {
	if m.productsClient == nil {
		m.productsClient = NewProductsClient(pg.GrpcConnection("localhost:50051", m.dialOpts["Products"]...))
	}
	return m.productsClient
}

// Fields returns all GraphQL query/mutation fields from all services in this module
func (m *ProductsModule) Fields() gql.Fields {
	fields := gql.Fields{}

	// Products service: GetProducts method
	fields["products_GetProducts"] = &gql.Field{
		Name: "products_GetProducts",
		Type: GetProductsResponseGraphqlType,
		Args: GetProductsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			req := GetProductsRequestFromArgs(p.Args)
			if m.productsService != nil {
				return m.productsService.GetProducts(p.Context, req)
			}
			return m.getProductsClient().GetProducts(p.Context, req)
		},
	}

	// Products service: SearchProducts method
	fields["products_SearchProducts"] = &gql.Field{
		Name: "products_SearchProducts",
		Type: SearchProductsResponseGraphqlType,
		Args: SearchProductsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			req := SearchProductsRequestFromArgs(p.Args)
			if m.productsService != nil {
				return m.productsService.SearchProducts(p.Context, req)
			}
			return m.getProductsClient().SearchProducts(p.Context, req)
		},
	}

	return fields
}

// WithLoaders registers all dataloaders from all services into the context
func (m *ProductsModule) WithLoaders(ctx context.Context) context.Context {
	// Products service: GetProductsByCategory loader
	ctx = context.WithValue(ctx, "GetProductsByCategoryLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *GetProductsByCategoryResponse
			var err error

			req := &pg.BatchRequest{
				Keys: keys.Keys(),
			}
			if m.productsService != nil {
				resp, err = m.productsService.GetProductsByCategory(ctx, req)
			} else {
				resp, err = m.getProductsClient().GetProductsByCategory(ctx, req)
			}

			if err != nil {
				// Return error result for each key - dataloader requires same number of results as keys
				for range keys {
					results = append(results, &dataloader.Result{Error: err})
				}
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *ProductsByCategory
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	// Products service: LoadProducts loader
	ctx = context.WithValue(ctx, "LoadProductsLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var requests []*GetProductsRequest
			for _, key := range keys {
				requests = append(requests, key.(*GetProductsRequestKey).GetProductsRequest)
			}
			var resp *LoadProductsResponse
			var err error

			req := &LoadProductsRequest{
				Reqs: requests,
			}
			if m.productsService != nil {
				resp, err = m.productsService.LoadProducts(ctx, req)
			} else {
				resp, err = m.getProductsClient().LoadProducts(ctx, req)
			}

			if err != nil {
				// Return error result for each key - dataloader requires same number of results as keys
				for range keys {
					results = append(results, &dataloader.Result{Error: err})
				}
				return results
			}

			for _, key := range keys.Keys() {
				if val, ok := resp.Results[key]; ok {
					results = append(results, &dataloader.Result{Data: val})
				} else {
					var empty *GetProductsResponse
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	return ctx
}

// Messages returns all message types from this package
func (m *ProductsModule) Messages() []pg.GraphqlMessage {
	return allMessages
}

// PackageName returns the proto package name
func (m *ProductsModule) PackageName() string {
	return "products"
}

// Type-safe field customization methods

// AddFieldToGetProductsRequest adds a custom field to the GetProductsRequest GraphQL type
func (m *ProductsModule) AddFieldToGetProductsRequest(fieldName string, field *gql.Field) {
	GetProductsRequestGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToGetProductsResponse adds a custom field to the GetProductsResponse GraphQL type
func (m *ProductsModule) AddFieldToGetProductsResponse(fieldName string, field *gql.Field) {
	GetProductsResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToGetProductsByCategoryResponse adds a custom field to the GetProductsByCategoryResponse GraphQL type
func (m *ProductsModule) AddFieldToGetProductsByCategoryResponse(fieldName string, field *gql.Field) {
	GetProductsByCategoryResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToLoadProductsRequest adds a custom field to the LoadProductsRequest GraphQL type
func (m *ProductsModule) AddFieldToLoadProductsRequest(fieldName string, field *gql.Field) {
	LoadProductsRequestGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToLoadProductsResponse adds a custom field to the LoadProductsResponse GraphQL type
func (m *ProductsModule) AddFieldToLoadProductsResponse(fieldName string, field *gql.Field) {
	LoadProductsResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToSearchProductsRequest adds a custom field to the SearchProductsRequest GraphQL type
func (m *ProductsModule) AddFieldToSearchProductsRequest(fieldName string, field *gql.Field) {
	SearchProductsRequestGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToSearchProductsResponse adds a custom field to the SearchProductsResponse GraphQL type
func (m *ProductsModule) AddFieldToSearchProductsResponse(fieldName string, field *gql.Field) {
	SearchProductsResponseGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToProductsByCategory adds a custom field to the ProductsByCategory GraphQL type
func (m *ProductsModule) AddFieldToProductsByCategory(fieldName string, field *gql.Field) {
	ProductsByCategoryGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToProduct adds a custom field to the Product GraphQL type
func (m *ProductsModule) AddFieldToProduct(fieldName string, field *gql.Field) {
	ProductGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToProductVariant adds a custom field to the ProductVariant GraphQL type
func (m *ProductsModule) AddFieldToProductVariant(fieldName string, field *gql.Field) {
	ProductVariantGraphqlType.AddFieldConfig(fieldName, field)
}

// AddFieldToInventory adds a custom field to the Inventory GraphQL type
func (m *ProductsModule) AddFieldToInventory(fieldName string, field *gql.Field) {
	InventoryGraphqlType.AddFieldConfig(fieldName, field)
}

// GraphQL type accessors

// GetProductsRequestType returns the GraphQL type for GetProductsRequest
func (m *ProductsModule) GetProductsRequestType() *gql.Object {
	return GetProductsRequestGraphqlType
}

// GetProductsResponseType returns the GraphQL type for GetProductsResponse
func (m *ProductsModule) GetProductsResponseType() *gql.Object {
	return GetProductsResponseGraphqlType
}

// GetProductsByCategoryResponseType returns the GraphQL type for GetProductsByCategoryResponse
func (m *ProductsModule) GetProductsByCategoryResponseType() *gql.Object {
	return GetProductsByCategoryResponseGraphqlType
}

// LoadProductsRequestType returns the GraphQL type for LoadProductsRequest
func (m *ProductsModule) LoadProductsRequestType() *gql.Object {
	return LoadProductsRequestGraphqlType
}

// LoadProductsResponseType returns the GraphQL type for LoadProductsResponse
func (m *ProductsModule) LoadProductsResponseType() *gql.Object {
	return LoadProductsResponseGraphqlType
}

// SearchProductsRequestType returns the GraphQL type for SearchProductsRequest
func (m *ProductsModule) SearchProductsRequestType() *gql.Object {
	return SearchProductsRequestGraphqlType
}

// SearchProductsResponseType returns the GraphQL type for SearchProductsResponse
func (m *ProductsModule) SearchProductsResponseType() *gql.Object {
	return SearchProductsResponseGraphqlType
}

// ProductsByCategoryType returns the GraphQL type for ProductsByCategory
func (m *ProductsModule) ProductsByCategoryType() *gql.Object {
	return ProductsByCategoryGraphqlType
}

// ProductType returns the GraphQL type for Product
func (m *ProductsModule) ProductType() *gql.Object {
	return ProductGraphqlType
}

// ProductVariantType returns the GraphQL type for ProductVariant
func (m *ProductsModule) ProductVariantType() *gql.Object {
	return ProductVariantGraphqlType
}

// InventoryType returns the GraphQL type for Inventory
func (m *ProductsModule) InventoryType() *gql.Object {
	return InventoryGraphqlType
}

// DataLoader accessor methods

// ProductsGetProductsByCategory loads a single *ProductsByCategory using the products service dataloader
func (m *ProductsModule) ProductsGetProductsByCategory(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return GetProductsByCategory(p, key)
}

// ProductsGetProductsByCategoryMany loads multiple *ProductsByCategory using the products service dataloader
func (m *ProductsModule) ProductsGetProductsByCategoryMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return GetProductsByCategoryMany(p, keys)
}

// ProductsLoadProducts loads a single *GetProductsResponse using the products service dataloader
func (m *ProductsModule) ProductsLoadProducts(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProducts(p, key)
}

// ProductsLoadProductsMany loads multiple *GetProductsResponse using the products service dataloader
func (m *ProductsModule) ProductsLoadProductsMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProductsMany(p, keys)
}

// Service instance accessors

// ProductsInstance is a unified interface for calling Products methods
// It works with both gRPC clients and direct service implementations
type ProductsInstance interface {
	GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error)
	SearchProducts(ctx context.Context, req *SearchProductsRequest) (*SearchProductsResponse, error)
	GetProductsByCategory(ctx context.Context, req *pg.BatchRequest) (*GetProductsByCategoryResponse, error)
	GetProductsByCategoryBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error)
	GetProductsByCategoryBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error)
	LoadProducts(ctx context.Context, req *LoadProductsRequest) (*LoadProductsResponse, error)
	LoadProductsBatch(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error)
	LoadProductsBatchMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error)
}

type productsServerAdapter struct {
	server ProductsServer
}

func (a *productsServerAdapter) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	return a.server.GetProducts(ctx, req)
}

func (a *productsServerAdapter) SearchProducts(ctx context.Context, req *SearchProductsRequest) (*SearchProductsResponse, error) {
	return a.server.SearchProducts(ctx, req)
}

func (a *productsServerAdapter) GetProductsByCategory(ctx context.Context, req *pg.BatchRequest) (*GetProductsByCategoryResponse, error) {
	return a.server.GetProductsByCategory(ctx, req)
}

func (a *productsServerAdapter) GetProductsByCategoryBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return GetProductsByCategory(p, key)
}

func (a *productsServerAdapter) GetProductsByCategoryBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return GetProductsByCategoryMany(p, keys)
}

func (a *productsServerAdapter) LoadProducts(ctx context.Context, req *LoadProductsRequest) (*LoadProductsResponse, error) {
	return a.server.LoadProducts(ctx, req)
}

func (a *productsServerAdapter) LoadProductsBatch(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProducts(p, key)
}

func (a *productsServerAdapter) LoadProductsBatchMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProductsMany(p, keys)
}

type productsClientAdapter struct {
	client ProductsClient
}

func (a *productsClientAdapter) GetProducts(ctx context.Context, req *GetProductsRequest) (*GetProductsResponse, error) {
	return a.client.GetProducts(ctx, req)
}

func (a *productsClientAdapter) SearchProducts(ctx context.Context, req *SearchProductsRequest) (*SearchProductsResponse, error) {
	return a.client.SearchProducts(ctx, req)
}

func (a *productsClientAdapter) GetProductsByCategory(ctx context.Context, req *pg.BatchRequest) (*GetProductsByCategoryResponse, error) {
	return a.client.GetProductsByCategory(ctx, req)
}

func (a *productsClientAdapter) GetProductsByCategoryBatch(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
	return GetProductsByCategory(p, key)
}

func (a *productsClientAdapter) GetProductsByCategoryBatchMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
	return GetProductsByCategoryMany(p, keys)
}

func (a *productsClientAdapter) LoadProducts(ctx context.Context, req *LoadProductsRequest) (*LoadProductsResponse, error) {
	return a.client.LoadProducts(ctx, req)
}

func (a *productsClientAdapter) LoadProductsBatch(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProducts(p, key)
}

func (a *productsClientAdapter) LoadProductsBatchMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error) {
	return LoadProductsMany(p, keys)
}

// Products returns a unified ProductsInstance that works with both clients and services
// Returns nil if neither client nor service is configured
func (m *ProductsModule) Products() ProductsInstance {
	if m.productsClient != nil {
		return &productsClientAdapter{client: m.productsClient}
	}
	if m.productsService != nil {
		return &productsServerAdapter{server: m.productsService}
	}
	return nil
}

// Backward compatibility layer for v0.51.7 API
// All functions below are deprecated and will be removed in a future version.
// Please migrate to the module-based API using NewProductsModule()

var defaultModule *ProductsModule

// ProductsClientInstance provides a unified Products client interface
// Deprecated: Use NewProductsModule().Products() instead
var ProductsClientInstance ProductsInstance

// SetDefaultModule allows you to set a custom module instance as the default
// for use with deprecated package-level functions.
// This allows you to configure a module once and have all deprecated functions use it.
// Example:
//
//	module := NewProductsModule(WithDialOptions(...))
//	SetDefaultModule(module)
//	// Now all deprecated Init(), WithLoaders(), etc. will use your module
func SetDefaultModule(module *ProductsModule) {
	defaultModule = module
}

func getDefaultModule() *ProductsModule {
	if defaultModule == nil {
		defaultModule = NewProductsModule()
	}
	return defaultModule
}

func init() {
	// Initialize ProductsClientInstance with lazy-loading adapter
	ProductsClientInstance = &productsClientAdapter{client: nil}
}

// ProductsInit initializes the Products service.
// Deprecated: Use NewProductsModule() and configure with WithModuleProductsClient() or WithModuleProductsService() instead.
func ProductsInit(ctx context.Context, opts ...ProductsModuleOption) (context.Context, []*gql.Field) {
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
	servicePrefix := "products_"
	for name, field := range fields {
		if strings.HasPrefix(name, servicePrefix) {
			serviceFields = append(serviceFields, field)
		}
	}

	return ctx, serviceFields
}

// ProductsWithLoaders registers dataloaders for the Products service into the context.
// Deprecated: Use NewProductsModule().WithLoaders(ctx) instead.
func ProductsWithLoaders(ctx context.Context) context.Context {
	return getDefaultModule().WithLoaders(ctx)
}

// WithLoaders registers all dataloaders from all services into the context.
// Deprecated: Use NewProductsModule().WithLoaders(ctx) instead.
func WithLoaders(ctx context.Context) context.Context {
	return getDefaultModule().WithLoaders(ctx)
}

// Fields returns all GraphQL query/mutation fields from all services.
// Deprecated: Use NewProductsModule().Fields() instead.
func Fields() gql.Fields {
	return getDefaultModule().Fields()
}
