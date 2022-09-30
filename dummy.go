package xpretty

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func dlog(cfn func(a ...interface{}) string, msg ...interface{}) {
	if !ctrl.dummyLog {
		return
	}
	c, l := Caller(3)
	if len(msg) == 0 {
		fmt.Printf("%v %v(%v) %v\n", Faint(NowAsStr()), cfn(c), cfn(l), Info("dummy log"))
	} else {
		s := AnyJoin(", ", msg...)
		fmt.Printf("%v %v(%v) %v\n", Faint(NowAsStr()), cfn(c), cfn(l), Info(s))
	}
}

// DummyLog will print a dummy log with green bg
func DummyLog(msg ...interface{}) {
	dlog(Green, msg...)
}

// DummyErrorLog will print a dummy log with red bg
func DummyErrorLog(msg ...interface{}) {
	dlog(Red, msg...)
}

// Caller wraps runtime.Caller and returns file and line number information
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

// AnyJoin:
// join all args with sep, even value in args is zero
// if you don't want zero values, use `AnyJoinNon0` instead
func AnyJoin(sep string, args ...interface{}) string {
	arr := []string{}
	for _, v := range args {
		arr = append(arr, fmt.Sprintf("%v", v))
	}
	return strings.Join(arr, sep)
}

func NowAsStr() string {
	return time.Now().Format("15:04:05")
}
