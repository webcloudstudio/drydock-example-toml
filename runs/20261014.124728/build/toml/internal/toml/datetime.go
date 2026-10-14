package toml

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// maxFractionDigits is the fractional-second precision kept. Digits past it are
// truncated, never rounded, and a shorter fraction is never padded.
const maxFractionDigits = 9

// dateTimePattern is the TOML v1.0.0 grammar for a date, optionally followed by a time and
// an offset. The date and time are separated by T, t, or a space. The offset is part of the
// date-time form only, so a bare date or a local date-time stops before one.
var dateTimePattern = regexp.MustCompile(`^(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})` +
	`(?:[Tt ](?P<hour>\d{2}):(?P<minute>\d{2}):(?P<second>\d{2})(?P<fraction>\.\d+)?` +
	`(?P<offset>[Zz]|[+-](?P<offsetHour>\d{2}):(?P<offsetMinute>\d{2}))?)?`)

// timePattern is the TOML v1.0.0 grammar for a time with neither a date nor an offset.
var timePattern = regexp.MustCompile(`^(?P<hour>\d{2}):(?P<minute>\d{2}):(?P<second>\d{2})(?P<fraction>\.\d+)?`)

// temporal reads a date, time, or date-time that starts at p.pos. It reports false and
// consumes nothing when no temporal form starts there.
func (p *parser) temporal() (map[string]any, bool, error) {
	rest := p.src[p.pos:]
	fields, n := namedMatch(dateTimePattern, rest)
	if n == 0 {
		fields, n = namedMatch(timePattern, rest)
	}
	if n == 0 {
		return nil, false, nil
	}
	if next := p.pos + n; next < len(p.src) && (isScalarByte(p.src[next]) || p.src[next] == ':') {
		return nil, true, p.errorf("invalid value %q", rest[:n])
	}
	kind, value, err := temporalForm(fields)
	if err != nil {
		return nil, true, p.errorf("%v", err)
	}
	p.pos += n
	return tagged(kind, value), true, nil
}

// namedMatch matches re at the start of s. It returns the named groups by name and the
// length of the match, or a length of zero when re does not match.
func namedMatch(re *regexp.Regexp, s string) (map[string]string, int) {
	m := re.FindStringSubmatch(s)
	if m == nil {
		return nil, 0
	}
	fields := make(map[string]string, len(m))
	for i, name := range re.SubexpNames() {
		if name != "" {
			fields[name] = m[i]
		}
	}
	return fields, len(m[0])
}

// temporalForm classifies a matched form by the parts it has, checks each range, and returns
// the tag and the canonical text. The date-time delimiter is written as T, and a lowercase Z is
// written as Z. Any other offset is kept as written.
func temporalForm(f map[string]string) (string, string, error) {
	date := ""
	if f["year"] != "" {
		if err := checkDate(f["year"], f["month"], f["day"]); err != nil {
			return "", "", err
		}
		date = f["year"] + "-" + f["month"] + "-" + f["day"]
	}
	if f["hour"] == "" {
		return "date-local", date, nil
	}
	if err := checkClock(f["hour"], f["minute"], f["second"]); err != nil {
		return "", "", err
	}
	clock := f["hour"] + ":" + f["minute"] + ":" + f["second"] + truncateFraction(f["fraction"])
	if date == "" {
		return "time-local", clock, nil
	}
	text := date + "T" + clock
	if f["offset"] == "" {
		return "datetime-local", text, nil
	}
	if err := checkOffset(f["offsetHour"], f["offsetMinute"]); err != nil {
		return "", "", err
	}
	return "datetime", text + canonicalOffset(f["offset"]), nil
}

// checkDate rejects a month outside 01-12 and a day that the month does not have in that year.
func checkDate(year, month, day string) error {
	m, d := number(month), number(day)
	if m < 1 || m > 12 {
		return fmt.Errorf("month %s is out of range", month)
	}
	// time.Date rolls an out-of-range day into the next month, so a changed day means it was invalid.
	if time.Date(number(year), time.Month(m), d, 0, 0, 0, 0, time.UTC).Day() != d {
		return fmt.Errorf("day %s is out of range for %s-%s", day, year, month)
	}
	return nil
}

// checkClock rejects an hour past 23, a minute past 59, or a second past 60. A second of 60
// is the leap second that RFC 3339 allows.
func checkClock(hour, minute, second string) error {
	if number(hour) > 23 {
		return fmt.Errorf("hour %s is out of range", hour)
	}
	if number(minute) > 59 {
		return fmt.Errorf("minute %s is out of range", minute)
	}
	if number(second) > 60 {
		return fmt.Errorf("second %s is out of range", second)
	}
	return nil
}

// checkOffset rejects an offset hour past 23 or an offset minute past 59.
func checkOffset(hour, minute string) error {
	if number(hour) > 23 {
		return fmt.Errorf("offset hour %s is out of range", hour)
	}
	if number(minute) > 59 {
		return fmt.Errorf("offset minute %s is out of range", minute)
	}
	return nil
}

// truncateFraction keeps at most maxFractionDigits digits after the point. The point itself
// is kept, so a fraction is never rounded up into the seconds.
func truncateFraction(fraction string) string {
	if len(fraction) > 1+maxFractionDigits {
		return fraction[:1+maxFractionDigits]
	}
	return fraction
}

// canonicalOffset writes a lowercase z as Z and leaves a numeric offset as written.
func canonicalOffset(offset string) string {
	if offset == "z" {
		return "Z"
	}
	return offset
}

// number converts digits that a temporal pattern has already matched. The conversion
// cannot fail, so its error is discarded.
func number(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
