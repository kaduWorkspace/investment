package app_account_dto

import (
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/dto"

	"github.com/go-playground/validator/v10"
)

type CreateUserInput struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6,containsany=!@#$%^&*(),containsany=0123456789"`
}

func NewCreateUserInput(name, email, password string) dto.Dto {
    return CreateUserInput{
        Name:     name,
        Email:    email,
        Password: password,
    }
}

func (input CreateUserInput) Validate() error {
    validate := validator.New()
    return validate.Struct(input)
}

func (input CreateUserInput) FormatValidationError(err error, language string) map[string]string {
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
