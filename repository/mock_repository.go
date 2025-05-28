package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetByID(model interface{}, id interface{}) error {
	args := m.Called(model, id)
	return args.Error(0)
}

func (m *MockRepository) GetAll(destSlicePtr interface{}, modelInstance interface{}, limit, offset *int, processor []QueryProcessor, preloads ...string) error {
	args := m.Called(destSlicePtr, modelInstance, limit, offset, processor, preloads)
	return args.Error(0)
}

func (m *MockRepository) GetWithFilters(model interface{}, processors []QueryProcessor) error {
	args := m.Called(model, processors)
	return args.Error(0)
}

func (m *MockRepository) Create(model interface{}) error {
	args := m.Called(model)
	return args.Error(0)
}
