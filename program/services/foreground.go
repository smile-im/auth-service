package services

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/micro-kit/micro-common/cache"
	"github.com/micro-kit/micro-common/crypto"
	"github.com/micro-kit/micro-common/logger"
	"github.com/micro-kit/micro-common/microerror"
	"github.com/smile-im/auth-service/program/models"
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
	if req.Username == "" || req.Password == "" {
		return nil, microerror.GetMicroError(10001)
	}
	// 逻辑代码
	user := new(models.User)
	password := crypto.PasswordHash(req.Password)
	err := user.UserLogin(req.Username, password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = microerror.GetMicroError(10001)
		} else {
			logger.Logger.Errorw("登录查询用户错误", "err", err, "req", req)
		}
		return nil, err
	}
	// 将登录信息缓存入redis
	key, token := cache.GetUserLoginToken()
	userJs, _ := json.Marshal(user)
	err = cache.GetClient().Set(key, string(userJs), 0).Err()
	if err != nil {
		err = microerror.GetMicroError(10000, err)
		return nil, err
	}

	// 返回结果
	reply := &authpb.LoginReply{
		UserInfo: user.ToPb(),
		Token:    token,
	}
	return reply, nil
}
