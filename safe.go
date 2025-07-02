/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/4/15 -- 14:40
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: saferun /safe.go 给运行的函数f封装，避免panic导致全局退出
*/

package saferun

import (
	"context"
	"fmt"
	"runtime"
)

type RunFn func() error
type RunArgsFn func(args ...interface{}) error

// SafeGoWithArgs 带参数函数
// f 内部需要对 [interface] 做反解才可以使用
func SafeGoWithArgs(f RunArgsFn, args ...interface{}) (err error) {

	defer func() {
		if p := recover(); p != nil {
			err = DumpStack(p)
		}
	}()

	return f(args...)
}

// SafeGo 无参数
func SafeGo(ctx context.Context, f RunFn) (err error) {

	defer func() {
		if p := recover(); p != nil {
			err = DumpStack(p)
		}
	}()

	return f()
}

func DumpStack(e interface{}) (err error) {
	if e == nil {
		return
	}
	err = fmt.Errorf("%+v", e)
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		err = fmt.Errorf("%s\t %s:%d", err.Error(), file, line)
	}
	return
}
