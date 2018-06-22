package v2

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h handler) iGetReviewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(403)
		return
	}

	review, err := h.f.GetReview(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(review.Message))
}
