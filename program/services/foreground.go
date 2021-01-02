package services

import (
	"context"

	"github.com/micro-kit/micro-common/microerror"
	"github.com/smile-im/microkit-client/proto/authpb"
)

/* 提供给客户端使用的rpc */

// Foreground 实现grpc客户端rpc接口
type Foreground struct {
	Base
}

// NewForeground 创建客户端rpc对象
func NewForeground() *Foreground {
	return &Foreground{
		Base: NewBase(),
	}
}

// Login 登录
func (s *Foreground) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginReply, error) {
	// 验证参数是否错误
	return nil, microerror.GetMicroError(10001)
	// TODO 逻辑代码

	// 返回结果
	reply := &authpb.LoginReply{}
	return reply, nil
}
