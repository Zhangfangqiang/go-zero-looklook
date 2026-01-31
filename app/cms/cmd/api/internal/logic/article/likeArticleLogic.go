// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package article

import (
	"context"

	"looklook/app/cms/cmd/api/internal/svc"
	"looklook/app/cms/cmd/api/internal/types"
	"looklook/app/cms/cmd/rpc/cms"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikeArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 点赞文章
func NewLikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeArticleLogic {
	return &LikeArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeArticleLogic) LikeArticle(req *types.LikeArticleReq) (resp *types.LikeArticleResp, err error) {
	likeResp, err := l.svcCtx.CmsRpc.LikeArticle(l.ctx, &cms.LikeArticleReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "like article failed, req: %+v", req)
	}

	return &types.LikeArticleResp{
		LikeCount: likeResp.LikeCount,
	}, nil
}
