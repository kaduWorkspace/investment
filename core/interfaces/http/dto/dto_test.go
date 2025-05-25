package http_dto

import (
	"testing"
)

type fakeFieldError struct {
	field string
	tag   string
}

func TestCoumpoundInterestInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   CoumpoundInterestInput
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   CoumpoundInterestInput{Periods: 12, TaxDecimal: 0.5, InitialValue: 1000},
			wantErr: false,
		},
		{
			name:    "missing initial value",
			input:   CoumpoundInterestInput{Periods: 12, TaxDecimal: 0.5},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, validationErrors := tt.input.Validate(tt.input)
            if err != nil {
                t.Errorf("%s: error on function", tt.name)
            }
			if tt.wantErr && validationErrors == nil{
                t.Errorf("%s: error expected", tt.name)
			}
		})
	}
}

func TestFutureValueOfASeriesInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   FutureValueOfASeriesInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: FutureValueOfASeriesInput{
				Periods: 12, TaxDecimal: 0.5,
				FirstDay: true, Contribution: 100,
			},
			wantErr: false,
		},
		{
			name: "invalid contribution",
			input: FutureValueOfASeriesInput{
				Periods: 12, TaxDecimal: 0.5,
				FirstDay: true, Contribution: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, validationErrors := tt.input.Validate(tt.input)
            if err != nil {
                t.Errorf("%s: error on function", tt.name)
            }
			if tt.wantErr && validationErrors == nil{
                t.Errorf("%s: error expected", tt.name)
			}
		})
	}
}

func TestPredictContributionFVSInput_Validate(t *testing.T) {
	tests := []struct {
		name    string
		input   PredictContributionFVSInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: PredictContributionFVSInput{
				Periods: 12, TaxDecimal: 0.5,
				FinalValue:             2000,
				InitialValue:           1000,
				ContributionOnFirstDay: true,
			},
			wantErr: false,
		},
		{
			name: "final value not greater than initial",
			input: PredictContributionFVSInput{
				Periods: 12, TaxDecimal: 0.5,
				FinalValue:             500,
				InitialValue:           1000,
				ContributionOnFirstDay: true,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, validationErrors := tt.input.Validate(tt.input)
            if err != nil {
                t.Errorf("%s: error on function", tt.name)
            }
			if tt.wantErr && validationErrors == nil{
                t.Errorf("%s: error expected", tt.name)
			}
		})
	}
}

