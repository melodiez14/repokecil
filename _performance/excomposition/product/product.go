package product

type response struct {
	StatusCode int
	Err        error
}

type product struct {
	ID   int64
	Name string
}

type GetProducts struct {
	response
	List []product
}

func Get() (x GetProducts) {
	return
}

func GetRouter() (URL string) {
	return
}
