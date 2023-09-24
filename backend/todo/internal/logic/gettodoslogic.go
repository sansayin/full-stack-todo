package logic

import (
	"context"

	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTodosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTodosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTodosLogic {
	return &GetTodosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTodosLogic) GetTodos() (resp []types.Todo, err error) {
	builder := l.svcCtx.Model.SelectBuilder()
	res, err := l.svcCtx.Model.FindAll(l.ctx, builder, "description desc")
	if err != nil {
		return nil, err
	}
	var data []types.Todo
	for _, todo := range res {
		data = append(data, types.Todo{
			Id:          todo.Id,
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	}
	// todo: add your logic here and delete this line

	return data, nil
}
