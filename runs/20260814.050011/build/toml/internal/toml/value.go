package toml

import (
	"encoding/json"
	"fmt"
)

// Kind identifies the TOML type represented by a Value.
type Kind string

const (
	String        Kind = "string"
	Integer       Kind = "integer"
	Float         Kind = "float"
	Bool          Kind = "bool"
	Datetime      Kind = "datetime"
	DatetimeLocal Kind = "datetime-local"
	DateLocal     Kind = "date-local"
	TimeLocal     Kind = "time-local"
)

// Value is a scalar TOML value. Lexeme is retained as supplied by the parser
// so integer values never pass through floating-point representation.
type Value struct {
	Kind   Kind
	Lexeme string
}

// Object is the internal representation of a TOML table or array.
type Object map[string]any

// Tagged returns the JSON-compatible tagged representation of v.
func Tagged(v any) (any, error) {
	switch value := v.(type) {
	case Value:
		return struct {
			Type  Kind   `json:"type"`
			Value string `json:"value"`
		}{value.Kind, value.Lexeme}, nil
	case Object:
		result := make(map[string]any, len(value))
		for key, item := range value {
			encoded, err := Tagged(item)
			if err != nil {
				return nil, fmt.Errorf("encode %q: %w", key, err)
			}
			result[key] = encoded
		}
		return result, nil
	case []any:
		result := make([]any, len(value))
		for i, item := range value {
			encoded, err := Tagged(item)
			if err != nil {
				return nil, fmt.Errorf("encode array item %d: %w", i, err)
			}
			result[i] = encoded
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported TOML value %T", v)
	}
}

// MarshalJSON encodes a Value as the required tagged JSON scalar.
func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type  Kind   `json:"type"`
		Value string `json:"value"`
	}{v.Kind, v.Lexeme})
}
