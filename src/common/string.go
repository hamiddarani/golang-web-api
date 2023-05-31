package common

import (
	"golang-web-api/config"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func GenerateOtp() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
