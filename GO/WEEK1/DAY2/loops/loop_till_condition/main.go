package main

func main() {
	var i int
	for i < 5 {
		println(i)
		i++
		// if i == 3 {
		// 	break
		// }
		if i == 4 {
			continue
		}
		println("continuing...")

	}
}
