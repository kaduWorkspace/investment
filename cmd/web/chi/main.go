package main_chi

import (
	interface_chi "kaduhod/fin_v3/core/interfaces/http/handlers/chi"
)

func Run() {
    server := interface_chi.ServerChi{}
    server.Setup()
    server.Start(":8989")
}
