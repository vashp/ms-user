package services

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/vashp/ms-user/src/cmd/domain"
	"github.com/vashp/ms-user/src/cmd/infrastructure/connection"
	"github.com/vashp/ms-user/src/cmd/infrastructure/repositories"
	"testing"
)

func TestCreateIntegration(t *testing.T) {
	db, err := connection.TestConnectPostgres()
	assert.NoError(t, err)

	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &domain.User{Email: "integration-test@example.com"}
	err = service.Create(context.Background(), user)
	assert.NoError(t, err)

	createdUser, err := repo.GetByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "integration-test@example.com", createdUser.Email)
}
