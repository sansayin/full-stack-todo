package logic

import (
	"context"

	"github.com/sansayin/todo-zero/model"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type ToggleTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewToggleTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ToggleTodoLogic {
	return &ToggleTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ToggleTodoLogic) ToggleTodo(req *types.ToggleTodoReq) (resp *types.Todo, err error) {
	result, err := l.svcCtx.Model.ToggleTodo(l.ctx, nil, &model.Todo{
		Id:          req.Id,
		Description: req.Description,
		Completed:   req.Completed,
	})
	var todo types.Todo
	todo.Id = result.Id
	todo.Completed = result.Completed
	todo.Description = result.Description

	return &todo, err 
}
