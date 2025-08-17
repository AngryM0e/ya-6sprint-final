package service

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(content []byte) (string, error) {
	if len(content) == 0 {
		return "", errors.New("empty content")
	}

	strContent := toUTF8(string(content))
	strContent = strings.TrimSpace(strContent)
	if strContent == "" {
		return "", errors.New("empty content")
	}

	if isMorseCode(strContent) {
		return morse.DefaultConverter.ToText(strContent), nil
	}

	return morse.DefaultConverter.ToMorse(strContent), nil
}

func toUTF8(s string) string {
	if utf8.ValidString(s) {
		return s
	}

	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		v = append(v, r)
	}
	return string(v)
}

func isMorseCode(s string) bool {
	allowedChars := ".-/ "

	for _, r := range s {
		if !strings.ContainsRune(allowedChars, r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
