// Package utils
// @Time:2023/08/24 06:36
// @File:jsRun.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package utils

import (
	"fmt"
	js "github.com/dop251/goja"
)

type JsUtil struct{}

func (j *JsUtil) JsRun(funcContent []string, params ...any) any {
	vm := js.New()
	_, err := vm.RunString(funcContent[0])
	if err != nil {
		return err
	}
	jsfn, ok := js.AssertFunction(vm.Get(funcContent[1]))
	if !ok {
		return fmt.Errorf("执行函数失败")
	}
	jsValues := make([]js.Value, 0, len(params))
	for _, v := range params {
		jsValues = append(jsValues, vm.ToValue(v))
	}
	result, err := jsfn(
		js.Undefined(),
		jsValues...,
	)
	if err != nil {
		return err
	}
	return result
}
