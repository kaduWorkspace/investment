package pg_repository

import (
	"context"
	"kaduhod/fin_v3/core/domain/repository"
	entitys "kaduhod/fin_v3/core/domain/user"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
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
func NewUserRepository(connection *pg_connection.PgxConnextion) repository.Repository[entitys.User] {
    return &UserRepository{
        pgx: connection,
    }
}
func (r *UserRepository) Save(fields entitys.User) (int, error) {
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
func (r *UserRepository) Update(fields entitys.User) error {
    return nil
}
func (r *UserRepository) Get(filters entitys.User) (entitys.User, error) {
    return entitys.User{}, nil
}
func (r *UserRepository) Delete(fields entitys.User) error {
    return nil
}
