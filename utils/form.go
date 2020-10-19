package utils

import (
	"fmt"
	"math"
	"time"
)

const ver string = "0.1.3"

var constamp string = "\x5b\x54\x6f\x72\x63\x68\x2f\x49\x6e\x66\x6f\x5d"
var wrapstamp string = "\x5b\x57\x72\x61\x70\x70\x65\x72\x2f\x47\x6f\x72\x6f\x75\x74\x69\x6e\x65\x5d"

// HMS represents hours minutes and seconds used for formatting time left as well as the string that represents the struct
type HMS struct {
	H   int
	M   int
	S   int
	Msg string
}

// Converts seconds to a hh/mm string
func (hms HMS) secondsToHms(d int) HMS {
	hms.H = int(math.Round(float64(d / 3600)))
	hms.M = int(math.Round(float64(d % 3600 / 60)))
	hms.S = int(math.Round(float64(d % 3600 % 60)))

	hms.Msg = formatHMS(&hms)
	return hms
}

// Take pointer to HMS struct and return formatted string to use as message
func formatHMS(hms *HMS) string {
	switch {
	case hms.M == 0:
		return fmt.Sprintf(" %d seconds", hms.S)
	case hms.H == 0:
		return fmt.Sprintf(" %d minute(s), %d seconds", hms.M, hms.S)
	default:
		return fmt.Sprintf(" %d hour(s), %d minute(s), %d seconds", hms.H, hms.M, hms.S)
	}

}

//Timestamp returns current time
func Timestamp() string {
	t := time.Now()
	return fmt.Sprintf("[%d:%d:%d]", t.Hour(), t.Minute(), t.Second())
}

// Banner returns program banner
func Banner(conf *Config) string {

	var banner string = fmt.Sprintf(`
_|_  _  ._ _ |_  | A simple server restart utility - Ver %v
 |_ (_) | (_ | | | by disastrpc @ https://github.com/disastrpc/torch
-----------------|----------------------------------------------------
[Λ] Java path:   %v
[Λ] Server jar:  %v
[Λ] Arguments:   %v
[Λ] Interval:    %d secs
[Λ] Warn count:  %d
`, ver, conf.JavPath, conf.JarFile, conf.JvmArgs[0:len(conf.JvmArgs)], conf.Interval, conf.WarnCount)

	return banner
}

// FormatWarn returns byte slice containing warning message that will be pushed to stdin
func FormatWarn(T int) []byte {

	var m string

	var hms HMS

	m1 := "\x20\x75\x6e\x74\x69\x6c\x20\x72\x65\x73\x74\x61\x72\x74\x20"
	m2 := "\x20\x73\x65\x72\x76\x65\x72\x20\x69\x73\x20\x72\x65\x73\x74\x61\x72\x74\x69\x6e\x67\x20\x6e\x6f\x77\x20"
	hms = hms.secondsToHms(T)

	switch {
	case (T / 3600) > 1:
		m = fmt.Sprintf("\x73\x61\x79\x20%v%v%v\n", constamp, hms.Msg, m1)
	case (T % 3600 / 60) > 1:
		T = T / 60
		m = fmt.Sprintf("\x73\x61\x79\x20%v%v%v\n", constamp, hms.Msg, m1)
	case T == 0:
		m = fmt.Sprintf("\x73\x61\x79\x20%v%v\n", constamp, m2)
	default:
		m = fmt.Sprintf("\x73\x61\x79\x20%v%v%v\n", constamp, T, m1)
	}

	return []byte(m)
}

// FormatLog Formats message to include Torch stamp and timestamp
func FormatLog(s string) string {
	s = fmt.Sprintf("%v %v %v %v\n", Timestamp(), constamp, wrapstamp, s)
	return s
}

// FormatConsole returns byte slice to push to stdin server console
func FormatConsole(s string) []byte {
	s = fmt.Sprintf("\x73\x61\x79\x20%v %v\n", constamp, s)
	return []byte(s)
}
