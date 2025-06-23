package pg_repository

import (
	"context"
	"kaduhod/fin_v3/core/domain/user"
	"kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)
func TestMain(m *testing.M) {

    enviroment := os.Args[len(os.Args) - 1]
    envFile := ".env.development"
    if enviroment == "local" {
        envFile = ".env.local"
    }
    err := godotenv.Load("../../../../../" + envFile)
    if err != nil {
        log.Fatal("Erro ao carregar .env:", err)
    }

    os.Exit(m.Run())
}

func TestUserRepository_Save(t *testing.T) {
	ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
	defer conn.Conn.Close()

	repo := NewUserRepository(conn)

	defer func() {
		_, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email LIKE 'test_%' or email like 'duplicate_%'")
	}()

	t.Run("successfully saves user", func(t *testing.T) {
		testUser := user.User{
			Name:     "Test User",
			Email:    "test_user@example.com",
			Password: "securepassword",
		}

		id, err := repo.Save(testUser)
		assert.NoError(t, err)
		assert.Greater(t, id, 0)

		var dbID int
		err = conn.Conn.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", testUser.Email).Scan(&dbID)
		assert.NoError(t, err)
		assert.Equal(t, id, dbID)
	})

	t.Run("returns error for duplicate email", func(t *testing.T) {
		testUser := user.User{
			Name:     "Duplicate User",
			Email:    "duplicate@example.com",
			Password: "password",
		}

		_, err := repo.Save(testUser)
		assert.NoError(t, err)

		_, err = repo.Save(testUser)
		assert.Error(t, err)
	})
}
func TestUserRepository_Get(t *testing.T) {
    ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
    defer conn.Conn.Close()

    repo := NewUserRepository(conn)

    testUser := user.User{
        Name:     "Test Get User",
        Email:    "test_get_user@example.com",
        Password: "password",
    }
    id, err := repo.Save(testUser)
    assert.NoError(t, err)

    defer func() {
        _, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email = $1", testUser.Email)
    }()

    t.Run("successfully gets user by ID", func(t *testing.T) {
        filters := user.User{Id: id}
        result, err := repo.Get(filters)

        assert.NoError(t, err)
        assert.Equal(t, id, result.Id)
        assert.Equal(t, testUser.Name, result.Name)
        assert.Equal(t, testUser.Email, result.Email)
    })

    t.Run("successfully gets user by email", func(t *testing.T) {
        filters := user.User{Email: testUser.Email}
        result, err := repo.Get(filters)

        assert.NoError(t, err)
        assert.Equal(t, id, result.Id)
        assert.Equal(t, testUser.Name, result.Name)
        assert.Equal(t, testUser.Email, result.Email)
    })

    t.Run("returns error when no filters provided", func(t *testing.T) {
        filters := user.User{}
        _, err := repo.Get(filters)

        assert.Error(t, err)
        assert.Equal(t, "no filter criteria provided", err.Error())
    })

    t.Run("returns error when user doesn't exist", func(t *testing.T) {
        filters := user.User{Email: "nonexistent@example.com"}
        _, err := repo.Get(filters)

        assert.Error(t, err)
        assert.Contains(t, err.Error(), "failed to get user")
    })
}
func TestUserRepository_Update(t *testing.T) {
    ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
    defer conn.Conn.Close()

    repo := NewUserRepository(conn)

    testUser := user.User{
        Name:     "Test Update User",
        Email:    "test_update_user@example.com",
        Password: "password",
    }
    id, err := repo.Save(testUser)
    assert.NoError(t, err)

    defer func() {
        _, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email LIKE 'test_update%' or LIKE 'updated_email%'")
    }()

    t.Run("successfully updates user", func(t *testing.T) {
        updatedUser := user.User{
            Id:       id,
            Name:     "Updated Name",
            Email:    "updated_email@example.com",
            Password: "newpassword",
        }

        err := repo.Update(updatedUser)
        assert.NoError(t, err)

        result, err := repo.Get(user.User{Id: id})
        assert.NoError(t, err)
        assert.Equal(t, updatedUser.Name, result.Name)
        assert.Equal(t, updatedUser.Email, result.Email)
    })

    t.Run("returns error when no ID provided", func(t *testing.T) {
        err := repo.Update(user.User{})
        assert.Error(t, err)
        assert.Equal(t, "user ID is required for update", err.Error())
    })

    t.Run("returns error when user doesn't exist", func(t *testing.T) {
        err := repo.Update(user.User{Id: 999999})
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "failed to update user")
    })
    _, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email = 'updated_email@example.com'")
    _, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email = 'test_update_user@example.com'")

}

func TestUserRepository_Delete(t *testing.T) {
    ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
    defer conn.Conn.Close()

    repo := NewUserRepository(conn)

    testUser := user.User{
        Name:     "Test Delete User",
        Email:    "test_delete_user@example.com",
        Password: "password",
    }
    id, err := repo.Save(testUser)
    assert.NoError(t, err)

    defer func() {
        _, _ = conn.Conn.Exec(ctx, "DELETE FROM users WHERE email LIKE 'test_delete%'")
    }()

    t.Run("successfully deletes user", func(t *testing.T) {
        err := repo.Delete(user.User{Id: id})
        assert.NoError(t, err)

        var deletedAt *time.Time
        err = conn.Conn.QueryRow(ctx,
            "SELECT deleted_at FROM users WHERE id = $1", id).Scan(&deletedAt)
        assert.NoError(t, err)
        assert.NotNil(t, deletedAt)
    })

    t.Run("returns error when no ID provided", func(t *testing.T) {
        err := repo.Delete(user.User{})
        assert.Error(t, err)
        assert.Equal(t, "user ID is required for deletion", err.Error())
    })

    t.Run("returns error when user doesn't exist", func(t *testing.T) {
        err := repo.Delete(user.User{Id: 999999})
        assert.Error(t, err)
        assert.Equal(t, "user not found or already deleted", err.Error())
    })
}
