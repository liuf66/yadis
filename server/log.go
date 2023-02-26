package yadis

import "log"

func ylog(format string, args ...any) {
	log.Printf(format, args...)
}
