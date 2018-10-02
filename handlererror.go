package main

import (
	"fmt"
	"log"
	"runtime"
)

func HandlerError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(3)
		log.Printf("\n[error] %s:%d %v", fn, line, err)
		pc := make([]uintptr, 10) // at least 1 entry needed
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		fmt.Printf("\n[error] %s:%d %s", file, line, f.Name())
		b = true
	}
	return
}
