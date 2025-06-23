package validators_dto

import (
	"testing"
)

func TestCoumpoundInterestInput_Validate(t *testing.T) {
    tests := []struct {
        name    string
        input   CoumpoundInterestInput
        wantErr bool
        messagesEn map[string]string
        messagesPt map[string]string
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
            messagesPt: map[string]string{"InitialValue":"campo obrigatório"},
            messagesEn: map[string]string{"InitialValue":"required field"},
        },
        {
            name:    "Required tax decimal",
            input:   CoumpoundInterestInput{Periods: 12, TaxDecimal: 0, InitialValue: 1},
            wantErr: true,
            messagesPt: map[string]string{"TaxDecimal":"campo obrigatório"},
            messagesEn: map[string]string{"TaxDecimal":"required field"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.input.Validate(tt.input)
            if tt.wantErr && err == nil{
                t.Errorf("%s: error expected", tt.name)
            }
            if tt.wantErr {
                enMessages := tt.input.FormatValidationError(err, "en")
                for key, value := range tt.messagesEn{
                    if enMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", enMessages[key], key, value)
                    }
                }
                ptMessages := tt.input.FormatValidationError(err, "pt")
                for key, value := range tt.messagesPt{
                    if ptMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", ptMessages[key], key, value)
                    }
                }
            }
        })
    }
}

func TestFutureValueOfASeriesInput_Validate(t *testing.T) {
    tests := []struct {
        name    string
        input   FutureValueOfASeriesInput
        wantErr bool
        messagesEn map[string]string
        messagesPt map[string]string
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
            messagesPt: map[string]string{"Contribution":"deve ser maior que [0]"},
            messagesEn: map[string]string{"Contribution":"must be greater than [0]"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.input.Validate(tt.input)
            if tt.wantErr && err == nil{
                t.Errorf("%s: error expected", tt.name)
            }
            if tt.wantErr {
                enMessages := tt.input.FormatValidationError(err, "en")
                for key, value := range tt.messagesEn{
                    if enMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", enMessages[key], key,value)
                    }
                }
                ptMessages := tt.input.FormatValidationError(err, "pt")
                for key, value := range tt.messagesPt{
                    if ptMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", ptMessages[key], key, value)
                    }
                }
            }
        })
    }
}

func TestPredictContributionFVSInput_Validate(t *testing.T) {
    tests := []struct {
        name    string
        input   PredictContributionFVSInput
        wantErr bool
        messagesEn map[string]string
        messagesPt map[string]string
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
            messagesEn: map[string]string{"FinalValue":"must be greater than [InitialValue]"},
            messagesPt: map[string]string{"FinalValue":"deve ser maior que [InitialValue]"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.input.Validate(tt.input)
            if tt.wantErr && err == nil{
                t.Errorf("%s: error expected", tt.name)
            }
            if tt.wantErr {
                enMessages := tt.input.FormatValidationError(err, "en")
                for key, value := range tt.messagesEn{
                    if enMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", enMessages[key], key ,value)
                    }
                }
                ptMessages := tt.input.FormatValidationError(err, "pt")
                for key, value := range tt.messagesPt{
                    if ptMessages[key] != value {
                        t.Errorf("%s[%s]: wrong massege, want: %s", enMessages[key], key ,value)
                    }
                }
            }
        })
    }
}
func TestFutureValueOfASeriesWithPeriodsInput_Validate(t *testing.T) {
    tests := []struct {
        name      string
        input     FutureValueOfASeriesWithPeriodsInput
        wantErr   bool
        messagesEn map[string]string
        messagesPt map[string]string
    }{
        {
            name: "valid with contribution only",
            input: FutureValueOfASeriesWithPeriodsInput{
                Periods: 12, TaxDecimal: 0.5,
                FirstDay: true, Contribution: 100,
                InitialValue: 0, InitialDate: "01/01/2023",
            },
            wantErr: false,
        },
        {
            name: "valid with initial value only",
            input: FutureValueOfASeriesWithPeriodsInput{
                Periods: 12, TaxDecimal: 0.5,
                FirstDay: true, Contribution: 0,
                InitialValue: 1000, InitialDate: "01/01/2023",
            },
            wantErr: false,
        },
        {
            name: "valid with both values",
            input: FutureValueOfASeriesWithPeriodsInput{
                Periods: 12, TaxDecimal: 0.5,
                FirstDay: true, Contribution: 100,
                InitialValue: 1000, InitialDate: "01/01/2023",
            },
            wantErr: false,
        },
        {
            name: "invalid - both zero",
            input: FutureValueOfASeriesWithPeriodsInput{
                Periods: 12, TaxDecimal: 0.5,
                FirstDay: true, Contribution: 0,
                InitialValue: 0, InitialDate: "01/01/2023",
            },
            wantErr: true,
            messagesEn: map[string]string{"Contribution": "at least one between contribution or initial_value must be greater than zero", "InitialValue": "at least one between contribution or initial_value must be greater than zero"},
            messagesPt: map[string]string{"Contribution": "pelo menos um entre contribution ou initial_value deve ser maior que zero", "InitialValue": "pelo menos um entre contribution ou initial_value deve ser maior que zero"},
        },
        {
            name: "invalid - missing initial date",
            input: FutureValueOfASeriesWithPeriodsInput{
                Periods: 12, TaxDecimal: 0.5,
                FirstDay: true, Contribution: 100,
                InitialValue: 1000, InitialDate: "",
            },
            wantErr: true,
            messagesEn: map[string]string{"InitialDate": "required field"},
            messagesPt: map[string]string{"InitialDate": "campo obrigatório"},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.input.Validate(tt.input)
            if tt.wantErr && err == nil {
                t.Errorf("%s: error = %v, wantErr %v", tt.name, err, tt.wantErr)
            }
            if tt.wantErr {
                enMessages := tt.input.FormatValidationError(err, "en")
                for key, value := range tt.messagesEn {
                    if enMessages[key] != value {
                        t.Errorf("%s[en]: wrong message for %s, got: %s, want: %s", tt.name, key, enMessages[key], value)
                    }
                }
                ptMessages := tt.input.FormatValidationError(err, "pt")
                for key, value := range tt.messagesPt {
                    if ptMessages[key] != value {
                        t.Errorf("%s[pt]: wrong message for %s, got: %s, want: %s", tt.name, key, ptMessages[key], value)
                    }
                }
            }
        })
    }
}
