package pg_connection

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxConnextion struct {
    Conn *pgxpool.Pool
}
func NewPgxConnection() *PgxConnextion {
    ctx := context.Background()
    host := os.Getenv("DATABASE_HOST")
    user := os.Getenv("DATABASE_USER")
    pass := os.Getenv("DATABASE_PWD")
    port := os.Getenv("DATABASE_PORT")
    db := os.Getenv("DATABASE_DB")
    url := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/" + db
    conn, err := pgxpool.New(ctx, url)
    if err != nil {
        panic(err)
    }
    return &PgxConnextion{
        Conn: conn,
    }
}
