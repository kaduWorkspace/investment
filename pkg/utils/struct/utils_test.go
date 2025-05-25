package struct_utils

import (
	"testing"
)
type test struct {
    Periods int `json:"periods" form:"periods" validate:"required,gte=1,number"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal" validate:"required,gt=0,number"`
    InitialValue float64 `json:"initial_value" form:"initial_value" validate:"number,required,gte=1"`
}
func TestInvalidJsonToStruct(t *testing.T) {
    t.Run("Invalid string value", func(t *testing.T) {
        json_value := []byte(`{"initial_value":"afdsdsf","periods":"13","tax_decimal":0.1475}`)
        err, _ := FromJson[test](json_value)
        if err == nil {
            t.Error("Invalid json to struct should not be permitted!")
        }
        if err.Error() != "initial_value should be number, but got string" {
            t.Errorf("Invalid json to struct error message not expected: %s", err.Error())
        }
    })
}
