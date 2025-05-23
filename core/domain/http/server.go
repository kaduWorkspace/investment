package http

type Server interface {
    Setup()
    Start(port string)
}
