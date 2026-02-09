package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TPriClearingDataGeneralYunnanModel = (*customTPriClearingDataGeneralYunnanModel)(nil)

type (
	// TPriClearingDataGeneralYunnanModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTPriClearingDataGeneralYunnanModel.
	TPriClearingDataGeneralYunnanModel interface {
		tPriClearingDataGeneralYunnanModel
	}

	customTPriClearingDataGeneralYunnanModel struct {
		*defaultTPriClearingDataGeneralYunnanModel
	}
)

// NewTPriClearingDataGeneralYunnanModel returns a model for the database table.
func NewTPriClearingDataGeneralYunnanModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TPriClearingDataGeneralYunnanModel {
	return &customTPriClearingDataGeneralYunnanModel{
		defaultTPriClearingDataGeneralYunnanModel: newTPriClearingDataGeneralYunnanModel(conn, c, opts...),
	}
}
