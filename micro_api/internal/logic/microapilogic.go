package logic

import (
	"context"

	"micro/micro_api/internal/svc"
	"micro/micro_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Micro_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMicro_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Micro_apiLogic {
	return &Micro_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Micro_apiLogic) Micro_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
