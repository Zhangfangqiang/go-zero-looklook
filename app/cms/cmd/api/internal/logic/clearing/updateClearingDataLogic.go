// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClearingDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新出清数据
func NewUpdateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClearingDataLogic {
	return &UpdateClearingDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClearingDataLogic) UpdateClearingData(req *types.UpdateClearingDataReq) (resp *types.UpdateClearingDataResp, err error) {
	// todo: add your logic here and delete this line

	return
}
