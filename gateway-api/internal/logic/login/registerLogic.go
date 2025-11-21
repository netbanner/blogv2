// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"blogV2/model"
	"context"
	"database/sql"
	"errors"

	"blogV2/gateway-api/internal/svc"
	"blogV2/gateway-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegisterLogic 注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.ResisterResp, err error) {
	_, err = l.svcCtx.UsersModel.FindByUserName(l.ctx, req.Username)
	if err == nil {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(err, model.ErrNotFound) { // 其他错误
		return nil, err
	}
	// 插入（自动生成缓存）
	user := model.Users{
		UserName: sql.NullString{String: req.Username, Valid: true},
		Password: sql.NullString{String: req.Password, Valid: true}, // 生产环境请使用 bcrypt 加密
		Email:    sql.NullString{String: req.Email, Valid: true},
	}
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, err
	}
	return &types.ResisterResp{Code: 200, Msg: "注册成功"}, nil
}
