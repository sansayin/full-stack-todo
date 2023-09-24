package handler

import (
	"net/http"

	"github.com/sansayin/todo-zero/todo/internal/logic"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteTodosHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteTodosLogic(r.Context(), svcCtx)
		resp, err := l.DeleteTodos()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
