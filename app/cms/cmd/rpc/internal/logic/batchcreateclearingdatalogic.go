package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateClearingDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateClearingDataLogic {
	return &BatchCreateClearingDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchCreateClearingDataLogic) BatchCreateClearingData(in *pb.BatchCreateClearingDataReq) (*pb.BatchCreateClearingDataResp, error) {
	// todo: add your logic here and delete this line

	return &pb.BatchCreateClearingDataResp{}, nil
}
