package receipt

import (
	"math"

	"gitlab.com/sorleone/pepper/pkg/product"
)

//------------------------------------------------------------------------------

// Record represents a record of a Receipt
type Record struct {
	Name     string  `json:"name,omitempty"`
	Imported bool    `json:"imported,omitempty"`
	Quantity uint64  `json:"quantity,omitempty"`
	Total    float64 `json:"total,omitempty"`
}

// Total represents the total of a Receipt
type Total struct {
	Records []Record `json:"records,omitempty"`
	Total   float64  `json:"total,omitempty"`
	Taxes   float64  `json:"taxes,omitempty"`
}

// Receipt represents a receipt
type Receipt struct {
	total                 float64
	taxes                 float64
	basicSalesTax         float64
	importDuty            float64
	recordCollisionMap    map[string]*Record
	taxExemptProductTypes map[product.Type]bool
}

//------------------------------------------------------------------------------

// MakeReceipt makes a new Receipt with standard parameters
func MakeReceipt() *Receipt {
	return &Receipt{
		basicSalesTax:      0.1,
		importDuty:         0.05,
		recordCollisionMap: make(map[string]*Record),
		taxExemptProductTypes: map[product.Type]bool{
			product.BookProductType:    true,
			product.FoodProductType:    true,
			product.MedicalProductType: true,
		},
	}
}

// Add adds Product(s) to the Receipt and returns it
func (receipt *Receipt) Add(products ...product.Product) *Receipt {
	var tax float64

	for _, product := range products {
		taxExempt := receipt.taxExemptProductTypes[product.Type]

		if product.Imported {
			if taxExempt {
				tax = receipt.importDuty
			} else {
				tax = receipt.basicSalesTax + receipt.importDuty
			}
		} else if taxExempt {
			tax = 0
		} else {
			tax = receipt.basicSalesTax
		}

		hash := product.Hash()
		record, ok := receipt.recordCollisionMap[hash]
		taxesOnPrice := roundNearest(product.Price * tax)
		taxedPrice := roundCents(product.Price + taxesOnPrice)
		receipt.taxes += taxesOnPrice
		receipt.total += taxedPrice

		if ok {
			record.Quantity++
			record.Total += taxedPrice
		} else {
			receipt.recordCollisionMap[hash] = &Record{
				Name:     product.Name,
				Quantity: 1,
				Total:    taxedPrice,
				Imported: product.Imported,
			}
		}
	}

	return receipt
}

// GetTotal returns the total of a Receipt
func (receipt *Receipt) GetTotal() *Total {
	records := make([]Record, 0, len(receipt.recordCollisionMap))
	for _, record := range receipt.recordCollisionMap {
		records = append(records, *record)
	}

	return &Total{
		Records: records,
		Total:   roundCents(receipt.total),
		Taxes:   roundCents(receipt.taxes),
	}
}

//------------------------------------------------------------------------------

func roundCents(value float64) float64 {
	return math.Round(value*100) / 100
}

func roundNearest(value float64) float64 {
	return float64(int64(5*math.Ceil(value/0.05))) / 100
}
