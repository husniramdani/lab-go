package main

import "fmt"

func main() {
	fmt.Println("Array was here")

	var names [4]string
	names[0] = "trafalgar"
	names[1] = "budi"
	names[3] = "supri"
	// names[4] = "error nih" // error out of bound

	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// inisiasi tanpa jumlah element
	var numbers = [...]int{1,2,3,4,5}

	fmt.Println("Jumlah element \t :", len(numbers))

	fmt.Println("\n")

	// array multi dimensi
	var numbers1 = [2][3]int{[3]int{1,2,3}, [3]int{4,5,6}}
	var numbers2 = [2][3]int{{1,2,3},{4,5,6}}

	fmt.Println("Numbers 1 : ", numbers1)
	fmt.Println("Numbers 2 : ", numbers2)

	fmt.Println("\n")

	// looping array
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Element %d : %s \n", i, fruits[i])
	}

	fmt.Println("\n")
	// looping array using range
	for i, fruit := range fruits {
		fmt.Printf("Element %d : %s \n", i, fruit)
	}

	// Alokasi element array menggunakan keyword make
	var buahs = make([]string, 2)
	buahs[0] = "apple"
	buahs[1] = "manggo"

	fmt.Println(buahs)

	// Slice seperti Array tapi tidak perlu define jumlah element
	var tamago = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(tamago[0]) // "apple"

	// tamago[0:2] [apple, grape]
	// tamago[:] [apple, grape, banana, melon]
	// tamago[2:] [banana, melon]
	// tamago[:2] [apple, grape]

	// Fungsi cap()
	// digunakan untuk menghitung lebar atau kapasitas maksimum Slice

	// Fungsi append()
	// untuk menambahkan elemen pada slice
	var cFruits = append(tamago, "Papaya")
	fmt.Println(cFruits)

	// copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	resultCopy := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(resultCopy)
}
