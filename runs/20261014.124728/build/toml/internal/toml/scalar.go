package toml

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Each pattern is the TOML v1.0.0 grammar for one form. Underscores sit only between
// digits, and leading zeros are allowed only in the fraction and exponent of a float.
var (
	decimalIntPattern   = regexp.MustCompile(`^[+-]?(0|[1-9](_?[0-9])*)$`)
	hexIntPattern       = regexp.MustCompile(`^0x[0-9A-Fa-f](_?[0-9A-Fa-f])*$`)
	octalIntPattern     = regexp.MustCompile(`^0o[0-7](_?[0-7])*$`)
	binaryIntPattern    = regexp.MustCompile(`^0b[01](_?[01])*$`)
	floatPattern        = regexp.MustCompile(`^[+-]?(0|[1-9](_?[0-9])*)((\.[0-9](_?[0-9])*)([eE][+-]?[0-9](_?[0-9])*)?|[eE][+-]?[0-9](_?[0-9])*)$`)
	specialFloatPattern = regexp.MustCompile(`^[+-]?(inf|nan)$`)
)

// scalar reads a date, time, or date-time when one starts at p.pos. Otherwise it reads the
// bare token at p.pos and classifies it as a boolean, integer, or float.
// The whole token must match one grammar, so a partly valid token is rejected, not truncated.
func (p *parser) scalar() (map[string]any, error) {
	if tv, ok, err := p.temporal(); ok {
		return tv, err
	}
	start := p.pos
	for !p.done() && isScalarByte(p.src[p.pos]) {
		p.pos++
	}
	token := p.src[start:p.pos]
	switch {
	case token == "":
		return nil, p.errorf("expected a value")
	case token == "true" || token == "false":
		return tagged("bool", token), nil
	case specialFloatPattern.MatchString(token):
		return tagged("float", formatFloat(parseSpecialFloat(token))), nil
	case isInteger(token):
		n, err := parseInteger(token)
		if err != nil {
			return nil, p.errorf("%v", err)
		}
		return tagged("integer", strconv.FormatInt(n, 10)), nil
	case floatPattern.MatchString(token):
		f, err := strconv.ParseFloat(stripUnderscores(token), 64)
		if err != nil {
			return nil, p.errorf("float %s is outside the binary64 range", token)
		}
		return tagged("float", formatFloat(f)), nil
	}
	return nil, p.errorf("invalid value %q", token)
}

// isScalarByte reports whether c can appear in a bare boolean, integer, or float token.
func isScalarByte(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
		c == '+' || c == '-' || c == '_' || c == '.'
}

func isInteger(token string) bool {
	return decimalIntPattern.MatchString(token) || hexIntPattern.MatchString(token) ||
		octalIntPattern.MatchString(token) || binaryIntPattern.MatchString(token)
}

// parseInteger converts a token that passed isInteger to a signed 64-bit value. A value
// outside that range is an error, never a rounded or wrapped result.
func parseInteger(token string) (int64, error) {
	digits := stripUnderscores(token)
	base := 10
	switch {
	case hexIntPattern.MatchString(token):
		base, digits = 16, digits[2:]
	case octalIntPattern.MatchString(token):
		base, digits = 8, digits[2:]
	case binaryIntPattern.MatchString(token):
		base, digits = 2, digits[2:]
	}
	n, err := strconv.ParseInt(digits, base, 64)
	if err != nil {
		return 0, fmt.Errorf("integer %s is outside the signed 64-bit range", token)
	}
	return n, nil
}

// parseSpecialFloat maps a token that passed specialFloatPattern to its IEEE 754 value.
func parseSpecialFloat(token string) float64 {
	switch {
	case strings.HasSuffix(token, "nan"):
		return math.NaN()
	case strings.HasPrefix(token, "-"):
		return math.Inf(-1)
	}
	return math.Inf(1)
}

// formatFloat renders a float as the shortest decimal that reads back to the same
// binary64 value. Special values use the TOML spelling without a leading plus.
func formatFloat(f float64) string {
	switch {
	case math.IsNaN(f):
		return "nan"
	case math.IsInf(f, 1):
		return "inf"
	case math.IsInf(f, -1):
		return "-inf"
	}
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func stripUnderscores(token string) string {
	return strings.ReplaceAll(token, "_", "")
}

// tagged builds the {"type", "value"} model for a scalar whose value is already a string.
func tagged(kind, value string) map[string]any {
	return map[string]any{"type": kind, "value": value}
}
