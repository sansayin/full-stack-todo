package logic

import (
	"context"

	"github.com/sansayin/todo-zero/model"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTodoLogic {
	return &CreateTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTodoLogic) CreateTodo(req *types.Todo) (*types.Todo, error) {
	todo, err := l.svcCtx.Model.Insert(l.ctx, nil, &model.Todo{
		Id:          req.Id,
		Description: req.Description,
		Completed:   req.Completed,
	})
	if err != nil {
		return nil, err
	}
  return &types.Todo{
    Id:todo.Id,
    Description: todo.Description,
    Completed: todo.Completed,
  }, nil
}
