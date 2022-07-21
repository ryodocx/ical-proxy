package util

// TODO: move to https://github.com/ryodocx/golib

import (
	"fmt"
	"runtime"
)

func WrapError(err error, msg ...string) error {
	pc, f, l, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fmt.Errorf("%s: %+v\n%s@\n%s:%d", msg, err, fn.Name(), f, l)
}
