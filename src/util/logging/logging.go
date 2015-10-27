package logging

import (
	"fmt"
	"time"
)

//Level declares how much data will be printed to console. 0: debugging, 1: verbose, 2: warning, 3: error, 4: nothing(only panic)
var Level int

func log(str string, level int) {
	if level < Level {
		return
	}
	t := time.Now()
	hour, min, sec := t.Clock()
	fmt.Printf("[%02d:%02d:%02d] %s\n", hour, min, sec, str)
}

//Debug prints DEBUG-level log
func Debug(str string) {
	log("DEBUG: "+str, 0)
}

//Verbose prints INFO-level log
func Verbose(str string) {
	log("INFO: "+str, 1)
}

//Warning prints WARN-level log
func Warning(str string) {
	log("WARN: "+str, 2)
}

//FromError prints ERROR-level log from error
func FromError(err error, level int) {
	log("ERROR: "+err.Error(), level)
}

//Error prints ERROR-level log from given string
func Error(str string) {
	log("ERROR: "+str, 3)
}
