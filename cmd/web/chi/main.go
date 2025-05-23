package main

import (
	interface_chi "kaduhod/fin_v3/core/interfaces/http/handlers/chi"
)

func main() {
    server := interface_chi.ServerChi{}
    server.Setup()
    server.Start(":8989")
}
