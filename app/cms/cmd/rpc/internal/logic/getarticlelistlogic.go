package logic

import (
	"context"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleListLogic) GetArticleList(in *pb.GetArticleListReq) (*pb.GetArticleListResp, error) {
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	// 查询已发布的文章
	articles, total, err := l.svcCtx.ArticleModel.FindPageListByPage(l.ctx, page, pageSize, in.Category, 1)
	if err != nil {
		return nil, errors.Wrapf(err, "find article list failed, req: %+v", in)
	}

	var list []*pb.Article
	for _, article := range articles {
		list = append(list, &pb.Article{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			PublishTime: article.PublishTime.Format("2006-01-02 15:04:05"),
			Category:    article.Category,
			LikeCount:   article.LikeCount,
			AuthorId:    article.AuthorId,
			CoverImage:  article.CoverImage,
			Status:      article.Status,
		})
	}

	return &pb.GetArticleListResp{
		List:  list,
		Total: total,
	}, nil
}
