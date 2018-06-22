package core

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestReviewModel_Get(t *testing.T) {
	type mock struct {
		query   string
		want    *sqlmock.Rows
		wantErr error
	}
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mock       mock
		wantReview Review
		wantErr    bool
	}{
		{
			name: "Success",
			args: args{
				id: 100,
			},
			mock: mock{
				query: "SELECT feedback_id, review FROM ws_product_feedback WHERE feedback_id = (.+);",
				want: sqlmock.NewRows([]string{"feedback_id", "review"}).
					AddRow(100, "Hello World"),
				wantErr: nil,
			},
			wantReview: Review{
				ID:      100,
				Message: "Hello World",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			args := ReviewModel{
				DB: sqlx.NewDb(db, "sqlmock"),
			}

			mock.ExpectQuery(tt.mock.query).WillReturnRows(tt.mock.want).WillReturnError(tt.mock.wantErr)

			gotReview, err := args.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReview, tt.wantReview) {
				t.Errorf("ReviewModel.Get() = %v, want %v", gotReview, tt.wantReview)
			}
		})
	}
}
func TestReviewModel_IGet(t *testing.T) {
	type mock struct {
		query   string
		want    *sqlmock.Rows
		wantErr error
	}
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mock       mock
		wantReview Review
		wantErr    bool
	}{
		{
			name: "Success",
			args: args{
				id: 100,
			},
			mock: mock{
				query: "SELECT feedback_id, review FROM ws_product_feedback WHERE feedback_id = (.+);",
				want: sqlmock.NewRows([]string{"feedback_id", "review"}).
					AddRow(100, "Hello World"),
				wantErr: nil,
			},
			wantReview: Review{
				ID:      100,
				Message: "Hello World",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			args := ReviewModel{
				DB: sqlx.NewDb(db, "sqlmock"),
			}

			mock.ExpectQuery(tt.mock.query).WillReturnRows(tt.mock.want).WillReturnError(tt.mock.wantErr)

			gotReview, err := IReview(args).Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReviewModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReview, tt.wantReview) {
				t.Errorf("ReviewModel.Get() = %v, want %v", gotReview, tt.wantReview)
			}
		})
	}
}

func BenchmarkReviewModel_Get(b *testing.B) {
	db, _, _ := sqlmock.New()
	args := ReviewModel{
		DB: sqlx.NewDb(db, "sqlmock"),
	}

	for i := 0; i < b.N; i++ {
		args.Get(1)
	}
}

func BenchmarkReviewModel_IGet(b *testing.B) {
	db, _, _ := sqlmock.New()
	args := ReviewModel{
		DB: sqlx.NewDb(db, "sqlmock"),
	}

	for i := 0; i < b.N; i++ {
		x := IReview(args)
		x.Get(1)
	}
}
