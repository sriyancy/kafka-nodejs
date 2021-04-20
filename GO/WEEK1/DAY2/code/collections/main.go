package main

import "fmt"

func main() {
	//array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	//slice
	//built on top of array
	slice := arr[:]
	arr[1] = 47
	slice[2] = 27
	fmt.Println(arr, slice)

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	s2 := slice1[1:]
	s3 := slice1[:2]
	s4 := slice1[1:2]

	fmt.Println(s2, s3, s4)

	//maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)
	//strct
	type user struct {
		ID        int
		firstname string
		lastname  string
	}
	var u user
	u.ID = 1
	u.firstname = "sai"
	u.lastname = "prem"
	fmt.Println(u)

	u2 := user{ID: 1, firstname: "vijay", lastname: "rohan"}
	fmt.Println(u2)
	fmt.Println(u2.firstname)
}
