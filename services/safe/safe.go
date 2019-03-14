package safe

import (
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

// GoRoutine safe go routine with recover
func GoRoutine(f func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				stack := debug.Stack()
				logrus.Error(string(stack))
			}
		}()

		f()
	}()
}
