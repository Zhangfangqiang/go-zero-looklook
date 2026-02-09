// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearingDataDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取出清数据详情
func NewClearingDataDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearingDataDetailLogic {
	return &ClearingDataDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearingDataDetailLogic) ClearingDataDetail(req *types.ClearingDataDetailReq) (resp *types.ClearingDataDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
