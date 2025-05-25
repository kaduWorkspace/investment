package struct_utils

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func FromJson[T any](b []byte) (error, T) {
    var dto T
    err := json.Unmarshal(b, &dto)
    if err == nil {
        return nil, dto
    }
    re := regexp.MustCompile(`cannot unmarshal (\w+) into Go struct field \w+\.(\w+) of type (\w+)`)
    match := re.FindStringSubmatch(err.Error())
    if len(match) != 4 {
        return err, dto
    }
    receivedType := match[1]
    field := match[2]
    expectedType := match[3]
    if expectedType == "float64" || expectedType == "int" {
        expectedType = "number"
    }
    err = fmt.Errorf("%s should be %s, but got %s", field, expectedType, receivedType)
    return err, dto
}
