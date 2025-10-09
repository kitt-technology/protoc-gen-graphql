package service

import (
	"context"
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"github.com/kitt-technology/protoc-gen-graphql/graphql"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// mockProductRepository is a mock implementation of ProductRepository for testing
type mockProductRepository struct {
	products map[string]*products.Product
}

func newMockProductRepository() *mockProductRepository {
	return &mockProductRepository{
		products: map[string]*products.Product{
			"1": {
				Id:       "1",
				Name:     "Product 1",
				Category: products.Category_ELECTRONICS,
				Price: &common_example.Money{
					CurrencyCode: "USD",
					Units:        100,
				},
				Inventory: &products.Inventory{
					Quantity: 10,
				},
			},
			"2": {
				Id:       "2",
				Name:     "Product 2",
				Category: products.Category_SPORTS,
				Price: &common_example.Money{
					CurrencyCode: "USD",
					Units:        50,
				},
				Inventory: &products.Inventory{
					Quantity: 0, // Out of stock
				},
			},
			"3": {
				Id:       "3",
				Name:     "Product 3",
				Category: products.Category_ELECTRONICS,
				Price: &common_example.Money{
					CurrencyCode: "USD",
					Units:        200,
				},
				Inventory: &products.Inventory{
					Quantity: 5,
				},
			},
		},
	}
}

func (m *mockProductRepository) GetByID(id string) (*products.Product, bool) {
	product, ok := m.products[id]
	return product, ok
}

func (m *mockProductRepository) GetByIDs(ids []string) []*products.Product {
	result := make([]*products.Product, 0, len(ids))
	for _, id := range ids {
		if product, ok := m.products[id]; ok {
			result = append(result, product)
		}
	}
	return result
}

func (m *mockProductRepository) GetByCategory(category products.Category) []*products.Product {
	var result []*products.Product
	for _, product := range m.products {
		if product.Category == category {
			result = append(result, product)
		}
	}
	return result
}

func (m *mockProductRepository) GetByCategories(categories []products.Category) []*products.Product {
	categorySet := make(map[products.Category]bool)
	for _, cat := range categories {
		categorySet[cat] = true
	}

	var result []*products.Product
	for _, product := range m.products {
		if categorySet[product.Category] {
			result = append(result, product)
		}
	}
	return result
}

func (m *mockProductRepository) GetInStock() []*products.Product {
	var result []*products.Product
	for _, product := range m.products {
		if product.Inventory != nil && product.Inventory.Quantity > 0 {
			result = append(result, product)
		}
	}
	return result
}

func (m *mockProductRepository) GetAll() []*products.Product {
	result := make([]*products.Product, 0, len(m.products))
	for _, product := range m.products {
		result = append(result, product)
	}
	return result
}

func (m *mockProductRepository) Search(query string, limit int) []*products.Product {
	result := make([]*products.Product, 0)
	for _, product := range m.products {
		result = append(result, product)
		if len(result) >= limit {
			break
		}
	}
	return result
}

func TestProductService_GetProducts(t *testing.T) {
	repo := newMockProductRepository()
	service := NewProductService(repo)
	ctx := context.Background()

	tests := []struct {
		name       string
		request    *products.GetProductsRequest
		wantCount  int
		wantErr    bool
	}{
		{
			name:      "get all products",
			request:   &products.GetProductsRequest{},
			wantCount: 3,
			wantErr:   false,
		},
		{
			name: "get by IDs",
			request: &products.GetProductsRequest{
				Ids: []string{"1", "2"},
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "get by single ID",
			request: &products.GetProductsRequest{
				Ids: []string{"1"},
			},
			wantCount: 1,
			wantErr:   false,
		},
		{
			name: "get by non-existent ID",
			request: &products.GetProductsRequest{
				Ids: []string{"999"},
			},
			wantCount: 0,
			wantErr:   false,
		},
		{
			name: "get by categories",
			request: &products.GetProductsRequest{
				Categories: []products.Category{products.Category_ELECTRONICS},
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "get by multiple categories",
			request: &products.GetProductsRequest{
				Categories: []products.Category{products.Category_ELECTRONICS, products.Category_SPORTS},
			},
			wantCount: 3,
			wantErr:   false,
		},
		{
			name: "get in stock only",
			request: &products.GetProductsRequest{
				InStockOnly: wrapperspb.Bool(true),
			},
			wantCount: 2, // Product 2 is out of stock
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetProducts(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Products) != tt.wantCount {
				t.Errorf("GetProducts() count = %d, want %d", len(resp.Products), tt.wantCount)
			}
			if !tt.wantErr && resp.PageInfo == nil {
				t.Error("GetProducts() PageInfo is nil")
			}
			if !tt.wantErr && resp.PageInfo != nil && resp.PageInfo.TotalCount != int32(tt.wantCount) {
				t.Errorf("GetProducts() PageInfo.TotalCount = %d, want %d", resp.PageInfo.TotalCount, tt.wantCount)
			}
		})
	}
}

func TestProductService_SearchProducts(t *testing.T) {
	repo := newMockProductRepository()
	service := NewProductService(repo)
	ctx := context.Background()

	tests := []struct {
		name      string
		request   *products.SearchProductsRequest
		wantCount int
		wantErr   bool
	}{
		{
			name: "search with default limit",
			request: &products.SearchProductsRequest{
				Query: "test",
			},
			wantCount: 3, // All products (less than default limit of 10)
			wantErr:   false,
		},
		{
			name: "search with custom limit",
			request: &products.SearchProductsRequest{
				Query: "test",
				Limit: 2,
			},
			wantCount: 2,
			wantErr:   false,
		},
		{
			name: "search with limit greater than total",
			request: &products.SearchProductsRequest{
				Query: "test",
				Limit: 100,
			},
			wantCount: 3,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.SearchProducts(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Products) != tt.wantCount {
				t.Errorf("SearchProducts() count = %d, want %d", len(resp.Products), tt.wantCount)
			}
			if !tt.wantErr && resp.PageInfo == nil {
				t.Error("SearchProducts() PageInfo is nil")
			}
		})
	}
}

func TestProductService_GetProductsByCategory(t *testing.T) {
	repo := newMockProductRepository()
	service := NewProductService(repo)
	ctx := context.Background()

	tests := []struct {
		name          string
		request       *graphql.BatchRequest
		wantResultLen int
		wantErr       bool
	}{
		{
			name: "single category",
			request: &graphql.BatchRequest{
				Keys: []string{"ELECTRONICS"},
			},
			wantResultLen: 1,
			wantErr:       false,
		},
		{
			name: "multiple categories",
			request: &graphql.BatchRequest{
				Keys: []string{"ELECTRONICS", "SPORTS"},
			},
			wantResultLen: 2,
			wantErr:       false,
		},
		{
			name: "invalid category",
			request: &graphql.BatchRequest{
				Keys: []string{"INVALID_CATEGORY"},
			},
			wantResultLen: 0,
			wantErr:       false,
		},
		{
			name: "mixed valid and invalid",
			request: &graphql.BatchRequest{
				Keys: []string{"ELECTRONICS", "INVALID"},
			},
			wantResultLen: 1,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetProductsByCategory(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductsByCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Results) != tt.wantResultLen {
				t.Errorf("GetProductsByCategory() result count = %d, want %d", len(resp.Results), tt.wantResultLen)
			}
		})
	}
}

func TestProductService_GetProductsBatch(t *testing.T) {
	repo := newMockProductRepository()
	service := NewProductService(repo)
	ctx := context.Background()

	tests := []struct {
		name          string
		request       *products.GetProductsBatchRequest
		wantResultLen int
		wantErr       bool
	}{
		{
			name: "single batch request",
			request: &products.GetProductsBatchRequest{
				Reqs: []*products.GetProductsRequest{
					{Ids: []string{"1"}},
				},
			},
			wantResultLen: 1,
			wantErr:       false,
		},
		{
			name: "multiple batch requests",
			request: &products.GetProductsBatchRequest{
				Reqs: []*products.GetProductsRequest{
					{Ids: []string{"1"}},
					{Ids: []string{"2"}},
					{Categories: []products.Category{products.Category_ELECTRONICS}},
				},
			},
			wantResultLen: 3,
			wantErr:       false,
		},
		{
			name: "empty batch requests",
			request: &products.GetProductsBatchRequest{
				Reqs: []*products.GetProductsRequest{},
			},
			wantResultLen: 0,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetProductsBatch(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductsBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(resp.Results) != tt.wantResultLen {
				t.Errorf("GetProductsBatch() result count = %d, want %d", len(resp.Results), tt.wantResultLen)
			}
		})
	}
}
