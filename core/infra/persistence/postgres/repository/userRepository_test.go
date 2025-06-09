package pg_repository

import (
	"context"
	"fmt"
	"kaduhod/fin_v3/core/domain/user"
	"kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Save(t *testing.T) {
    err := godotenv.Load("/home/carlos/projetos/meu-app/chi_version/.env")
    if err != nil {
        fmt.Println(err)
        log.Fatal("Error loading .env file")
    }

	// Setup test database connection
	ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
	defer conn.Conn.Close()

	// Create repository
	repo := NewUserRepository(conn)

	// Clean up test data
	defer func() {
		_, _ = conn.Conn.Exec(ctx, "DELETE FROM user WHERE email LIKE 'test_%'")
	}()

	t.Run("successfully saves user", func(t *testing.T) {
		testUser := entitys.User{
			Name:     "Test User",
			Email:    "test_user@example.com",
			Password: "securepassword",
		}

		id, err := repo.Save(testUser)
		assert.NoError(t, err)
		assert.Greater(t, id, 0)

		// Verify the record exists
		var dbID int
		err = conn.Conn.QueryRow(ctx, "SELECT id FROM user WHERE email = $1", testUser.Email).Scan(&dbID)
		assert.NoError(t, err)
		assert.Equal(t, id, dbID)
	})

	t.Run("returns error for duplicate email", func(t *testing.T) {
		testUser := entitys.User{
			Name:     "Duplicate User",
			Email:    "duplicate@example.com",
			Password: "password",
		}

		// First insert should succeed
		_, err := repo.Save(testUser)
		assert.NoError(t, err)

		// Second insert should fail
		_, err = repo.Save(testUser)
		assert.Error(t, err)
	})

	t.Run("returns error for empty required fields", func(t *testing.T) {
		testCases := []struct {
			name  string
			user  entitys.User
			field string
		}{
			{"empty name", entitys.User{Email: "test1@example.com", Password: "pass"}, "name"},
			{"empty email", entitys.User{Name: "Test", Password: "pass"}, "email"},
			{"empty password", entitys.User{Name: "Test", Email: "test2@example.com"}, "password"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := repo.Save(tc.user)
				assert.Error(t, err)
			})
		}
	})
}
