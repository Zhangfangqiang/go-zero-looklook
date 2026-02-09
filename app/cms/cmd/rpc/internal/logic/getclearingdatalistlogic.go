package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClearingDataListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClearingDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClearingDataListLogic {
	return &GetClearingDataListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClearingDataListLogic) GetClearingDataList(in *pb.GetClearingDataListReq) (*pb.GetClearingDataListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetClearingDataListResp{}, nil
}
