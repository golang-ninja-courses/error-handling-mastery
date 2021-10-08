package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID    string
	Email string
}

type ByEmail []User

func (s ByEmail) Len() int {
	return len(s)
}

func (s ByEmail) Less(i, j int) bool {
	return s[i].Email < s[j].Email
}

func (s ByEmail) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{ID: "1", Email: "bob@gmail.com"},
		{ID: "2", Email: "alex@gmail.com"},
		{ID: "2", Email: "alice@gmail.com"},
	}

	sort.Sort(ByEmail(users))

	// [{2 alex@gmail.com} {2 alice@gmail.com} {1 bob@gmail.com}]
	fmt.Println(users)
}
