package utils

import (
	"fmt"
	"time"
)

var ver string = "1.0.0"

var constamp string = "\x5b\x54\x6f\x72\x63\x68\x2f\x49\x6e\x66\x6f\x5d"
var wrapstamp string = "\x5b\x57\x72\x61\x70\x70\x65\x72\x2f\x47\x6f\x72\x6f\x75\x74\x69\x6e\x65\x5d"

func Timestamp() string {
	t := time.Now()
	return fmt.Sprintf("[%d:%d:%d]", t.Hour(), t.Minute(), t.Second())
}

func Banner(conf *Config) string {

	// Banner .
	var banner string = fmt.Sprintf(`
_|_  _  ._ _ |_  | A simple server restart utility - Ver %v
 |_ (_) | (_ | | | by disastrpc @ https://github.com/disastrpc/torch
-----------------|----------------------------------------------------
[Λ] Java path	: %v
[Λ] Server jar	: %v
[Λ] Arguments	: %v
[Λ] Interval	: %d secs
`, ver, conf.JavPath, conf.JarFile, conf.JvmArgs[0:len(conf.JvmArgs)], conf.Interval)

	return banner
}

// FormatWarn returns byte slice containing warning message that will be pushed to stdin
func FormatWarn(T int) []byte {

	var s string

	switch {
	case T >= 60:
		T = T / 60
		s = "minute(s)"
	default:
		s = "seconds"
	}

	m := fmt.Sprintf("\x73\x61\x79\x20%v %v %v until restart\n", constamp, T, s)

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
