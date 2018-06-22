package conn

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestInitDB(t *testing.T) {
	type args struct {
		cfg DBConfig
	}
	tests := []struct {
		name string
		args args
		want *sqlx.DB
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDB(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
