package err

import (
	"github.com/micro-kit/micro-common/microerror"
)

/* 当前服务错误码 */

var (
	errs = []*microerror.MicroError{
		{
			Code: 11001,
			Msg:  "用户名或密码错误",
		},
	}
)

func init() {
	microerror.InitError(errs...)
}
