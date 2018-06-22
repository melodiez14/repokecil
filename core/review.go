package core

// ReviewModel is struct which is composed from Model struct
type ReviewModel Model

// IReview is the interface of ReviewModel struct.
type IReview interface {
	Get(id int64) (review Review, err error)
}

// Review is a struct used to scan the data from database
type Review struct {
	ID      int64  `db:"feedback_id"`
	Message string `db:"review"`
}

// Get returns the data from database
// FYI. This function has high cost
// There is 24 allocations to the heap for every call to this function
// To test this performance, you can bench this function which already been created in review_test.go
func (args ReviewModel) Get(id int64) (review Review, err error) {
	err = args.DB.Get(&review, "SELECT feedback_id, review FROM ws_product_feedback WHERE feedback_id = $1;", id)
	return
}
