package svc

import (
	_ "github.com/lib/pq"
	"github.com/sansayin/todo-zero/model"
	"github.com/sansayin/todo-zero/todo/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TodoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Model: model.NewTodoModel(sqlx.NewSqlConn("postgres", c.Psql.DataSource), c.CacheRedis),
	}
}
