package logic

import (
	"context"

	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTodoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTodoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTodoLogic {
	return &GetTodoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTodoLogic) GetTodo(req *types.GetTodoReq) (resp *types.Todo, err error) {
	todo, err := l.svcCtx.Model.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.Todo{
		Id:          todo.Id,
		Description: todo.Description,
		Completed:   todo.Completed,
	}, err
}
