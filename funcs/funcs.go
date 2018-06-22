package funcs

import (
	"github.com/jmoiron/sqlx"
	"github.com/melodiez14/repokecil/core"
)

// Funcs is the struct of all core module
type Funcs struct {
	Review core.IReview
}

// IFuncs is the interface of all available functions in this package
// This is used for mocking from another package
type IFuncs interface {
	GetReview(id int64) (review core.Review, err error)
}

// Init function is used to create pointer semantics Funcs object.
// *Funcs consist of all core modules
// This function will escaped to the heap.
// The reason is same as *sql.DB (to make a changes of this object available for all place. Safety reason)
func Init(db *sqlx.DB) *Funcs {
	return &Funcs{
		Review: core.ReviewModel{
			DB: db,
		},
	}
}
