package handler

import (
	"net/http"

	"github.com/sansayin/todo-zero/common/result"

	"github.com/sansayin/todo-zero/todo/internal/logic"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Todo
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewCreateTodoLogic(r.Context(), svcCtx)
		resp, err := l.CreateTodo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
