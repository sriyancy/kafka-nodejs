package main

type User struct {
	ID        int
	Firstname string
	Lastname  string
}

func main() {
	u1 := User{
		ID:        1,
		Firstname: "ford",
		Lastname:  "dent",
	}
	u2 := User{
		ID:        2,
		Firstname: "ford",
		Lastname:  "perfect",
	}
	if u1.ID == u2.ID {
		println("same user")
	} else if u1.Firstname == u2.Firstname {
		println("similat user")
	}
}
