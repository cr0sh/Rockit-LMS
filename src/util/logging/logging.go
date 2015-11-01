package logging

import (
	"fmt"
	"time"
)

var level = 1

func log(lv int, prefix string, f ...interface{}) {
	if lv < level {
		return
	}
	t := time.Now()
	hour, min, sec := t.Clock()
	fmt.Printf("[%02d:%02d:%02d] %s", hour, min, sec, prefix)
	fmt.Println(f...)
}

//Debug prints DEBUG-level log
func Debug(f ...interface{}) {
	log(0, "DEBUG: ", f...)
}

//Verbose prints INFO-level log
func Verbose(f ...interface{}) {
	log(1, " INFO: ", f...)
}

//Warning prints WARN-level log
func Warning(f ...interface{}) {
	log(2, " WARN: ", f...)
}

//FromError prints ERROR-level log from error
func FromError(err error, level int) {
	log(level, "***ERROR***:", err.Error())
}

//Error prints ERROR-level log from given string
func Error(f ...interface{}) {
	log(3, "***ERROR***: ", f...)
}

//SetLevel declares how much data will be printed to console. 0: debugging, 1: verbose, 2: warning, 3: error, 4: nothing(only panic)
func SetLevel(lv int) {
	if lv <= 4 && lv >= 0 {
		switch lv {
		case 0:
			level = lv
			Debug("Set loglevel to DEBUG:", lv)
		case 1:
			Debug("Set loglevel to VERBOSE:", lv)
			level = lv
		case 2:
			Verbose("Set loglevel to WARNING-and-upper:", lv)
			level = lv
		case 3:
			Verbose("Set loglevel to ERROR-only:", lv)
			level = lv
		case 4:
			Verbose("Set loglevel to PANIC:", lv)
			level = lv
		}
	}
}
