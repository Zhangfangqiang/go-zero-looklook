// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearingDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取出清数据列表
func NewClearingDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearingDataListLogic {
	return &ClearingDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearingDataListLogic) ClearingDataList(req *types.ClearingDataListReq) (resp *types.ClearingDataListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
