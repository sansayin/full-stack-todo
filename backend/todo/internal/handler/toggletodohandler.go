package handler

import (
	"net/http"

	"github.com/sansayin/todo-zero/common/result"
	"github.com/sansayin/todo-zero/todo/internal/logic"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func toggleTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ToggleTodoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewToggleTodoLogic(r.Context(), svcCtx)
		resp, err := l.ToggleTodo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
