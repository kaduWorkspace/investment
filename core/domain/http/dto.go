package http

type Dto interface {
    Validate() (error, interface{})
    FormatValidationError(err error, language string) map[string]string
}
