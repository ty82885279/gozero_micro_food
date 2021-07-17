package logic

import (
	"context"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"gozero_micro_food/usermanage/api/internal/svc"
	"gozero_micro_food/usermanage/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {
	resp, err := l.svcCtx.User.Login(l.ctx, &user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token := types.JwtToken{
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}
	response := types.UserReply{
		Email:    resp.Email,
		Id:       resp.Id,
		JwtToken: token,
	}
	return &types.LoginResponse{
		UserReply: response,
	}, nil
}
