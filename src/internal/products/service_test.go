package products

import (
	"testing"

	"github.com/caiowWillian/first-crud-golang/src/pkg/databases/mongo/mockMongo"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	mocked := new(mockMongo.Repository)

	mocked.On("Insert", mock.Anything).Return(nil)
	id, err := NewService(mocked).CreateProduct(Product{})

	assert.Equal(t, len(id) > 0, true)
	assert.NoError(t, err)
}

func TestGetAll(t *testing.T) {
	mocked := new(mockMongo.Repository)
	mocked.On("GetAll", mock.Anything, mock.Anything).Return(nil)

	_, err := NewService(mocked).GetAllProducts()
	assert.NoError(t, err)
}
