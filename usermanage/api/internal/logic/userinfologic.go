package logic

import (
	"context"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"gozero_micro_food/usermanage/api/internal/svc"
	"gozero_micro_food/usermanage/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req types.UserinfoRequest) (*types.UserinfoResponse, error) {

	logx.Infof("11-----userId: %v", l.ctx.Value("userId"))
	res, err := l.svcCtx.User.Userinfo(l.ctx, &user.UserinfoRequest{
		Userid: req.Userid,
	})
	if err != nil {
		return nil, err
	}
	reply := types.UserReply{
		Id:    res.Id,
		Email: res.Email,
	}
	return &types.UserinfoResponse{
		UserReply: reply,
	}, nil
}
