package xpretty

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func DLog(skip int, cfn func(a ...interface{}) string, msg ...interface{}) {
	if !ctrl.dummyLog {
		return
	}
	c, l := Caller(skip)
	if len(msg) == 0 {
		fmt.Printf("%v %v(%v) %v\n", Faint(nowAsStr()), cfn(c), cfn(l), Info("dummy log"))
	} else {
		s := anyJoin(", ", msg...)
		fmt.Printf("%v %v(%v) %v\n", Faint(nowAsStr()), cfn(c), cfn(l), Info(s))
	}
}

// DummyLog will print a dummy log with green bg
func DummyLog(msg ...interface{}) {
	DLog(3, Green, msg...)
}

// DummyErrorLog will print a dummy log with red bg
func DummyErrorLog(msg ...interface{}) {
	DLog(3, Red, msg...)
}

// Caller wraps runtime.Caller and returns file and line number information
//
// skip:
//   - 0: Caller
//   - 1: is where Caller is called
//
// Returns: filename, linenum
func Caller(skip int) (string, int) {
	pc, file, l, _ := runtime.Caller(skip)

	lst := strings.Split(file, "/")
	file = lst[len(lst)-1]

	funcName := runtime.FuncForPC(pc).Name()
	lst = strings.Split(funcName, "/")
	funcName = fmt.Sprintf("%s:%s", file, lst[len(lst)-1])

	return funcName, l
}

// anyJoin join all args with sep, even value in args is zero
func anyJoin(sep string, args ...interface{}) string {
	var arr []string
	for _, v := range args {
		arr = append(arr, fmt.Sprintf("%v", v))
	}
	return strings.Join(arr, sep)
}

func nowAsStr() string {
	return time.Now().Format("15:04:05")
}
