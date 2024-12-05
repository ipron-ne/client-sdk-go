package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"net/url"
	"strings"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func JSONParse(data string) map[string]any {
	var result map[string]any
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		log.Printf("Failed to parse JSON: %s", err)
		return nil
	}
	return result
}

func GetStr(data map[string]any, key string) string {
	d := data[key]
	if d == nil {
		return ""
	}
	return d.(string)
}

func GetInt(data map[string]any, key string) int {
	d := data[key]
	if d == nil {
		return 0
	}
	return d.(int)
}

func GetBool(data map[string]any, key string) bool {
	d := data[key]
	if d == nil {
		return false
	}
	return d.(bool)
}

func GetMap(data map[string]any, key string) map[string]any {
	d := data[key]
	if d == nil {
		return nil
	}

	m, ok := d.(map[string]any)
	if !ok {
		return nil
	}

	return m
}

// Helper function to format query parameters
func ParamsSerializer(params map[string]any) string {
	// url.Values는 기본 쿼리 문자열 저장용 맵
	values := url.Values{}

	for key, value := range params {
		switch v := value.(type) {
		case []int:
			// 숫자 배열 값을 쉼표로 연결
			strVals := make([]string, len(v))
			for i, num := range v {
				strVals[i] = strconv.Itoa(num)
			}
			values.Set(key, strings.Join(strVals, ","))
		case []string:
			// 배열 값을 쉼표로 연결
			values.Set(key, strings.Join(v, ","))
		case string:
			// 단일 문자열
			values.Set(key, v)
		case int:
			// 단일 숫자
			values.Set(key, strconv.Itoa(v))
		default:
			// 기타 타입은 문자열로 변환
			values.Set(key, fmt.Sprintf("%v", v))
		}
	}

	// Encode()로 쿼리 문자열 생성
	return values.Encode()
}

// createUUID: UUID 생성
func CreateUUID() string {
	return uuid.New().String()
}

// checkRegExp: 정규표현식 검사
func checkRegExp(input string, regExp string) (string, string, bool) {
	re, err := regexp.Compile(regExp)
	if err != nil {
		log.Printf("Invalid regular expression: %s", err)
		return input, regExp, false
	}
	result := re.MatchString(input)
	return input, regExp, result
}

// decodeJWT: JWT 디코딩
func DecodeJWT(token string) (map[string]any, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid JWT token")
	}
	// e := base64.StdEncoding.WithPadding(base64.NoPadding)
	// payload, err := e.DecodeString(parts[1])
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 payload: %w", err)
	}

	var result map[string]any
	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON payload: %w", err)
	}
	return result, nil
}

// convertUnixTimeStampToDate: 타임스탬프를 Date로 변환
func ConvertUnixTimeStampToDate(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// validEmail: 이메일 정규표현식 검사
func ValidEmail(input string) (string, string, bool) {
	return checkRegExp(input, `^[^\s@]+@[^\s@]+\.[^\s@]+$`)
}

// Log: 로깅 유틸리티
type Log struct{}

func (l *Log) log(level string, sfmt string, args ...interface{}) {
	colors := map[string]string{
		"success": "\033[32m", // Green
		"error":   "\033[31m", // Red
		"warn":    "\033[33m", // Yellow
		"info":    "\033[34m", // Blue
		"debug":   "\033[90m", // Gray
	}
	now := time.Now().Format("01/02 15:04:05")
	color := colors[level]
	reset := "\033[0m"
	levelFormatted := fmt.Sprintf("%-7s", strings.ToUpper(level))
	headStr := fmt.Sprintf("[%s] %s%s%s %s\n", now, color, levelFormatted, reset, sfmt)
	fmt.Printf(headStr, args...)
}

func (l *Log) Success(fmt string, args ...interface{}) {
	l.log("success", fmt, args...)
}

func (l *Log) Error(fmt string, args ...interface{}) {
	l.log("error", fmt, args...)
}

func (l *Log) Warn(fmt string, args ...interface{}) {
	l.log("warn", fmt, args...)
}

func (l *Log) Info(fmt string, args ...interface{}) {
	l.log("info", fmt, args...)
}

func (l *Log) Debug(fmt string, args ...interface{}) {
	l.log("debug", fmt, args...)
}

/*
func main() {
	// UUID 생성 예제
	fmt.Println("Generated UUID:", createUUID())

	// 정규표현식 검사 예제
	input, regExp, result := validEmail("example@domain.com")
	fmt.Printf("Input: %s, RegExp: %s, Result: %v\n", input, regExp, result)

	// JWT 디코딩 예제
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvbiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.sflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	decoded, err := decodeJWT(jwt)
	if err != nil {
		log.Printf("Failed to decode JWT: %v", err)
	} else {
		fmt.Printf("Decoded JWT: %+v\n", decoded)
	}

	// 타임스탬프 변환 예제
	date := convertUnixTimeStampToDate(1638382800)
	fmt.Println("Converted Date:", date)

	// 로깅 예제
	logger := &Log{}
	logger.Info("This is an info message")
	logger.Warn("This is a warning")
	logger.Error("This is an error")
	logger.Debug("This is debug information")
	logger.Success("This is a success message")
}
*/
