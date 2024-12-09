package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vashp/ms-user/src/cmd/domain"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)
	if user, ok := args.Get(0).(*domain.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &domain.User{Email: "test@example.com"}
	mockRepo.On("Create", mock.Anything, user).Return(nil)

	err := service.Create(context.Background(), user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateWithEmptyEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &domain.User{Email: ""}
	err := service.Create(context.Background(), user)
	assert.Error(t, err)
	assert.Equal(t, "email is required", err.Error())

	mockRepo.AssertExpectations(t)
}
