package logic

import (
	"context"
	"errors"
	"gozero_micro_food/lee"
	"gozero_micro_food/usermanage/model"
	"time"

	"gozero_micro_food/usermanage/rpc/user/internal/svc"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.Response, error) {
	_, err := l.svcCtx.Model.FindOneByEmail(in.Email)
	if err == model.ErrNotFound {
		// 可以注册
		newuser := model.User{
			Email:    in.Email,
			Name:     in.Username,
			Password: lee.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}
		if result, err := l.svcCtx.Model.Insert(newuser); err != nil {
			return nil, err
		} else {
			userid, err := result.LastInsertId()
			if err != nil {
				return nil, err
			}
			now := time.Now().Unix()
			accessExpire := now + l.svcCtx.Config.AccessExpire
			token, err := lee.GetJwtToken(l.svcCtx.Config.AccessSecret, now, accessExpire, userid)
			if err != nil {
				return nil, err
			}
			response := user.Response{
				Id:           userid,
				Email:        newuser.Email,
				AccessToken:  token,
				AccessExpire: now + accessExpire,
				RefreshAfter: now + accessExpire/2,
			}
			return &response, nil
		}

	} else if err == nil {
		return nil, errors.New("用户已存在")
	} else {
		return nil, err
	}
}
