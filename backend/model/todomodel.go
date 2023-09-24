package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TodoModel = (*customTodoModel)(nil)

type (
	// TodoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTodoModel.
	TodoModel interface {
		todo1Model
	}

	customTodoModel struct {
		*defaultTodoModel
	}
)

// NewTodoModel returns a model for the database table.
func NewTodoModel(conn sqlx.SqlConn, c cache.CacheConf) TodoModel {
	return &customTodoModel{
		defaultTodoModel: newTodoModel(conn, c),
	}
}
