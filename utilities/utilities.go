package utilities

import (
	"fmt"
	"reflect"
	"runtime"
)

// NewFuncInfo will return function information
func NewFuncInfo(fn interface{}) (fi FuncInfo) {
	rt := runtime.FuncForPC(reflect.ValueOf(fn).Pointer())
	fi.File, fi.Line = rt.FileLine(rt.Entry())
	fi.Name = rt.Name()
	return
}

// FuncInfo contains basic func info
type FuncInfo struct {
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func (fi *FuncInfo) String() string {
	return fmt.Sprintf("%s %s (%d)", fi.Name, fi.File, fi.Line)
}
