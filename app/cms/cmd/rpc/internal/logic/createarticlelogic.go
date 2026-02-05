package logic

import (
	"context"
	"database/sql"
	"time"

	"looklook/app/cms/cmd/rpc/internal/svc"
	"looklook/app/cms/cmd/rpc/pb"
	"looklook/app/cms/model"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateArticleLogic) CreateArticle(in *pb.CreateArticleReq) (*pb.CreateArticleResp, error) {
	article := &model.Article{
		Title:       in.Title,
		Content:     in.Content,
		Category:    in.Category,
		AuthorId:    in.AuthorId,
		CoverImage:  in.CoverImage,
		PublishTime: time.Now(),
		Status:      0, // 草稿状态
		LikeCount:   0,
		DelState:    0,
		DeleteTime:  sql.NullTime{Valid: false}, // 未删除，DeleteTime为NULL
		Version:     0,
	}

	result, err := l.svcCtx.ArticleModel.Insert(l.ctx, nil, article)
	if err != nil {
		return nil, errors.Wrapf(err, "create article failed, req: %+v", in)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(err, "get last insert id failed")
	}

	return &pb.CreateArticleResp{
		Id: id,
	}, nil
}
