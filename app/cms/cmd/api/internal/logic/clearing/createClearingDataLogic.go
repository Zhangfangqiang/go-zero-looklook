// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package clearing

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"
	"looklook/app/cms/cmd/rpc/cms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClearingDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建出清数据
func NewCreateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClearingDataLogic {
	return &CreateClearingDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClearingDataLogic) CreateClearingData(req *types.CreateClearingDataReq) (resp *types.CreateClearingDataResp, err error) {
	createResp, err := l.svcCtx.CmsRpc.CreateClearingData(l.ctx, &cms.CreateClearingDataReq{
		ProvinceId:             req.ProvinceId,
		CompanyId:              req.CompanyId,
		CompanyName:            req.CompanyName,
		UnitId:                 req.UnitId,
		UnitName:               req.UnitName,
		TargetDate:             req.TargetDate,
		Timeperiod:             req.Timeperiod,
		DayaheadClearingEnergy: req.DayaheadClearingEnergy,
		DayaheadClearingPrice:  req.DayaheadClearingPrice,
		RealtimeClearingEnergy: req.RealtimeClearingEnergy,
		RealtimeClearingPrice:  req.RealtimeClearingPrice,
		DayaheadBidPower:       req.DayaheadBidPower,
		RealtimeBidPower:       req.RealtimeBidPower,
		ActualEnergy:           req.ActualEnergy,
		ActualPower:            req.ActualPower,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "create article failed, req: %+v", req)
	}

	return &types.CreateClearingDataResp{
		Id: createResp.Id,
	}, nil
}
