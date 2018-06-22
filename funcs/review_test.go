package funcs

import (
	"reflect"
	"testing"

	"github.com/melodiez14/repokecil/core"
)

func TestFuncs_GetReview(t *testing.T) {
	type fields struct {
		Review core.IReview
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantReview core.Review
		wantErr    bool
	}{
		{
			name: "No Error",
			fields: fields{
				Review: ReviewModelTest{},
			},
			args: args{
				id: 100,
			},
			wantReview: core.Review{
				ID:      100,
				Message: "100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := Funcs{
				Review: tt.fields.Review,
			}
			gotReview, err := model.GetReview(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Funcs.GetReview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReview, tt.wantReview) {
				t.Errorf("Funcs.GetReview() = %v, want %v", gotReview, tt.wantReview)
			}
		})
	}
}
