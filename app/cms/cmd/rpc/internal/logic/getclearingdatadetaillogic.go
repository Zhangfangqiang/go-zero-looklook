package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClearingDataDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClearingDataDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClearingDataDetailLogic {
	return &GetClearingDataDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClearingDataDetailLogic) GetClearingDataDetail(in *pb.GetClearingDataDetailReq) (*pb.GetClearingDataDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetClearingDataDetailResp{}, nil
}
