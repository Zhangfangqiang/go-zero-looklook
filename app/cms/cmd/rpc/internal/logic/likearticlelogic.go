package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"
	"looklook/app/cms/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikeArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeArticleLogic {
	return &LikeArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeArticleLogic) LikeArticle(in *pb.LikeArticleReq) (*pb.LikeArticleResp, error) {
	// 先查询文章是否存在
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(err, "article not found, id: %d", in.Id)
		}
		return nil, errors.Wrapf(err, "find article failed, id: %d", in.Id)
	}

	// 更新点赞数
	err = l.svcCtx.ArticleModel.UpdateLikeCount(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "update like count failed, id: %d", in.Id)
	}

	return &pb.LikeArticleResp{
		LikeCount: article.LikeCount + 1,
	}, nil
}
