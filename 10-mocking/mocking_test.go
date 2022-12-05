package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) MockAdder(values []int) (int, error) {
	args := mock.Called(values)

	return args.Get(0).(int), args.Error(1)
}

func TestMockAdder(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("MockAdder", []int{2, 3}).Return(5, nil)

	testMathInterface := MathInterface(mockRepo)

	got, err := testMathInterface.MockAdder([]int{2, 3})

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 5, got)
	assert.Nil(t, err)
}
