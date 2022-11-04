package shorten

import (
	"net/url"
	"strings"

	"github.com/defer-panic/url-shortener-api/internal/utils"
)

const alphabet = "ynAJfoSgdXHB5VasEMtcbPCr1uNZ4LG723ehWkvwYR6KpxjTm8iQUFqz9D"

var alphabetLen = uint32(len(alphabet))

func Shorten(id uint32, src string) string {
	var (
		digits  []uint32
		num     = id
		builder strings.Builder
	)

	for num > 0 {
		digits = append(digits, num%alphabetLen)
		num /= alphabetLen
	}

	utils.Reverse(digits)

	for _, digit := range digits {
		builder.WriteString(string(alphabet[digit]))
	}

	return builder.String()
}

func PrependBaseURL(baseURL, identifier string) (string, error) {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	parsed.Path = identifier

	return parsed.String(), nil
}
