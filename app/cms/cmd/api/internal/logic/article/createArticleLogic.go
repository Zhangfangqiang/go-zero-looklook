// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"
	"looklook/app/cms/cmd/rpc/cms"
	"looklook/pkg/ctxdata"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章
func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (resp *types.CreateArticleResp, err error) {
	// 从context中获取当前登录用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)

	createResp, err := l.svcCtx.CmsRpc.CreateArticle(l.ctx, &cms.CreateArticleReq{
		Title:      req.Title,
		Content:    req.Content,
		Category:   req.Category,
		AuthorId:   userId,
		CoverImage: req.CoverImage,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create article failed, req: %+v", req)
	}

	return &types.CreateArticleResp{
		Id: createResp.Id,
	}, nil
}
