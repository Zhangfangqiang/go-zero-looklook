// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateClearingDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量创建出清数据
func NewBatchCreateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateClearingDataLogic {
	return &BatchCreateClearingDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchCreateClearingDataLogic) BatchCreateClearingData(req *types.BatchCreateClearingDataReq) (resp *types.BatchCreateClearingDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
