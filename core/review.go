package core

type ReviewModel Model

type IReview interface {
	Get(id int64) (review Review, err error)
}

type Review struct {
	ID      int64  `db:"feedback_id"`
	Message string `db:"review"`
}

func (args ReviewModel) Get(id int64) (review Review, err error) {
	err = args.DB.Get(&review, "SELECT feedback_id, review FROM ws_product_feedback WHERE feedback_id = $1;", id)
	return
}
