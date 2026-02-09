package svc

import (
	"looklook/app/cms/cmd/rpc/internal/config"
	"looklook/app/cms/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                             config.Config
	ArticleModel                       model.ArticleModel
	TPriClearingDataGeneralYunnanModel model.TPriClearingDataGeneralYunnanModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:                             c,
		ArticleModel:                       model.NewArticleModel(sqlConn, c.Cache),
		TPriClearingDataGeneralYunnanModel: model.NewTPriClearingDataGeneralYunnanModel(sqlConn, c.Cache),
	}
}
