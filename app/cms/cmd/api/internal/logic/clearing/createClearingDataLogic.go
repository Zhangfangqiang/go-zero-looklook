// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClearingDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建出清数据
func NewCreateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClearingDataLogic {
	return &CreateClearingDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClearingDataLogic) CreateClearingData(req *types.CreateClearingDataReq) (resp *types.CreateClearingDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
