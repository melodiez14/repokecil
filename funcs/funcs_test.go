package funcs

import (
	"database/sql"
	"reflect"
	"strconv"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/repokecil/core"
)

type FuncsTest struct {
	Review core.ReviewModel
}

type ReviewModelTest core.Model

func (args ReviewModelTest) Get(id int64) (review core.Review, err error) {
	switch id {
	case 1:
		return core.Review{}, sql.ErrNoRows
	case 2:
		return core.Review{}, sql.ErrConnDone
	default:
		review.ID = id
		review.Message = strconv.FormatInt(id, 10)
		return
	}
}

func TestInit(t *testing.T) {
	db := &sqlx.DB{}
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *Funcs
	}{
		{
			name: "Success",
			args: args{
				db: db,
			},
			want: &Funcs{
				Review: core.ReviewModel{
					DB: db,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInit(b *testing.B) {
	var db *sqlx.DB
	for i := 0; i < b.N; i++ {
		Init(db)
	}
}
