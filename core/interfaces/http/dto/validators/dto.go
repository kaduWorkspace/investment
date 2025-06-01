package validators_dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)
type HttpInput struct {}
func (i HttpInput) FormatValidationError(err error, language string) map[string]string {
    var validationMessages = map[string]map[string]string{
        "required": {
            "pt": "campo obrigatório",
            "en": "required field",
        },
        "gte": {
            "pt": "deve ser maior ou igual a",
            "en": "must be greater than or equal to",
        },
        "gt": {
            "pt": "deve ser maior que",
            "en": "must be greater than",
        },
        "number": {
            "pt": "deve ser um número válido",
            "en": "must be a valid number",
        },
        "boolean": {
            "pt": "deve ser verdadeiro ou falso",
            "en": "must be true or false",
        },
        "datetime": {
            "pt": "deve ser uma data/hora válida",
            "en": "must be a valid date/time",
        },
        "gtfield": {
            "pt": "deve ser maior que",
            "en": "must be greater than",
        },
    }

    result := make(map[string]string)
    var validationErrors validator.ValidationErrors
    if errors.As(err, &validationErrors) {
        for _, e := range validationErrors {
            field := e.Field()
            tag := e.Tag()
            if messages, exists := validationMessages[tag]; exists {
                if msg, langExists := messages[language]; langExists {
                    if tag == "gtfield" || tag == "gte" || tag == "gt" {
                        msg = fmt.Sprintf("%s [%s]", msg, e.Param())
                    }
                    result[field] = msg
                    continue
                }
            }
            if language == "pt" {
                result[field] = fmt.Sprintf("validação falhou para '%s'", tag)
            } else {
                result[field] = fmt.Sprintf("validation failed for '%s'", tag)
            }
        }
    }

    if len(result) == 0 {
        if language == "pt" {
            result["_error"] = "dados da requisição inválidos"
        } else {
            result["_error"] = "invalid request data"
        }
    }
    return result
}

func (I HttpInput) Validate(dto interface{}) (error) {
    validate := validator.New(validator.WithPrivateFieldValidation())
    validate.RegisterValidation("datetime", func(fl validator.FieldLevel) bool {
        _, err := time.Parse("02/01/2006", fl.Field().String())
        return err == nil
    })
    err := validate.Struct(dto)
    if err == nil {
        return nil
    }
    var invalidValidationError *validator.InvalidValidationError
    if errors.As(err, &invalidValidationError) {
        return err
    }
    var validatonErrors validator.ValidationErrors
    if errors.As(err, &validatonErrors) {
        return err
    }
    return err
}

func (I HttpInput) FromJson(b []byte) (error, any) {
    var dto interface{}
    err := json.Unmarshal(b, &dto)
    return err, dto
}

type CoumpoundInterestInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gt=0,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"number,required,gte=1"`
}

type FutureValueOfASeriesInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gt=0,number"`
    FirstDay bool `json:"first_day" form:"first_day" validate:"boolean"`
    Contribution float64 `json:"contribution" form:"contribution" validate:"gt=0,number"`
}

type FutureValueOfASeriesWithPeriodsInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gt=0,number"`
    FirstDay bool `json:"first_day" form:"first_day" validate:"boolean"`
    Contribution float64 `json:"contribution" form:"contribution" validate:"number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"gte=1,number"`
    InitialDate string `json:"initial_date" form:"initial_date" validate:"required,datetime"`
}

type PredictContributionFVSInput struct {
    HttpInput
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gt=0,number"`
    FinalValue float64 `json:"final_value" form:"final_value" validate:"required,gte=1,gtfield=InitialValue,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"gte=0,number"`
    ContributionOnFirstDay bool `json:"first_day" form:"contribution_on_first_day" validate:"required,boolean"`
}

