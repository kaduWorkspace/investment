package dto

type Dto interface {
    Validate() (error)
    FormatValidationError(err error, language string) map[string]string
}
