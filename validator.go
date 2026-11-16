// Package validator provides deterministic normalization and validation helpers
// for multi-country 15-column B2B company Excel (.xlsx) and UTF-8 CSV datasets.
package validator

import (
	"regexp"
	"strings"
)

var (
	naicsRe  = regexp.MustCompile(`^[0-9]{2,6}$`)
	kvkRe    = regexp.MustCompile(`^[0-9]{8}$`)
	e164Re   = regexp.MustCompile(`^\+[1-9][0-9]{6,14}$`)
	postalRe = regexp.MustCompile(`^[A-Za-z0-9 -]{3,12}$`)
)

// NormalizePostalZeroPad ensures national numeric ZIP/postal codes keep leading zeros
// for US East Coast (07102, 02110), Germany, Spain, and France.
func NormalizePostalZeroPad(raw string, width int) string {
	s := strings.TrimSpace(raw)
	if s == "" || width <= 0 || len(s) >= width {
		return s
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return s
		}
	}
	return strings.Repeat("0", width-len(s)) + s
}

// IsValidNAICS checks 2-to-6-digit U.S. Census Bureau NAICS industrial codes
// (e.g. NAICS 31-33 Manufacturing and NAICS 23 Construction).
func IsValidNAICS(code string) bool {
	return naicsRe.MatchString(strings.TrimSpace(code))
}

// IsValidKVK checks 8-digit Netherlands Kamer van Koophandel registry numbers.
func IsValidKVK(kvk string) bool {
	return kvkRe.MatchString(strings.TrimSpace(kvk))
}

// IsValidE164 checks strict E.164 international B2B phone formatting.
func IsValidE164(phone string) bool {
	return e164Re.MatchString(strings.TrimSpace(phone))
}

// IsValidPostal checks printable international postal codes after zero-padding.
func IsValidPostal(postal string) bool {
	return postalRe.MatchString(strings.TrimSpace(postal))
}
