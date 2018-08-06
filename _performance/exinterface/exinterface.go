package exinterface

// User is the struct
type User struct {
	FirstName string
	LastName  string
	Age       int8
}

type Response struct {
	Code int64
	Data interface{}
}

func New() User {
	return User{
		FirstName: "Risal",
		LastName:  "Falah",
		Age:       23,
	}
}

//go:noinline
func GetResponse() Response {
	user := New()
	return Response{
		Code: 200,
		Data: user,
	}
}

func (u *User) Print() {
	println(u.FirstName, u.LastName)
}

func (u *User) Scan(firstname, lastname string) {
	u.FirstName = firstname
	u.LastName = lastname
}

type IPrinter interface {
	Print()
	Scan(firstname, lastname string)
}
