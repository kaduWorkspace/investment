package main_chi

import (
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	interface_chi "kaduhod/fin_v3/core/interfaces/http/handlers/chi"
)

func Run() {
    conn := pg_connection.NewPgxConnection()
    server := interface_chi.ServerChi{
        Conn: conn,
    }
    server.Setup()
    defer server.Shutdown()
    server.Start(":8989")
}
