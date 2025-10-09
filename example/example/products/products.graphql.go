package products

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

var GetProductsBatchRequestGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetProductsBatchRequest",
	Fields: gql.Fields{
		"reqs": &gql.Field{
			Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlType)),
		},
	},
})

var GetProductsBatchRequestGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetProductsBatchRequestInput",
	Fields: gql.InputObjectConfigFieldMap{
		"reqs": &gql.InputObjectFieldConfig{
			Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlInputType)),
		},
	},
})

var GetProductsBatchRequestGraphqlArgs = gql.FieldConfigArgument{
	"reqs": &gql.ArgumentConfig{
		Type: gql.NewList(gql.NewNonNull(GetProductsRequestGraphqlInputType)),
	},
}

func GetProductsBatchRequestFromArgs(args map[string]interface{}) *GetProductsBatchRequest {
	return GetProductsBatchRequestInstanceFromArgs(&GetProductsBatchRequest{}, args)
}

func GetProductsBatchRequestInstanceFromArgs(objectFromArgs *GetProductsBatchRequest, args map[string]interface{}) *GetProductsBatchRequest {
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

func (objectFromArgs *GetProductsBatchRequest) FromArgs(args map[string]interface{}) {
	GetProductsBatchRequestInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetProductsBatchRequest) XXX_GraphqlType() *gql.Object {
	return GetProductsBatchRequestGraphqlType
}

func (msg *GetProductsBatchRequest) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetProductsBatchRequestGraphqlArgs
}

func (msg *GetProductsBatchRequest) XXX_Package() string {
	return "products"
}

var GetProductsBatchResponseGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "GetProductsBatchResponse",
	Fields: gql.Fields{
		"_null": &gql.Field{
			Type: gql.Boolean,
		},
	},
})

var GetProductsBatchResponseGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "GetProductsBatchResponseInput",
	Fields: gql.InputObjectConfigFieldMap{
		"_null": &gql.InputObjectFieldConfig{
			Type: gql.Boolean,
		},
	},
})

var GetProductsBatchResponseGraphqlArgs = gql.FieldConfigArgument{
	"_null": &gql.ArgumentConfig{
		Type: gql.Boolean,
	},
}

func GetProductsBatchResponseFromArgs(args map[string]interface{}) *GetProductsBatchResponse {
	return GetProductsBatchResponseInstanceFromArgs(&GetProductsBatchResponse{}, args)
}

func GetProductsBatchResponseInstanceFromArgs(objectFromArgs *GetProductsBatchResponse, args map[string]interface{}) *GetProductsBatchResponse {
	return objectFromArgs
}

func (objectFromArgs *GetProductsBatchResponse) FromArgs(args map[string]interface{}) {
	GetProductsBatchResponseInstanceFromArgs(objectFromArgs, args)
}

func (msg *GetProductsBatchResponse) XXX_GraphqlType() *gql.Object {
	return GetProductsBatchResponseGraphqlType
}

func (msg *GetProductsBatchResponse) XXX_GraphqlArgs() gql.FieldConfigArgument {
	return GetProductsBatchResponseGraphqlArgs
}

func (msg *GetProductsBatchResponse) XXX_Package() string {
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

var ProductsClientInstance ProductsClient
var ProductsServiceInstance ProductsServer
var ProductsDialOpts []grpc.DialOption

type ProductsOption func(*ProductsConfig)

type ProductsConfig struct {
	service  ProductsServer
	client   ProductsClient
	dialOpts []grpc.DialOption
}

// WithProductsService sets the service implementation for direct calls (no gRPC)
func WithProductsService(service ProductsServer) ProductsOption {
	return func(cfg *ProductsConfig) {
		cfg.service = service
	}
}

// WithProductsClient sets the gRPC client for remote calls
func WithProductsClient(client ProductsClient) ProductsOption {
	return func(cfg *ProductsConfig) {
		cfg.client = client
	}
}

// WithProductsDialOptions sets the dial options for the gRPC client
func WithProductsDialOptions(opts ...grpc.DialOption) ProductsOption {
	return func(cfg *ProductsConfig) {
		cfg.dialOpts = opts
	}
}

func ProductsInit(ctx context.Context, opts ...ProductsOption) (context.Context, []*gql.Field) {
	cfg := &ProductsConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	ProductsServiceInstance = cfg.service
	ProductsClientInstance = cfg.client
	ProductsDialOpts = cfg.dialOpts

	var fields []*gql.Field
	fields = append(fields, &gql.Field{
		Name: "products_GetProducts",
		Type: GetProductsResponseGraphqlType,
		Args: GetProductsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if ProductsServiceInstance != nil {
				return ProductsServiceInstance.GetProducts(p.Context, GetProductsRequestFromArgs(p.Args))
			}
			if ProductsClientInstance == nil {
				ProductsClientInstance = getProductsClient()
			}
			return ProductsClientInstance.GetProducts(p.Context, GetProductsRequestFromArgs(p.Args))
		},
	})

	fields = append(fields, &gql.Field{
		Name: "products_SearchProducts",
		Type: SearchProductsResponseGraphqlType,
		Args: SearchProductsRequestGraphqlArgs,
		Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if ProductsServiceInstance != nil {
				return ProductsServiceInstance.SearchProducts(p.Context, SearchProductsRequestFromArgs(p.Args))
			}
			if ProductsClientInstance == nil {
				ProductsClientInstance = getProductsClient()
			}
			return ProductsClientInstance.SearchProducts(p.Context, SearchProductsRequestFromArgs(p.Args))
		},
	})

	ctx = ProductsWithLoaders(ctx)

	return ctx, fields
}

func getProductsClient() ProductsClient {
	host := "localhost:50051"
	envHost := os.Getenv("SERVICE_HOST")
	if envHost != "" {
		host = envHost
	}
	return NewProductsClient(pg.GrpcConnection(host, ProductsDialOpts...))
}

// SetProductsService sets the service implementation for direct calls (no gRPC)
func SetProductsService(service ProductsServer) {
	ProductsServiceInstance = service
}

// SetProductsClient sets the gRPC client for remote calls
func SetProductsClient(client ProductsClient) {
	ProductsClientInstance = client
}

func ProductsWithLoaders(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "GetProductsByCategoryLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var resp *GetProductsByCategoryResponse
			var err error
			if ProductsServiceInstance != nil {
				resp, err = ProductsServiceInstance.GetProductsByCategory(ctx, &pg.BatchRequest{
					Keys: keys.Keys(),
				})
			} else {
				if ProductsClientInstance == nil {
					ProductsClientInstance = getProductsClient()
				}
				resp, err = ProductsClientInstance.GetProductsByCategory(ctx, &pg.BatchRequest{
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
					var empty *ProductsByCategory
					results = append(results, &dataloader.Result{Data: empty})
				}
			}

			return results
		},
	))

	ctx = context.WithValue(ctx, "GetProductsBatchLoader", dataloader.NewBatchedLoader(
		func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
			var results []*dataloader.Result

			var requests []*GetProductsRequest
			for _, key := range keys {
				requests = append(requests, key.(*GetProductsRequestKey).GetProductsRequest)
			}
			var resp *GetProductsBatchResponse
			var err error
			if ProductsServiceInstance != nil {
				resp, err = ProductsServiceInstance.GetProductsBatch(ctx, &GetProductsBatchRequest{
					Reqs: requests,
				})
			} else {
				if ProductsClientInstance == nil {
					ProductsClientInstance = getProductsClient()
				}
				resp, err = ProductsClientInstance.GetProductsBatch(ctx, &GetProductsBatchRequest{
					Reqs: requests,
				})
			}

			if err != nil {
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

func ProductsGetProductsByCategory(p gql.ResolveParams, key string) (func() (interface{}, error), error) {
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

func ProductsGetProductsByCategoryMany(p gql.ResolveParams, keys []string) (func() (interface{}, error), error) {
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

func ProductsGetProductsBatch(p gql.ResolveParams, key *GetProductsRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetProductsBatchLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetProductsBatchLoader").(*dataloader.Loader)
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

func ProductsGetProductsBatchMany(p gql.ResolveParams, keys []*GetProductsRequest) (func() (interface{}, error), error) {
	var loader *dataloader.Loader
	switch p.Context.Value("GetProductsBatchLoader").(type) {
	case *dataloader.Loader:
		loader = p.Context.Value("GetProductsBatchLoader").(*dataloader.Loader)
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
