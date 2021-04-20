package main

func main() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	for j, v := range slice {
		println(j, v)
	}
	//index and value
	//map
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for j, v := range wellKnownPorts {
		println(j, v)
	}
	for v := range wellKnownPorts {
		println(v)
	}
	for _, v := range wellKnownPorts {
		println(v)
	}
}
