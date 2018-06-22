package funcs

import (
	"github.com/melodiez14/repokecil/core"
)

func (model Funcs) GetReview(id int64) (review core.Review, err error) {
	return model.Review.Get(id)
}
