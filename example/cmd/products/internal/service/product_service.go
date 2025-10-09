package service

import (
	"context"

	"github.com/kitt-technology/protoc-gen-graphql/example/cmd/products/internal/repository"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
)

// ProductService implements the Products gRPC service
type ProductService struct {
	products.UnimplementedProductsServer
	repo repository.ProductRepository
}

// NewProductService creates a new ProductService with the given repository
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// GetProductsByCategory implements batch loading by category
func (s *ProductService) GetProductsByCategory(ctx context.Context, request *graphql.BatchRequest) (*products.GetProductsByCategoryResponse, error) {
	results := make(map[string]*products.ProductsByCategory)

	for _, categoryStr := range request.Keys {
		// Parse category string to enum
		categoryValue, ok := products.Category_value[categoryStr]
		if !ok {
			// Invalid category, skip
			continue
		}
		category := products.Category(categoryValue)

		productList := s.repo.GetByCategory(category)
		results[categoryStr] = &products.ProductsByCategory{
			Results: productList,
		}
	}

	return &products.GetProductsByCategoryResponse{Results: results}, nil
}

// GetProductsBatch implements batch loading for complex requests
func (s *ProductService) GetProductsBatch(ctx context.Context, request *products.GetProductsBatchRequest) (*products.GetProductsBatchResponse, error) {
	results := make(map[string]*products.GetProductsResponse)

	for _, req := range request.Reqs {
		key := graphql.ProtoKey(req)
		resp, err := s.GetProducts(ctx, req)
		if err == nil {
			results[key] = resp
		}
	}

	return &products.GetProductsBatchResponse{Results: results}, nil
}

// GetProducts retrieves products based on the request criteria
func (s *ProductService) GetProducts(ctx context.Context, request *products.GetProductsRequest) (*products.GetProductsResponse, error) {
	var prods []*products.Product

	// Priority: InStockOnly > Ids > Categories > All
	switch {
	case request.InStockOnly != nil && request.InStockOnly.GetValue():
		prods = s.repo.GetInStock()
	case len(request.Ids) > 0:
		prods = s.repo.GetByIDs(request.Ids)
	case len(request.Categories) > 0:
		prods = s.repo.GetByCategories(request.Categories)
	default:
		prods = s.repo.GetAll()
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

// SearchProducts performs a search query on products
func (s *ProductService) SearchProducts(ctx context.Context, request *products.SearchProductsRequest) (*products.SearchProductsResponse, error) {
	limit := int(request.Limit)
	if limit <= 0 {
		limit = 10 // Default limit
	}

	prods := s.repo.Search(request.Query, limit)
	allProducts := s.repo.GetAll()

	return &products.SearchProductsResponse{
		Products: prods,
		PageInfo: &graphql.PageInfo{
			TotalCount:  int32(len(prods)),
			EndCursor:   "search_cursor",
			HasNextPage: len(allProducts) > limit,
		},
	}, nil
}
