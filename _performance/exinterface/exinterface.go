package exinterface

// User is the struct
type User struct {
	FirstName string
	LastName  string
}

func New() User {
	return User{
		FirstName: "Risal",
		LastName:  "Falah",
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
