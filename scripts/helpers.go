package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}

func Validate(data interface{}) error {
	validate := validator.New()
	return validate.Struct(data)
}

func IsBlank(s string) bool {
	return s == "" || strings.TrimSpace(s) == ""
}

func IsEmailValid(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func GetHeader(key string, r *http.Request) string {
	if v := r.Header.Get(key); v!= "" {
		return v
	}
	return ""
}

func GetIP(r *http.Request) string {
	ip := GetHeader("X-Forwarded-For", r)
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func GetCountry(ip string) (string, error) {
	resp, err := http.Get("https://ip-api.com/json/" + ip)
	if err!= nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err!= nil {
		return "", err
	}

	return data["country"], nil
}

func GetCity(ip string) (string, error) {
	resp, err := http.Get("https://ip-api.com/json/" + ip)
	if err!= nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err!= nil {
		return "", err
	}

	return data["city"], nil
}

func GetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDateStr() string {
	return time.Now().Format("2006-01-02")
}

func Str2Int(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err!= nil {
		return 0, err
	}
	return i, nil
}

func Str2Float(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err!= nil {
		return 0, err
	}
	return f, nil
}

func LogError(err error) {
	log.Println(err)
}

func LogErrorWithMsg(err error, msg string) {
	log.Printf("%s: %v", msg, err)
}

func LogInfo(msg string) {
	log.Println(msg)
}

func LogDebug(msg string) {
	log.Println(msg)
}