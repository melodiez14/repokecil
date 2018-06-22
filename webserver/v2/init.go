package v2

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/repokecil/funcs"
)

type handler struct {
	f funcs.IFuncs
}

func InitRouter(f funcs.IFuncs, r *httprouter.Router) {
	h := handler{f: f}

	r.GET("/api/review/v2", h.iGetReviewHandler)
}
