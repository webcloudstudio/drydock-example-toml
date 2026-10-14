package toml

import (
	"reflect"
	"testing"
)

func TestDecodeDateTimeForms(t *testing.T) {
	tests := []struct {
		name  string
		input string
		kind  string
		want  string
	}{
		{name: "local date", input: "1979-05-27", kind: "date-local", want: "1979-05-27"},
		{name: "local time", input: "07:32:00", kind: "time-local", want: "07:32:00"},
		{name: "local time with fraction", input: "00:32:00.999999", kind: "time-local", want: "00:32:00.999999"},
		{name: "local datetime with T", input: "1979-05-27T07:32:00", kind: "datetime-local", want: "1979-05-27T07:32:00"},
		{name: "local datetime with space", input: "1979-05-27 07:32:00", kind: "datetime-local", want: "1979-05-27T07:32:00"},
		{name: "local datetime with lowercase t", input: "1979-05-27t07:32:00", kind: "datetime-local", want: "1979-05-27T07:32:00"},
		{name: "offset datetime Z", input: "1979-05-27T07:32:00Z", kind: "datetime", want: "1979-05-27T07:32:00Z"},
		{name: "offset datetime lowercase z", input: "1979-05-27T00:32:00z", kind: "datetime", want: "1979-05-27T00:32:00Z"},
		{name: "offset datetime negative offset", input: "1979-05-27T00:32:00-07:00", kind: "datetime", want: "1979-05-27T00:32:00-07:00"},
		{name: "offset datetime positive offset", input: "1979-05-27T00:32:00+07:32", kind: "datetime", want: "1979-05-27T00:32:00+07:32"},
		{name: "offset datetime with space", input: "1979-05-27 07:32:00Z", kind: "datetime", want: "1979-05-27T07:32:00Z"},
		{name: "offset datetime with fraction", input: "1979-05-27T00:32:00.999999-07:00", kind: "datetime", want: "1979-05-27T00:32:00.999999-07:00"},
		{name: "fraction keeps fewer digits unpadded", input: "1979-05-27T00:32:00.5Z", kind: "datetime", want: "1979-05-27T00:32:00.5Z"},
		{name: "fraction truncated past nine digits", input: "1979-05-27T00:32:00.1234567899Z", kind: "datetime", want: "1979-05-27T00:32:00.123456789Z"},
		{name: "fraction truncated not rounded", input: "00:00:00.9999999999", kind: "time-local", want: "00:00:00.999999999"},
		{name: "leap second", input: "1990-12-31T23:59:60Z", kind: "datetime", want: "1990-12-31T23:59:60Z"},
		{name: "leap day in leap year", input: "2000-02-29", kind: "date-local", want: "2000-02-29"},
		{name: "first and last of year", input: "0000-01-01", kind: "date-local", want: "0000-01-01"},
		{name: "last day of December", input: "9999-12-31", kind: "date-local", want: "9999-12-31"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("d = " + tt.input + "\n")
			if err != nil {
				t.Fatalf("Decode(%q) error = %v, want nil", tt.input, err)
			}
			want := map[string]any{"d": map[string]any{"type": tt.kind, "value": tt.want}}
			if !reflect.DeepEqual(doc, want) {
				t.Errorf("Decode(%q) = %v, want %v", tt.input, doc, want)
			}
		})
	}
}

func TestDecodeRejectsInvalidDateTimes(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "month zero", input: "1979-00-01"},
		{name: "month thirteen", input: "1979-13-01"},
		{name: "day zero", input: "1979-05-00"},
		{name: "day past month end", input: "1979-04-31"},
		{name: "February 29 in common year", input: "1979-02-29"},
		{name: "February 30 in leap year", input: "2000-02-30"},
		{name: "date-time with hour twenty-four", input: "1979-05-27T24:00:00"},
		{name: "minute sixty", input: "07:60:00"},
		{name: "second sixty-one", input: "07:32:61"},
		{name: "offset hour twenty-four", input: "1979-05-27T07:32:00+24:00"},
		{name: "offset minute sixty", input: "1979-05-27T07:32:00-07:60"},
		{name: "time with hour twenty-four", input: "24:00:00"},
		{name: "date without seconds", input: "1979-05-27T07:32"},
		{name: "time without seconds", input: "07:32"},
		{name: "short year", input: "79-05-27"},
		{name: "offset without time", input: "1979-05-27Z"},
		{name: "offset without minutes", input: "1979-05-27T07:32:00+07"},
		{name: "trailing text after date-time", input: "1979-05-27T07:32:00Zx"},
		{name: "date followed by letters", input: "1979-05-27abc"},
		{name: "fraction with no digits", input: "07:32:00."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := decodeString("d = " + tt.input + "\n")
			if err == nil {
				t.Errorf("Decode(%q) = %v, want an error", tt.input, doc)
			}
		})
	}
}
