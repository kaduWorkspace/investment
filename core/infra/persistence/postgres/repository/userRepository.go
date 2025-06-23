package pg_repository

import (
	"context"
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/repository"
	user "kaduhod/fin_v3/core/domain/user"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"strings"
)

/*type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}*/
type UserRepository struct {
    pgx *pg_connection.PgxConnextion
}
func NewUserRepository(connection *pg_connection.PgxConnextion) repository.Repository[user.User] {
    return &UserRepository{
        pgx: connection,
    }
}
func (r *UserRepository) Save(fields user.User) (int, error) {
    ctx := context.Background()
    tx, err := r.pgx.Conn.Begin(ctx)
    var id int
    if err != nil {
        return id, err
    }
    err = tx.QueryRow(ctx, "insert into users (name, email, password) values ($1, $2, $3) returning id", fields.Name, fields.Email, fields.Password).Scan(&id);
    if err != nil {
        tx.Rollback(ctx)
        return id, err
    }
    tx.Commit(ctx)
    return id, nil
}

func (r *UserRepository) Get(filters user.User) (user.User, error) {
    ctx := context.Background()
    var user user.User

    // Build WHERE clause and arguments
    var whereClause strings.Builder
    var args []interface{}
    var conditions []string

    if filters.Id != 0 {
        args = append(args, filters.Id)
        conditions = append(conditions, "id = $1")
    } else {
        cont := 1
        if filters.Name != "" {
            args = append(args, filters.Name)
            conditions = append(conditions, fmt.Sprintf("name = $%d", cont))
            cont++
        }
        if filters.Email != "" {
            args = append(args, filters.Email)
            conditions = append(conditions, fmt.Sprintf("email = $%d", cont))
            cont++
        }
    }
    if len(args) == 0 {
        return user, errors.New("no filter criteria provided")
    }

    whereClause.WriteString("WHERE ")
    whereClause.WriteString(strings.Join(conditions, " AND "))

    query := "SELECT id, name, email, password FROM users " + whereClause.String()
    err := r.pgx.Conn.QueryRow(ctx, query, args...).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
    if err != nil {
        return user, fmt.Errorf("failed to get user: %w", err)
    }

    return user, nil
}

func (r *UserRepository) Update(user user.User) error {
    if user.Id == 0 {
        return errors.New("user ID is required for update")
    }

    ctx := context.Background()
    userDb, err := r.Get(user)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    tx, err := r.pgx.Conn.Begin(ctx)
    if err != nil {
        return err
    }
    defer tx.Rollback(ctx)

    var setClause strings.Builder
    var args []string
    var values []interface{}
    paramCount := 1

    if userDb.Email != user.Email {
        args = append(args, fmt.Sprintf("email = $%d", paramCount))
        values = append(values, user.Email)
        paramCount++
    }

    if userDb.Password != user.Password {
        args = append(args, fmt.Sprintf("password = $%d", paramCount))
        values = append(values, user.Password)
        paramCount++
    }

    if userDb.Name != user.Name {
        args = append(args, fmt.Sprintf("name = $%d", paramCount))
        values = append(values, user.Name)
        paramCount++
    }

    if len(args) == 0 {
        return errors.New("no fields to update")
    }

    setClause.WriteString("SET ")
    setClause.WriteString(strings.Join(args, ", "))
    values = append(values, user.Id)

    query := fmt.Sprintf("UPDATE users %s, updated_at = NOW() WHERE id = $%d",
        setClause.String(),
        paramCount)

    _, err = tx.Exec(ctx, query, values...)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    return tx.Commit(ctx)
}


func (r *UserRepository) Delete(user user.User) error {
    ctx := context.Background()

    if user.Id == 0 {
        return errors.New("user ID is required for deletion")
    }

    tx, err := r.pgx.Conn.Begin(ctx)
    if err != nil {
        return err
    }

    result, err := tx.Exec(ctx,
        "UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL",
        user.Id)
    if err != nil {
        tx.Rollback(ctx)
        return fmt.Errorf("failed to delete user: %w", err)
    }

    if rows := result.RowsAffected(); rows == 0 {
        tx.Rollback(ctx)
        return errors.New("user not found or already deleted")
    }

    return tx.Commit(ctx)
}
