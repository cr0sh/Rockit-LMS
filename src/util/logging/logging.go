package logging

import (
	"fmt"
	"time"
)

//Level declares how much data will be printed to console. 0: debugging, 1: verbose, 2: warning, 3: error, 4: nothing(only panic)
var Level int

func log(level int, prefix string, f ...interface{}) {
	if level < Level {
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
	log(1, "INFO: ", f...)
}

//Warning prints WARN-level log
func Warning(f ...interface{}) {
	log(2, "WARN: ", f...)
}

//FromError prints ERROR-level log from error
func FromError(err error, level int) {
	log(level, "ERROR:", err.Error())
}

//Error prints ERROR-level log from given string
func Error(f ...interface{}) {
	log(3, "ERROR: ", f...)
}
