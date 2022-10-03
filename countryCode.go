package countryCode

import (
	"errors"
	"strconv"
	"strings"
)

func Country(code string) (string, error) {

	country := map[string]string{"TH": "Thailand", "HK": "Hongkong", "JP": "Japan", "CN": "China", "US": "America"}

	if _, err := strconv.ParseFloat(code, 64); err == nil {
		return "", errors.New("error, your input look like a number")
	}

	code2 := strings.ToUpper(code)
	result := country[code2]
	return result, nil

}
