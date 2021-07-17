package logic

import (
	"context"
	"strconv"

	"gozero_micro_food/usermanage/rpc/user/internal/svc"
	"gozero_micro_food/usermanage/rpc/user/pro/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoRequest) (*user.Response, error) {
	userid, _ := strconv.ParseInt(in.Userid, 10, 64)
	result, err := l.svcCtx.Model.FindOne(userid)
	if err == nil {
		return &user.Response{
			Id:    result.Id,
			Email: result.Email,
		}, nil
	} else {
		return nil, err
	}

}
