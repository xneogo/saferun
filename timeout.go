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
 @Time    : 2025/7/2 -- 12:19
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: saferun /timeout.go
*/

package saferun

import (
	"context"
	"errors"
	"time"
)

func TimeoutRun(ctx context.Context, fn RunFn, ttl time.Duration) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		err = SafeGo(ctx, fn)
		cancel()
	}()

	select {
	case <-ctx.Done():
		return
	case <-time.After(ttl):
		return errors.New("out of time")
	}
}
