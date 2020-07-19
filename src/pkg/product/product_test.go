package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	wanted := "music CD14.99othertrue"
	product := Product{"music CD", 14.99, OtherProductType, true}
	assert.Equal(t, wanted, product.Hash())
}
