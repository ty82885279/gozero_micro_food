package logic

import (
	"context"
	"errors"
	"gozero_micro_food/lee"
	"gozero_micro_food/usermanage/model"
	"gozero_micro_food/usermanage/rpc/user/internal/svc"
	"gozero_micro_food/usermanage/rpc/user/pro/user"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.Response, error) {
	res, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if err == nil {
		password := lee.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
		if password == res.Password {
			now := time.Now().Unix()
			accessExpire := l.svcCtx.Config.AccessExpire
			token, err := lee.GetJwtToken(l.svcCtx.Config.AccessSecret, now, accessExpire, res.Id)
			if err != nil {
				return nil, err
			}
			response := user.Response{
				Email:        res.Email,
				Id:           res.Id,
				AccessExpire: now + accessExpire,
				AccessToken:  token,
				RefreshAfter: now + accessExpire/2,
			}
			return &response, nil
		} else {
			return nil, errors.New("密码错误")
		}
	}
	if err == model.ErrNotFound {

		return nil, errors.New("用户不存在")
	}
	return nil, err
}
