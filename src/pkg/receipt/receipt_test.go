package receipt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/sorleone/pepper/pkg/product"
)

func TestBasicSalesTax(t *testing.T) {
	receipt := NewStandardReceipt().Add(
		product.Product{
			Name:  "book",
			Price: 12.49,
			Type:  product.BookProductType,
		},
		product.Product{
			Name:  "book",
			Price: 12.49,
			Type:  product.BookProductType,
		},
		product.Product{
			Name:  "music CD",
			Price: 14.99,
			Type:  product.OtherProductType,
		},
		product.Product{
			Name:  "chocolate bar",
			Price: 0.85,
			Type:  product.FoodProductType,
		},
	).GetTotal()

	assert.EqualValues(t, 42.32, receipt.Total)
	assert.EqualValues(t, 1.5, receipt.Taxes)

	for _, record := range receipt.Records {
		switch record.Name {
		case "book":
			assert.EqualValues(t, 2, record.Quantity)
			assert.EqualValues(t, 24.98, record.Total)
			break
		case "music CD":
			assert.EqualValues(t, 1, record.Quantity)
			assert.EqualValues(t, 16.49, record.Total)
			break
		case "chocolate bar":
			assert.EqualValues(t, 1, record.Quantity)
			assert.EqualValues(t, 0.85, record.Total)
			break
		}
	}
}

func TestImportedTax(t *testing.T) {
	receipt := NewStandardReceipt().Add(
		product.Product{
			Name:     "bottle of perfume",
			Price:    47.50,
			Type:     product.OtherProductType,
			Imported: true,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    10.00,
			Type:     product.FoodProductType,
			Imported: true,
		},
	).GetTotal()

	assert.EqualValues(t, 65.15, receipt.Total)
	assert.EqualValues(t, 7.65, receipt.Taxes)

	for _, record := range receipt.Records {
		switch record.Name {
		case "bottle of perfume":
			assert.EqualValues(t, 1, record.Quantity)
			assert.EqualValues(t, 54.65, record.Total)
			break
		case "box of chocolates":
			assert.EqualValues(t, 1, record.Quantity)
			assert.EqualValues(t, 10.50, record.Total)
			break
		}
	}
}

func TestConsecutiveAdd(t *testing.T) {
	repeat := 2.0
	receipt := NewStandardReceipt().Add(
		product.Product{
			Name:     "bottle of perfume",
			Price:    47.50,
			Type:     product.OtherProductType,
			Imported: true,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    10.00,
			Type:     product.FoodProductType,
			Imported: true,
		},
	).Add(
		product.Product{
			Name:     "bottle of perfume",
			Price:    47.50,
			Type:     product.OtherProductType,
			Imported: true,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    10.00,
			Type:     product.FoodProductType,
			Imported: true,
		},
	).GetTotal()

	assert.EqualValues(t, repeat*65.15, receipt.Total)
	assert.EqualValues(t, repeat*7.65, receipt.Taxes)

	for _, record := range receipt.Records {
		switch record.Name {
		case "bottle of perfume":
			assert.EqualValues(t, repeat*1, record.Quantity)
			assert.EqualValues(t, repeat*54.65, record.Total)
			break
		case "box of chocolates":
			assert.EqualValues(t, repeat*1, record.Quantity)
			assert.EqualValues(t, repeat*10.50, record.Total)
			break
		}
	}
}

func TestMixedTax(t *testing.T) {
	receipt := NewStandardReceipt().Add(
		product.Product{
			Name:     "bottle of perfume",
			Price:    27.99,
			Type:     product.OtherProductType,
			Imported: true,
		},
		product.Product{
			Name:  "bottle of perfume",
			Price: 18.99,
			Type:  product.OtherProductType,
		},
		product.Product{
			Name:  "packet of headache pills",
			Price: 9.75,
			Type:  product.MedicalProductType,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    11.25,
			Type:     product.FoodProductType,
			Imported: true,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    11.25,
			Type:     product.FoodProductType,
			Imported: true,
		},
		product.Product{
			Name:     "box of chocolates",
			Price:    11.25,
			Type:     product.FoodProductType,
			Imported: true,
		},
	).GetTotal()

	assert.EqualValues(t, 98.38, receipt.Total)
	assert.EqualValues(t, 7.90, receipt.Taxes)

	for _, record := range receipt.Records {
		if record.Imported {
			switch record.Name {
			case "bottle of perfume":
				assert.EqualValues(t, 1, record.Quantity)
				assert.EqualValues(t, 32.19, record.Total)
				break
			case "box of chocolates":
				assert.EqualValues(t, 3, record.Quantity)
				assert.EqualValues(t, 35.55, record.Total)
				break
			}
		} else {
			switch record.Name {
			case "bottle of perfume":
				assert.EqualValues(t, 1, record.Quantity)
				assert.EqualValues(t, 20.89, record.Total)
				break
			case "packet of headache pills":
				assert.EqualValues(t, 1, record.Quantity)
				assert.EqualValues(t, 9.75, record.Total)
				break
			}
		}
	}
}
