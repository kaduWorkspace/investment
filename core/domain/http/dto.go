package http

type Dto interface {
    Validate() (error, interface{})
}
