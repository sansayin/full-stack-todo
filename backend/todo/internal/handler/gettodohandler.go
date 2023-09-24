package handler

import (
	"fmt"
	"net/http"

	"github.com/sansayin/todo-zero/common/result"

	"github.com/sansayin/todo-zero/todo/internal/logic"
	"github.com/sansayin/todo-zero/todo/internal/svc"
	"github.com/sansayin/todo-zero/todo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getTodoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTodoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetTodoLogic(r.Context(), svcCtx)
		resp, err := l.GetTodo(&req)
    fmt.Printf("%+v", resp)
		//httpx.OkJsonCtx(r.Context(), w, resp)
		result.HttpResult(r, w, resp, err)
	}
}
