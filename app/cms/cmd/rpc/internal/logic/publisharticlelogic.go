package logic

import (
	"context"
	"time"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"
	"looklook/app/cms/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishArticleLogic {
	return &PublishArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishArticleLogic) PublishArticle(in *pb.PublishArticleReq) (*pb.PublishArticleResp, error) {
	article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(err, "article not found, id: %d", in.Id)
		}
		return nil, errors.Wrapf(err, "find article failed, id: %d", in.Id)
	}

	// 检查是否是文章作者
	if article.AuthorId != in.AuthorId {
		return nil, errors.New("no permission to publish this article")
	}

	// 更新文章状态为已发布
	article.Status = 1
	article.PublishTime = time.Now()

	err = l.svcCtx.ArticleModel.Update(l.ctx, nil, article)
	if err != nil {
		return nil, errors.Wrapf(err, "update article failed, id: %d", in.Id)
	}

	return &pb.PublishArticleResp{
		Success: true,
	}, nil
}
