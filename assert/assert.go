package assert

import (
	"fmt"
	"log"
)

func True(b bool, s string) {
	if b != true {
		log.Panicf("should be true: %s", s)
	}
}

func Truef(b bool, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if b != true {
		log.Panicf("should be true: %s", msg)
	}
}
