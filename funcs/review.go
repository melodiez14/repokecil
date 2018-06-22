package funcs

import (
	"github.com/melodiez14/repokecil/core"
)

// GetReview is just a sample function to get a review from core.
// In production, this function usualy has a validatation if the specified user has a privilege access the core or not
func (model Funcs) GetReview(id int64) (review core.Review, err error) {
	return model.Review.Get(id)
}
