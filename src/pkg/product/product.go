package product

import "fmt"

//------------------------------------------------------------------------------

// Type represents the type of a Product
type Type string

const (
	// MedicalProductType represents a medical Product
	MedicalProductType Type = "medical"

	// BookProductType represents a book Product
	BookProductType Type = "book"

	// FoodProductType represents a food Product
	FoodProductType Type = "food"

	// OtherProductType represents a generic Product
	OtherProductType Type = "other"
)

// Product represents a product
type Product struct {
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Type     Type    `json:"type,omitempty"`
	Imported bool    `json:"imported,omitempty"`
}

//------------------------------------------------------------------------------

// Hash returns the hash of a Product
// by concatenating the properties from first to last
func (product *Product) Hash() string {
	return fmt.Sprintf(
		"%v%v%v%v",
		product.Name,
		product.Price,
		product.Type,
		product.Imported,
	)
}
