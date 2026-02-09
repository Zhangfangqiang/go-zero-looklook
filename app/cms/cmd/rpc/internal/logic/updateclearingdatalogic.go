package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClearingDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClearingDataLogic {
	return &UpdateClearingDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateClearingDataLogic) UpdateClearingData(in *pb.UpdateClearingDataReq) (*pb.UpdateClearingDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateClearingDataResp{}, nil
}
