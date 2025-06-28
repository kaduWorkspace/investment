package validators_dto

import (
	"kaduhod/fin_v3/core/domain/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignInInput_Validations(t *testing.T) {
	tests := []struct {
		name     string
		input    dto.Dto
		wantErr  bool
		errField string
		errTag   string
	}{
		{
			name:     "empty email",
			input:    NewSignInInput("", "password"),
			wantErr:  true,
			errField: "Email",
			errTag:   "required",
		},
		{
			name:     "invalid email format",
			input:    NewSignInInput("invalid-email", "password"),
			wantErr:  true,
			errField: "Email",
			errTag:   "email",
		},
		{
			name:     "empty password",
			input:    NewSignInInput("valid@email.com", ""),
			wantErr:  true,
			errField: "Password",
			errTag:   "required",
		},
		{
			name:    "valid input",
			input:   NewSignInInput("valid@email.com", "password"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.wantErr && err == nil {
				t.Log("Error expected")
				t.Fail()
			}
			if tt.wantErr && err != nil {
				assert.Error(t, err)
				errs := tt.input.FormatValidationError(err, "en")
				assert.Contains(t, errs, tt.errField)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSignInInput_Translation(t *testing.T) {
	tests := []struct {
		lang     string
		tag      string
		param    string
		expected string
	}{
		{"en", "required", "", "required field"},
		{"pt", "required", "", "campo obrigatório"},
		{"en", "email", "", "must be a valid email"},
		{"pt", "email", "", "deve ser um email válido"},
	}

	for _, tt := range tests {
		t.Run(tt.lang+"_"+tt.tag, func(t *testing.T) {
			input := NewSignInInput("", "")
			err := input.Validate()
			errs := input.FormatValidationError(err, tt.lang)

			if tt.tag == "required" {
				assert.Contains(t, errs["Email"], tt.expected)
			}
		})
	}
}
