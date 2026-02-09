package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClearingDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClearingDataLogic {
	return &DeleteClearingDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteClearingDataLogic) DeleteClearingData(in *pb.DeleteClearingDataReq) (*pb.DeleteClearingDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteClearingDataResp{}, nil
}
