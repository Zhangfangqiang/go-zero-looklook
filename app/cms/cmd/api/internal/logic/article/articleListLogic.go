// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"
	"looklook/app/cms/cmd/rpc/cms"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleListLogic {
	return &ArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleListLogic) ArticleList(req *types.ArticleListReq) (resp *types.ArticleListResp, err error) {
	listResp, err := l.svcCtx.CmsRpc.GetArticleList(l.ctx, &cms.GetArticleListReq{
		Category: req.Category,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "get article list failed, req: %+v", req)
	}

	var list []types.Article
	if len(listResp.List) > 0 {
		for _, item := range listResp.List {
			var article types.Article
			_ = copier.Copy(&article, item)
			list = append(list, article)
		}
	}

	return &types.ArticleListResp{
		List:  list,
		Total: listResp.Total,
	}, nil
}
