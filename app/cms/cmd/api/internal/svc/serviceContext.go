// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"looklook/app/cms/cmd/api/internal/config"
	"looklook/app/cms/cmd/rpc/cms"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	CmsRpc cms.Cms
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CmsRpc: cms.NewCms(zrpc.MustNewClient(c.CmsRpcConf)),
	}
}
