package repository

import (
	"sync"
	"time"

	"github.com/kitt-technology/protoc-gen-graphql/example/common-example"
	"github.com/kitt-technology/protoc-gen-graphql/example/products"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	GetByID(id string) (*products.Product, bool)
	GetByIDs(ids []string) []*products.Product
	GetByCategory(category products.Category) []*products.Product
	GetByCategories(categories []products.Category) []*products.Product
	GetInStock() []*products.Product
	GetAll() []*products.Product
	Search(query string, limit int) []*products.Product
}

// InMemoryProductRepository is an in-memory implementation of ProductRepository
type InMemoryProductRepository struct {
	mu       sync.RWMutex
	products map[string]*products.Product
}

// NewInMemoryProductRepository creates a new in-memory product repository with sample data
func NewInMemoryProductRepository() *InMemoryProductRepository {
	repo := &InMemoryProductRepository{
		products: make(map[string]*products.Product),
	}
	repo.seedData()
	return repo
}

// GetByID retrieves a product by its ID
func (r *InMemoryProductRepository) GetByID(id string) (*products.Product, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, ok := r.products[id]
	return product, ok
}

// GetByIDs retrieves products by their IDs
func (r *InMemoryProductRepository) GetByIDs(ids []string) []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*products.Product, 0, len(ids))
	for _, id := range ids {
		if product, ok := r.products[id]; ok {
			result = append(result, product)
		}
	}
	return result
}

// GetByCategory retrieves all products in a specific category
func (r *InMemoryProductRepository) GetByCategory(category products.Category) []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*products.Product
	for _, product := range r.products {
		if product.Category == category {
			result = append(result, product)
		}
	}
	return result
}

// GetByCategories retrieves all products in the specified categories
func (r *InMemoryProductRepository) GetByCategories(categories []products.Category) []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categorySet := make(map[products.Category]bool)
	for _, cat := range categories {
		categorySet[cat] = true
	}

	var result []*products.Product
	for _, product := range r.products {
		if categorySet[product.Category] {
			result = append(result, product)
		}
	}
	return result
}

// GetInStock retrieves all products that have inventory
func (r *InMemoryProductRepository) GetInStock() []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*products.Product
	for _, product := range r.products {
		if product.Inventory != nil && product.Inventory.Quantity > 0 {
			result = append(result, product)
		}
	}
	return result
}

// GetAll retrieves all products
func (r *InMemoryProductRepository) GetAll() []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*products.Product, 0, len(r.products))
	for _, product := range r.products {
		result = append(result, product)
	}
	return result
}

// Search performs a simple search on products (in a real app, this would be more sophisticated)
func (r *InMemoryProductRepository) Search(query string, limit int) []*products.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// For now, return all products (in a real app, we'd filter by query)
	result := make([]*products.Product, 0)
	for _, product := range r.products {
		result = append(result, product)
		if len(result) >= limit {
			break
		}
	}
	return result
}

// seedData initializes the repository with sample data
func (r *InMemoryProductRepository) seedData() {
	now := timestamppb.New(time.Now())
	lastMonth := timestamppb.New(time.Now().AddDate(0, -1, 0))

	r.products = map[string]*products.Product{
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
			Rating:      floatPtr(4.7),
			ReviewCount: int32Ptr(1523),
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
			Rating:      floatPtr(4.5),
			ReviewCount: int32Ptr(892),
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
			Rating:      floatPtr(4.8),
			ReviewCount: int32Ptr(2341),
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
			Rating:      floatPtr(4.3),
			ReviewCount: int32Ptr(567),
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
			Rating:      floatPtr(4.6),
			ReviewCount: int32Ptr(734),
		},
	}
}

// Helper functions for creating pointers
func floatPtr(f float32) *float32 {
	return &f
}

func int32Ptr(i int32) *int32 {
	return &i
}
