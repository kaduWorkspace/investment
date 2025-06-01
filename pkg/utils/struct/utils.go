package struct_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
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
func HttpRequest(url string, method string, headers map[string]string, body string) (string, error) {
	if method == "" {
		method = "GET"
	}
	var requestBody *bytes.Reader
	if body != "" {
		requestBody = bytes.NewReader([]byte(body))
	} else {
		requestBody = bytes.NewReader(nil)
	}
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return "", fmt.Errorf("erro ao criar a requisição: %w", err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erro ao realizar a requisição: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler a resposta: %w", err)
	}
	return string(respBody), nil
}
func EhMobile(userAgent string) bool {
	mobileKeywords := []string{
		"Android",
		"iPhone",
		"iPad",
		"Windows Phone",
		"BlackBerry",
		"Mobile",
	}
	for _, keyword := range mobileKeywords {
		if strings.Contains(userAgent, keyword) {
			return true
		}
	}
	return false
}
