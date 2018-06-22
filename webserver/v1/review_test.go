package v1

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/repokecil/core"
	"github.com/melodiez14/repokecil/funcs"
)

type FuncsTest struct {
	Review core.ReviewModel
}

func (model FuncsTest) GetReview(id int64) (review core.Review, err error) {
	switch id {
	case 1, 2:
		return review, sql.ErrConnDone
	default:
		review.ID = id
		review.Message = strconv.FormatInt(id, 10)
		return
	}
}

func Test_handler_iGetReviewHandler(t *testing.T) {
	type fields struct {
		f funcs.IFuncs
	}
	type args struct {
		w  *httptest.ResponseRecorder
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
	}{
		{
			name: "Non Numeric ID",
			fields: fields{
				f: FuncsTest{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  func() *http.Request { r, _ := http.NewRequest(http.MethodGet, "/api/review/v1", nil); return r }(),
				ps: httprouter.Params{},
			},
			wantCode: 403,
		},
		{
			name: "Review Not Found",
			fields: fields{
				f: FuncsTest{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  func() *http.Request { r, _ := http.NewRequest(http.MethodGet, "/api/review/v1?id=1", nil); return r }(),
				ps: httprouter.Params{},
			},
			wantCode: 500,
		},
		{
			name: "Success",
			fields: fields{
				f: FuncsTest{},
			},
			args: args{
				w:  httptest.NewRecorder(),
				r:  func() *http.Request { r, _ := http.NewRequest(http.MethodGet, "/api/review/v1?id=20", nil); return r }(),
				ps: httprouter.Params{},
			},
			wantCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				f: tt.fields.f,
			}

			h.iGetReviewHandler(tt.args.w, tt.args.r, tt.args.ps)
			if tt.args.w.Code != tt.wantCode {
				t.Fatalf("Non-expected status code %v:\n\tbody: %v", tt.wantCode, tt.args.w.Code)
			}
		})
	}
}
