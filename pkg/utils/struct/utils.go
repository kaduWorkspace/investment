package struct_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("erro ao realizar a requisição: %w", err)
	}
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
func SessionId(r *http.Request) string {
    // 1. Pega o User-Agent
    userAgent := r.UserAgent()

    // 2. Tenta pegar o IP do cabeçalho X-Forwarded-For (caso esteja atrás de proxy)
    ip := r.Header.Get("X-Forwarded-For")
    if ip == "" {
        // Se não tiver, usa o RemoteAddr
        ip = r.RemoteAddr
    } else {
        // Se o X-Forwarded-For tiver múltiplos IPs, pega o primeiro
        ip = strings.Split(ip, ",")[0]
    }

    // Se RemoteAddr contiver porta (tipo "192.168.0.1:12345"), remove
    if strings.Contains(ip, ":") {
        ip = strings.Split(ip, ":")[0]
    }

    return fmt.Sprintf("%s:%s", ip, userAgent)
}
func CreateCookie(w http.ResponseWriter) *http.Cookie {
	sessionID := uuid.New().String()
	cookie := &http.Cookie{
		Name:     "cookie",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(71 * time.Hour),
	}
	http.SetCookie(w, cookie)
	return cookie
}
func GetCookie(r *http.Request) *http.Cookie {
    cookie, err := r.Cookie("cookie")
    if err != nil {
        return nil
    }
    return cookie
}
func MinStringLength(fl validator.FieldLevel) bool {
	val := fl.Field()
	if val.Kind() != reflect.String {
		return false
	}
	str := val.String()
	param := fl.Param()
	minLen, err := strconv.Atoi(param)
	if err != nil {
		return false
	}
	return len(str) >= minLen
}
