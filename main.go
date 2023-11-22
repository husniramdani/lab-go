package main

import "fmt"

func main() {
	// variable
	var firstName string = "john"
	var lastName string

	// print
	fmt.Printf("Halo %s %s!\n", firstName, lastName)
	fmt.Println("john", "wick")

	// multiple const declaration
	const (
		square = "kotak"
		isToday bool = true
		numeric uint8 = 1
		floatNum = 2.2
	)

	// switch
	var point = 6
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// looping
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	var i = 0;
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	i = 0;
	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}

	// looping range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}


	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kvs {
	    fmt.Println("Key=", k, "Value=", v)
	}


	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("Done")
	}
}
