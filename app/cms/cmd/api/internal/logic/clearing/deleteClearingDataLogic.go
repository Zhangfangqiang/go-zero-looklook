// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClearingDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除出清数据
func NewDeleteClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClearingDataLogic {
	return &DeleteClearingDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClearingDataLogic) DeleteClearingData(req *types.DeleteClearingDataReq) (resp *types.DeleteClearingDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
