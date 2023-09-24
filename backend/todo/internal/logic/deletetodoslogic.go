package logic

import (
	"context"

	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTodosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTodosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTodosLogic {
	return &DeleteTodosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTodosLogic) DeleteTodos() (resp *types.RespStatus, err error) {
	// todo: add your logic here and delete this line

	return
}
