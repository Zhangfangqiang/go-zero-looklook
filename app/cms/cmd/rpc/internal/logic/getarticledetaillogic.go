package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"
	"looklook/app/cms/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleDetailLogic {
	return &GetArticleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleDetailLogic) GetArticleDetail(in *pb.GetArticleDetailReq) (*pb.GetArticleDetailResp, error) {
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(err, "article not found, id: %d", in.Id)
		}
		return nil, errors.Wrapf(err, "find article failed, id: %d", in.Id)
	}

	return &pb.GetArticleDetailResp{
		Article: &pb.Article{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			Category:    article.Category,
			LikeCount:   article.LikeCount,
			AuthorId:    article.AuthorId,
			CoverImage:  article.CoverImage,
			Status:      article.Status,
		},
	}, nil
}
