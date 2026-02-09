package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
		FindPageListByPage(ctx context.Context, page, pageSize int64, category string, status int64) ([]*Article, int64, error)
		UpdateLikeCount(ctx context.Context, id int64) error
		FindAllByCategory(ctx context.Context, category string) ([]*Article, error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn, c, opts...),
	}
}

// Trans 事务支持
func (m *customArticleModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

// FindPageListByPage 分页查询文章列表
func (m *customArticleModel) FindPageListByPage(ctx context.Context, page, pageSize int64, category string, status int64) ([]*Article, int64, error) {
	// 构建查询条件
	query := squirrel.Select().From(m.table).Where("del_state = ?", 0)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	// 查询总数
	countQuery := query.Column("COUNT(*)")
	countSql, countArgs, err := countQuery.ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = m.QueryRowNoCacheCtx(ctx, &total, countSql, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	// 查询分页数据
	dataQuery := query.Column(articleRows).OrderBy("publish_time DESC").Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
	dataSql, dataArgs, err := dataQuery.ToSql()
	if err != nil {
		return nil, 0, err
	}

	var resp []*Article
	err = m.QueryRowsNoCacheCtx(ctx, &resp, dataSql, dataArgs...)
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

// UpdateLikeCount 更新点赞数
func (m *customArticleModel) UpdateLikeCount(ctx context.Context, id int64) error {
	articleIdKey := fmt.Sprintf("%s%v", cacheArticleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set like_count = like_count + 1 where id = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, articleIdKey)
	return err
}

// FindAllByCategory 根据分类查询所有文章
func (m *customArticleModel) FindAllByCategory(ctx context.Context, category string) ([]*Article, error) {
	query := fmt.Sprintf("select %s from %s where category = ? and del_state = 0 and status = 1 order by publish_time desc", articleRows, m.table)

	var resp []*Article
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, category)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
