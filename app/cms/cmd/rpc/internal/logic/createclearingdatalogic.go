package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"
	"looklook/app/cms/model"
	"time"
)

type CreateClearingDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateClearingDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClearingDataLogic {
	return &CreateClearingDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateClearingDataLogic) CreateClearingData(in *pb.CreateClearingDataReq) (*pb.CreateClearingDataResp, error) {

	println("in.ProvinceId:", in.ProvinceId)

	// 将字符串转换为 time.Time
	targetDate, err := time.Parse("2006-01-02", in.TargetDate)
	if err != nil {
		return nil, err
	}

	tPriClearingDataGeneralYunnan := &model.TPriClearingDataGeneralYunnan{
		ProvinceId:             in.ProvinceId,
		CompanyId:              in.CompanyId,
		CompanyName:            in.CompanyName,
		UnitId:                 in.UnitId,
		UnitName:               in.UnitName,
		TargetDate:             targetDate,
		Timeperiod:             in.Timeperiod,
		DayaheadClearingEnergy: in.DayaheadClearingEnergy,
		DayaheadClearingPrice:  in.DayaheadClearingPrice,
		RealtimeClearingEnergy: in.RealtimeClearingEnergy,
		RealtimeClearingPrice:  in.RealtimeClearingPrice,
		DayaheadBidPower:       in.DayaheadBidPower,
		RealtimeBidPower:       in.RealtimeBidPower,
		ActualEnergy:           in.ActualEnergy,
		ActualPower:            in.ActualPower,
	}

	result, err := l.svcCtx.TPriClearingDataGeneralYunnanModel.Insert(l.ctx, nil, tPriClearingDataGeneralYunnan)

	if err != nil {
		return nil, errors.Wrapf(err, "create tPriClearingDataGeneralYunnan failed, req: %+v", in)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(err, "get last insert id failed")
	}

	return &pb.CreateClearingDataResp{
		Id: id,
	}, nil
}
