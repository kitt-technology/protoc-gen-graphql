package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	products.RegisterProductsServer(s, &ProductService{})
	reflection.Register(s)

	fmt.Println("========================================")
	fmt.Println("Products gRPC Service")
	fmt.Println("========================================")
	fmt.Println("Listening on: localhost:50051")
	fmt.Println("\nThis service provides product catalog data")
	fmt.Println("Available sample products:")
	fmt.Println("  - Wireless Bluetooth Headphones ($79.99)")
	fmt.Println("  - Running Shoes - Trail Edition ($129.99)")
	fmt.Println("  - Smart Watch Pro ($299.99)")
	fmt.Println("  - Premium Coffee Maker ($89.99)")
	fmt.Println("  - Eco-Friendly Yoga Mat ($29.99)")
	fmt.Println("\nReady to accept gRPC requests...")
	fmt.Println("========================================\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type ProductService struct {
	products.UnimplementedProductsServer
}

func (s ProductService) GetProductsByCategory(ctx context.Context, request *graphql.BatchRequest) (*products.GetProductsByCategoryResponse, error) {
	var results = make(map[string]*products.ProductsByCategory)

	for _, category := range request.Keys {
		productsByCategory := products.ProductsByCategory{}
		for _, product := range productsDb {
			if product.Category.String() == category {
				productsByCategory.Results = append(productsByCategory.Results, product)
			}
		}
		results[category] = &productsByCategory
	}

	return &products.GetProductsByCategoryResponse{Results: results}, nil
}

func (s ProductService) GetProductsBatch(ctx context.Context, request *products.GetProductsBatchRequest) (*products.GetProductsBatchResponse, error) {
	var res products.GetProductsBatchResponse
	res.Results = make(map[string]*products.GetProductsResponse)
	for _, req := range request.Reqs {
		key := graphql.ProtoKey(req)
		resp, err := s.GetProducts(ctx, req)
		if err == nil {
			res.Results[key] = resp
		}
	}
	return &res, nil
}

func (s ProductService) GetProducts(ctx context.Context, request *products.GetProductsRequest) (*products.GetProductsResponse, error) {
	var prods []*products.Product

	if request.InStockOnly != nil && request.InStockOnly.GetValue() {
		for _, product := range productsDb {
			if product.Inventory.Quantity > 0 {
				prods = append(prods, product)
			}
		}
	} else if len(request.Ids) > 0 {
		for _, id := range request.Ids {
			if p, ok := productsDb[id]; ok {
				prods = append(prods, p)
			}
		}
	} else if len(request.Categories) > 0 {
		for _, product := range productsDb {
			for _, cat := range request.Categories {
				if product.Category == cat {
					prods = append(prods, product)
					break
				}
			}
		}
	} else {
		for _, product := range productsDb {
			prods = append(prods, product)
		}
	}

	return &products.GetProductsResponse{
		Products: prods,
		PageInfo: &graphql.PageInfo{
			TotalCount:  int32(len(prods)),
			EndCursor:   "cursor_end",
			HasNextPage: false,
		},
	}, nil
}

func (s ProductService) SearchProducts(ctx context.Context, request *products.SearchProductsRequest) (*products.SearchProductsResponse, error) {
	var prods []*products.Product

	// Simple search implementation
	for _, product := range productsDb {
		prods = append(prods, product)
	}

	limit := int32(10)
	if request.Limit > 0 {
		limit = request.Limit
	}

	if int32(len(prods)) > limit {
		prods = prods[:limit]
	}

	return &products.SearchProductsResponse{
		Products: prods,
		PageInfo: &graphql.PageInfo{
			TotalCount:  int32(len(prods)),
			EndCursor:   "search_cursor",
			HasNextPage: int32(len(productsDb)) > limit,
		},
	}, nil
}

var productsDb map[string]*products.Product

func init() {
	now := timestamppb.New(time.Now())
	lastMonth := timestamppb.New(time.Now().AddDate(0, -1, 0))

	productsDb = map[string]*products.Product{
		"1": {
			Id:          "1",
			Name:        "Wireless Bluetooth Headphones",
			Description: "High-quality over-ear headphones with active noise cancellation and 30-hour battery life",
			Category:    products.Category_ELECTRONICS,
			Price: &common_example.Money{
				CurrencyCode: "USD",
				Units:        79,
				
			},
			Inventory: &products.Inventory{
				Quantity:          150,
				Reserved:          10,
				WarehouseLocation: "A-12",
				LastRestocked:     lastMonth,
			},
			Variants: []*products.ProductVariant{
				{
					Id:   "1a",
					Name: "Black",
					Sku:  "WBH-BLK-001",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        79,
						
					},
					StockQuantity: 80,
					Attributes: map[string]string{
						"color": "black",
					},
				},
				{
					Id:   "1b",
					Name: "Silver",
					Sku:  "WBH-SLV-001",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        79,
						
					},
					StockQuantity: 70,
					Attributes: map[string]string{
						"color": "silver",
					},
				},
			},
			CreatedAt:   lastMonth,
			UpdatedAt:   now,
			Featured:    wrapperspb.Bool(true),
			SellerId:    "seller1",
			ImageUrls:   []string{"https://example.com/headphones-1.jpg", "https://example.com/headphones-2.jpg"},
			Rating:      proto.Float32(4.7),
			ReviewCount: proto.Int32(1523),
		},
		"2": {
			Id:          "2",
			Name:        "Running Shoes - Trail Edition",
			Description: "Professional trail running shoes with enhanced grip and waterproof membrane",
			Category:    products.Category_SPORTS,
			Price: &common_example.Money{
				CurrencyCode: "USD",
				Units:        129,
				
			},
			Inventory: &products.Inventory{
				Quantity:          200,
				Reserved:          25,
				WarehouseLocation: "B-05",
				LastRestocked:     now,
			},
			Variants: []*products.ProductVariant{
				{
					Id:   "2a",
					Name: "Size 9 - Blue",
					Sku:  "RUN-TR-09-BLU",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        129,
						
					},
					StockQuantity: 50,
					Attributes: map[string]string{
						"size":  "9",
						"color": "blue",
					},
				},
				{
					Id:   "2b",
					Name: "Size 10 - Blue",
					Sku:  "RUN-TR-10-BLU",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        129,
						
					},
					StockQuantity: 75,
					Attributes: map[string]string{
						"size":  "10",
						"color": "blue",
					},
				},
				{
					Id:   "2c",
					Name: "Size 10 - Black",
					Sku:  "RUN-TR-10-BLK",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        129,
						
					},
					StockQuantity: 75,
					Attributes: map[string]string{
						"size":  "10",
						"color": "black",
					},
				},
			},
			CreatedAt:   lastMonth,
			UpdatedAt:   now,
			Featured:    wrapperspb.Bool(false),
			SellerId:    "seller2",
			ImageUrls:   []string{"https://example.com/shoes-1.jpg"},
			Rating:      proto.Float32(4.5),
			ReviewCount: proto.Int32(892),
		},
		"3": {
			Id:          "3",
			Name:        "Smart Watch Pro",
			Description: "Advanced fitness tracker with heart rate monitor, GPS, and 7-day battery life",
			Category:    products.Category_ELECTRONICS,
			Price: &common_example.Money{
				CurrencyCode: "USD",
				Units:        299,
				
			},
			Inventory: &products.Inventory{
				Quantity:          75,
				Reserved:          15,
				WarehouseLocation: "A-08",
				LastRestocked:     lastMonth,
			},
			CreatedAt:   lastMonth,
			UpdatedAt:   now,
			Featured:    wrapperspb.Bool(true),
			SellerId:    "seller1",
			ImageUrls:   []string{"https://example.com/watch-1.jpg", "https://example.com/watch-2.jpg"},
			Rating:      proto.Float32(4.8),
			ReviewCount: proto.Int32(2341),
		},
		"4": {
			Id:          "4",
			Name:        "Premium Coffee Maker",
			Description: "Programmable drip coffee maker with thermal carafe and brew strength control",
			Category:    products.Category_HOME_GARDEN,
			Price: &common_example.Money{
				CurrencyCode: "USD",
				Units:        89,
				
			},
			Inventory: &products.Inventory{
				Quantity:          120,
				Reserved:          5,
				WarehouseLocation: "C-15",
				LastRestocked:     now,
			},
			CreatedAt:   lastMonth,
			UpdatedAt:   now,
			Featured:    wrapperspb.Bool(false),
			SellerId:    "seller3",
			ImageUrls:   []string{"https://example.com/coffee-1.jpg"},
			Rating:      proto.Float32(4.3),
			ReviewCount: proto.Int32(567),
		},
		"5": {
			Id:          "5",
			Name:        "Eco-Friendly Yoga Mat",
			Description: "Non-slip yoga mat made from sustainable materials with alignment guides",
			Category:    products.Category_SPORTS,
			Price: &common_example.Money{
				CurrencyCode: "USD",
				Units:        29,
				
			},
			Inventory: &products.Inventory{
				Quantity:          300,
				Reserved:          20,
				WarehouseLocation: "B-22",
				LastRestocked:     now,
			},
			Variants: []*products.ProductVariant{
				{
					Id:   "5a",
					Name: "Purple",
					Sku:  "YOGA-PRP-001",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        29,
						
					},
					StockQuantity: 150,
					Attributes: map[string]string{
						"color": "purple",
					},
				},
				{
					Id:   "5b",
					Name: "Green",
					Sku:  "YOGA-GRN-001",
					Price: &common_example.Money{
						CurrencyCode: "USD",
						Units:        29,
						
					},
					StockQuantity: 150,
					Attributes: map[string]string{
						"color": "green",
					},
				},
			},
			CreatedAt:   lastMonth,
			UpdatedAt:   now,
			Featured:    wrapperspb.Bool(false),
			SellerId:    "seller2",
			ImageUrls:   []string{"https://example.com/yoga-1.jpg"},
			Rating:      proto.Float32(4.6),
			ReviewCount: proto.Int32(734),
		},
	}
}

// proto is a helper package for creating proto optional values
var proto = struct {
	Float32 func(float32) *float32
	Int32   func(int32) *int32
}{
	Float32: func(f float32) *float32 { return &f },
	Int32:   func(i int32) *int32 { return &i },
}