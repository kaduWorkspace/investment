package app_account_service_test

import (
	"context"
	"testing"

	app_account_dto "kaduhod/fin_v3/core/application/account/dto"
	app_account_service "kaduhod/fin_v3/core/application/account/service"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	pg_repository "kaduhod/fin_v3/core/infra/persistence/postgres/repository"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupDB(t *testing.T) *pg_connection.PgxConnextion {
	t.Helper()
    err := godotenv.Load("../../../../../.env.development")
    if err != nil {
        t.Error(err)
        t.Fail()
    }

	conn := pg_connection.NewPgxConnection()
	return conn
}

func TestCreateUserService_Integration(t *testing.T) {
	conn := setupDB(t)
	defer cleanupTestUsers(t, conn) // This will run after all tests complete
	userRepo := pg_repository.NewUserRepository(conn)
	userService := app_account_service.NewCreateUserService(userRepo)

	t.Run("Create user successfully", func(t *testing.T) {
		input := app_account_dto.CreateUserInput{
			Name:     "Test User",
			Email:    "test1@example.com",
			Password: "password123!",
		}

		err := userService.Create(input)
		assert.NoError(t, err)
	})

	t.Run("Fail to create user with duplicate email", func(t *testing.T) {
		input := app_account_dto.CreateUserInput{
			Name:     "Test User",
			Email:    "test2@example.com",
			Password: "password123!",
		}

		// First creation should succeed
		err := userService.Create(input)
		assert.NoError(t, err)

		// Second creation with same email should fail
		err = userService.Create(input)
		assert.Error(t, err)
		assert.Equal(t, "User email not available", err.Error())
	})

	t.Run("Fail to create user with invalid password", func(t *testing.T) {
		input := app_account_dto.CreateUserInput{
			Name:     "Test User",
			Email:    "test3@example.com",
			Password: "short", // Doesn't meet validation requirements
		}

		err := input.Validate()
		assert.Error(t, err)
	})
}
func cleanupTestUsers(t *testing.T, conn *pg_connection.PgxConnextion) {
	t.Helper()
	_, err := conn.Conn.Exec(context.Background(), `
		DELETE FROM users
		WHERE email IN (
			'test1@example.com',
			'test2@example.com',
			'test3@example.com'
		)
	`)
	if err != nil {
		t.Logf("Warning: failed to clean up test users: %v", err)
	}
}
