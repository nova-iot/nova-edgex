package rule

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"novaedgex/internal/logic/rule"
	"novaedgex/internal/svc"
	"novaedgex/internal/types"
)

// 创建规则
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRuleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rule.NewCreateLogic(r.Context(), svcCtx)
		err := l.Create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
