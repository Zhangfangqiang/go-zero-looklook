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

type PublishArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发布文章
func NewPublishArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishArticleLogic {
	return &PublishArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishArticleLogic) PublishArticle(req *types.PublishArticleReq) (resp *types.PublishArticleResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	publishResp, err := l.svcCtx.CmsRpc.PublishArticle(l.ctx, &cms.PublishArticleReq{
		Id:       req.Id,
		AuthorId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "publish article failed, req: %+v", req)
	}

	return &types.PublishArticleResp{
		Success: publishResp.Success,
	}, nil
}
