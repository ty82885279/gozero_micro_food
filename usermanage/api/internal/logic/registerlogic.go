package logic

import (
	"context"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"gozero_micro_food/usermanage/api/internal/svc"
	"gozero_micro_food/usermanage/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (*types.RegisteResponse, error) {
	resp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterRequest{
		Username: req.Username,
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
		Id:       resp.Id,
		Email:    resp.Email,
		JwtToken: token,
	}

	return &types.RegisteResponse{
		UserReply: response,
	}, nil
}
