// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/sansayin/todo-zero/common/globalkey"
)

var (
	todoFieldNames          = builder.RawFieldNames(&Todo{}, true)
	todoRows                = strings.Join(todoFieldNames, ",")
	todoRowsExpectAutoSet   = strings.Join(stringx.Remove(todoFieldNames,"completed","delete_time","del_state","version", "create_time", "update_time"), ",")
	todoRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(todoFieldNames, "id", "create_time", "update_time"))

	cachePublicTodoIdPrefix = "cache:public:todo:id:"
 	cachePublicAllTodoPrefix = "cache:public:todo:all"
)

type (
	todo1Model interface {
		Insert(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error)

		FindOne(ctx context.Context, id string) (*Todo, error)
		UpdateOne(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error)
 		ToggleTodo(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error)

		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Todo) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *Todo) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Todo, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Todo, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Todo, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Todo, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Todo, error)

		Delete(ctx context.Context, session sqlx.Session, id string) error
	}

	defaultTodoModel struct {
		sqlc.CachedConn
		table string
	}

	Todo struct {
		Id          string    `db:"id"`
		Description string    `db:"description"`
		Completed   bool      `db:"completed"`
		CreateTime  time.Time `db:"create_time"`
		UpdateTime  time.Time `db:"update_time"`
		DeleteTime  time.Time `db:"delete_time"`
		DelState    int64     `db:"del_state"`
		Version     int64     `db:"version"`
	}
)

func newTodoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTodoModel {
	return &defaultTodoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."todo"`,
	}
}

func (m *defaultTodoModel) Insert(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	var resp Todo

	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2) returning *", m.table, todoRowsExpectAutoSet)
		if session != nil {
			return nil, session.QueryRowCtx(ctx, &resp, query, data.Id, data.Description )
		}
		return nil, conn.QueryRowCtx(ctx, &resp, query, data.Id, data.Description)
	}, publicTodoIdKey)

	return &resp, err
}

func (m *defaultTodoModel) FindOne(ctx context.Context, id string) (*Todo, error) {
	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, id)
	var resp Todo
	err := m.QueryRowCtx(ctx, &resp, publicTodoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 and del_state = 0", todoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) UpdateOne(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error) {
	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, data.Id)
	var resp Todo
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1 returning *", m.table, todoRowsWithPlaceHolder)
		if session != nil {
			return nil, session.QueryRowCtx(ctx, &resp, query, data.Id, data.Description, data.Completed, data.CreateTime, data.UpdateTime, data.DeleteTime, data.DelState, data.Version)
		}
		return nil, conn.QueryRowCtx(ctx, &resp, query, data.Id, data.Description, data.Completed, data.CreateTime, data.UpdateTime, data.DeleteTime, data.DelState, data.Version)
	}, publicTodoIdKey)
	return &resp, err
}

func (m *defaultTodoModel) ToggleTodo(ctx context.Context, session sqlx.Session, data *Todo) (*Todo, error) {
	m.DelCacheCtx(ctx, cachePublicAllTodoPrefix)
  var resp Todo
	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, data.Id)
  _,err:=m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1 returning *", m.table, "completed=$2, update_time=CURRENT_TIMESTAMP")
		return nil,conn.QueryRowCtx(ctx,&resp, query, data.Id, data.Completed)
	}, publicTodoIdKey)
  return &resp, err
}


func (m *defaultTodoModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Todo) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1 and version = ? ", m.table, todoRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Description, data.Completed, data.CreateTime, data.UpdateTime, data.DeleteTime, data.DelState, data.Version, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.Description, data.Completed, data.CreateTime, data.UpdateTime, data.DeleteTime, data.DelState, data.Version, oldVersion)
	}, publicTodoIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultTodoModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *Todo) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "TodoModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultTodoModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTodoModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTodoModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Todo, error) {

	builder = builder.Columns(todoRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = 0").ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Todo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Todo, error) {

	builder = builder.Columns(todoRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Todo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Todo, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(todoRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*Todo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultTodoModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Todo, error) {

	builder = builder.Columns(todoRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Todo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Todo, error) {

	builder = builder.Columns(todoRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Todo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTodoModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultTodoModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTodoModel) Delete(ctx context.Context, session sqlx.Session, id string) error {
	publicTodoIdKey := fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, publicTodoIdKey)
	return err
}
func (m *defaultTodoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicTodoIdPrefix, primary)
}
func (m *defaultTodoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 and del_state = ? limit 1", todoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultTodoModel) tableName() string {
	return m.table
}
