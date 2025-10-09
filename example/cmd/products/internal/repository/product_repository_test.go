package repository

import (
	"testing"

	"github.com/kitt-technology/protoc-gen-graphql/example/products"
)

func TestInMemoryProductRepository_GetByID(t *testing.T) {
	repo := NewInMemoryProductRepository()

	tests := []struct {
		name      string
		id        string
		wantFound bool
		wantName  string
	}{
		{
			name:      "existing product",
			id:        "1",
			wantFound: true,
			wantName:  "Wireless Bluetooth Headphones",
		},
		{
			name:      "non-existing product",
			id:        "999",
			wantFound: false,
		},
		{
			name:      "empty id",
			id:        "",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, found := repo.GetByID(tt.id)
			if found != tt.wantFound {
				t.Errorf("GetByID() found = %v, want %v", found, tt.wantFound)
			}
			if found && product.Name != tt.wantName {
				t.Errorf("GetByID() name = %v, want %v", product.Name, tt.wantName)
			}
		})
	}
}

func TestInMemoryProductRepository_GetByIDs(t *testing.T) {
	repo := NewInMemoryProductRepository()

	tests := []struct {
		name      string
		ids       []string
		wantCount int
	}{
		{
			name:      "multiple existing products",
			ids:       []string{"1", "2", "3"},
			wantCount: 3,
		},
		{
			name:      "some existing, some not",
			ids:       []string{"1", "999", "2"},
			wantCount: 2,
		},
		{
			name:      "no existing products",
			ids:       []string{"999", "998"},
			wantCount: 0,
		},
		{
			name:      "empty ids",
			ids:       []string{},
			wantCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetByIDs(tt.ids)
			if len(result) != tt.wantCount {
				t.Errorf("GetByIDs() count = %d, want %d", len(result), tt.wantCount)
			}
		})
	}
}

func TestInMemoryProductRepository_GetByCategory(t *testing.T) {
	repo := NewInMemoryProductRepository()

	tests := []struct {
		name      string
		category  products.Category
		wantCount int
	}{
		{
			name:      "electronics category",
			category:  products.Category_ELECTRONICS,
			wantCount: 2, // Headphones and Smart Watch
		},
		{
			name:      "sports category",
			category:  products.Category_SPORTS,
			wantCount: 2, // Running Shoes and Yoga Mat
		},
		{
			name:      "home garden category",
			category:  products.Category_HOME_GARDEN,
			wantCount: 1, // Coffee Maker
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetByCategory(tt.category)
			if len(result) != tt.wantCount {
				t.Errorf("GetByCategory() count = %d, want %d", len(result), tt.wantCount)
			}
			// Verify all returned products are in the requested category
			for _, product := range result {
				if product.Category != tt.category {
					t.Errorf("GetByCategory() returned product with category %v, want %v", product.Category, tt.category)
				}
			}
		})
	}
}

func TestInMemoryProductRepository_GetByCategories(t *testing.T) {
	repo := NewInMemoryProductRepository()

	tests := []struct {
		name       string
		categories []products.Category
		wantCount  int
	}{
		{
			name:       "multiple categories",
			categories: []products.Category{products.Category_ELECTRONICS, products.Category_SPORTS},
			wantCount:  4, // 2 electronics + 2 sports
		},
		{
			name:       "single category",
			categories: []products.Category{products.Category_HOME_GARDEN},
			wantCount:  1,
		},
		{
			name:       "empty categories",
			categories: []products.Category{},
			wantCount:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.GetByCategories(tt.categories)
			if len(result) != tt.wantCount {
				t.Errorf("GetByCategories() count = %d, want %d", len(result), tt.wantCount)
			}
		})
	}
}

func TestInMemoryProductRepository_GetInStock(t *testing.T) {
	repo := NewInMemoryProductRepository()

	result := repo.GetInStock()

	// All seeded products should have stock
	if len(result) == 0 {
		t.Error("GetInStock() returned no products, expected at least some")
	}

	// Verify all returned products have inventory > 0
	for _, product := range result {
		if product.Inventory == nil || product.Inventory.Quantity <= 0 {
			t.Errorf("GetInStock() returned product %s with no stock", product.Id)
		}
	}
}

func TestInMemoryProductRepository_GetAll(t *testing.T) {
	repo := NewInMemoryProductRepository()

	result := repo.GetAll()

	// We seeded 5 products
	expectedCount := 5
	if len(result) != expectedCount {
		t.Errorf("GetAll() count = %d, want %d", len(result), expectedCount)
	}
}

func TestInMemoryProductRepository_Search(t *testing.T) {
	repo := NewInMemoryProductRepository()

	tests := []struct {
		name      string
		query     string
		limit     int
		wantCount int
	}{
		{
			name:      "limit less than total",
			query:     "test",
			limit:     3,
			wantCount: 3,
		},
		{
			name:      "limit greater than total",
			query:     "test",
			limit:     100,
			wantCount: 5, // All products
		},
		{
			name:      "limit of 1",
			query:     "test",
			limit:     1,
			wantCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.Search(tt.query, tt.limit)
			if len(result) != tt.wantCount {
				t.Errorf("Search() count = %d, want %d", len(result), tt.wantCount)
			}
		})
	}
}

func TestInMemoryProductRepository_Concurrency(t *testing.T) {
	repo := NewInMemoryProductRepository()

	// Test concurrent reads
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			_ = repo.GetAll()
			_, _ = repo.GetByID("1")
			_ = repo.GetByCategory(products.Category_ELECTRONICS)
			done <- true
		}()
	}

	// Wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}
}
