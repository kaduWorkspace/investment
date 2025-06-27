package validators_dto

import (
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/dto"

	"github.com/go-playground/validator/v10"
)

type SignInInput struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

func NewSignInInput(Email, Password string) dto.Dto {
    return SignInInput{
        Email:    Email,
        Password: Password,
    }
}
func (i SignInInput) FormatValidationError(err error, language string) map[string]string {
    var validationMessages = map[string]map[string]string{
        "required": {
            "pt": "campo obrigatório",
            "en": "required field",
        },
        "email": {
            "pt": "deve ser um email válido",
            "en": "must be a valid email",
        },
        "min": {
            "pt": "deve ter no mínimo %s caracteres",
            "en": "must be at least %s characters",
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
                    if tag == "min" || tag == "containsany" || tag == "eqfield" {
                        msg = fmt.Sprintf(msg, e.Param())
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
func (i SignInInput) Validate() error {
    validate := validator.New()
    return validate.Struct(i)
}
