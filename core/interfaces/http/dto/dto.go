package http_dto

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)
type HttpInput struct {}
func (i HttpInput) FormatValidationErrorToPortuguese(err error) string {
    var validationMessages = map[string]string{
        "required":  "campo obrigatório",
        "gte":       "deve ser maior ou igual a 0",  // More specific message for TaxDecimal
        "number":    "deve ser um número válido",
        "boolean":      "deve ser verdadeiro ou falso",
        "datetime":  "deve ser uma data/hora válida",
        "gtfield":   "deve ser maior que o valor de referência",
    }
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			field := e.Field()
			tag := e.Tag()
			if msg, exists := validationMessages[tag]; exists {
				return fmt.Sprintf("%s: %s", field, msg)
			}
			return fmt.Sprintf("%s: validação falhou para '%s'", field, tag)
		}
	}
	return "dados da requisição inválidos"
}

func (I HttpInput) Validate(dto interface{}) (error, interface{}) {
    validate := validator.New(validator.WithPrivateFieldValidation())
    err := validate.Struct(dto)
    if err != nil {
        return nil, err
    }
    var invalidValidationError *validator.InvalidValidationError
    if errors.As(err, &invalidValidationError) {
        return err, nil
    }
    var validatonErrors validator.ValidationErrors
    if errors.As(err, &validatonErrors) {
        return err, nil
    }
    return err, nil
}

type CoumpoundInterestInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gte=1,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"required,gte=1,number"`
}

type FutureValueOfASeriesInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gte=1,number"`
    FirstDay bool `json:"first_day" form:"first_day" validate:"boolean"`
    Contribution float64 `json:"contribution" form:"contribution" validate:"gte=1,number"`
}

type FutureValueOfASeriesWithPeriodsInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gte=1,number"`
    FirstDay bool `json:"first_day" form:"first_day" validate:"boolean"`
    Contribution float64 `json:"contribution" form:"contribution" validate:"gte=1,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"gte=1,number"`
    InitialDate time.Time `json:"initial_date" form:"initial_date" validate:"required,datetime"`
}

type PredictContributionFVSInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gte=1,number"`
    FinalValue float64 `json:"final_value" form:"final_value" validate:"required,gte=1,gtfield=InitialValue,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"required,gte=1,number"`
    ContributionOnFirstDay bool `json:"contribution_on_first_day" form:"contribution_on_first_day" validate:"required,boolean"`
}

