package utils

import (
	"fmt"
	"time"
)

var Ver string = "1.0.0"

var Constamp string = "\x5b\x54\x6f\x72\x63\x68\x2f\x57\x72\x61\x70\x70\x65\x72\x5d"

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
`, Ver, conf.JavPath, conf.JarFile, conf.JvmArgs[0:len(conf.JvmArgs)], conf.Interval)

	return banner
}
