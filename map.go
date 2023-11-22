package main
import "fmt"

func main(){
	fmt.Println("======= MAP ======")

	var chicken map[string]int
	chicken = map[string]int{} // inisialisasi nilai awal {}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("Mei", chicken["mei"]) // mei 0

	var chicken2 = map[string]int{
		"januari": 50,
		"februari": 40,
	}
	fmt.Println("Chicken 2 : ", chicken2)

	// deklarasi map tanpa nilai awal
	// var chicken3 = map[string]int{}
	// var chicken4 = make(map[string]int)
	// var chicken5 = *new(map[string]int)

	// menghapus item map
	// delete(chicken, "januari")
	// fmt.Println(len(chicken)) // 1

	// A.17.5 deteksi keberadaan item dengan key tertentu
	var value, isExist = chicken["januari"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exists")
	}

	// A.17.6 kombinasi slice dan map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	fmt.Println()

	// dalam []map[string]string tiap elemen bisa memiliki key yang berbeda-beda
	var chickenDatas = []map[string]string {
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	checkEmpty := func(value string, isExist bool) string {
		if isExist {
			return value
		}
		return "-"
	}

	for _, chicken := range chickenDatas {
		var name, isName = chicken["name"]
		var gender, isGender = chicken["gender"]
		var color, isColor = chicken["color"]
		var id, isId = chicken["id"]
		var address, isAddress = chicken["address"]
		var community, isCommunity = chicken["community"]

		fmt.Println("====================")
		fmt.Printf("name : %s\n", checkEmpty(name, isName))
		fmt.Printf("gender : %s\n", checkEmpty(gender, isGender))
		fmt.Printf("color : %s\n", checkEmpty(color, isColor))
		fmt.Println("====================")

		fmt.Printf("id : %s\n", checkEmpty(id, isId))
		fmt.Printf("address : %s\n", checkEmpty(address, isAddress))
		fmt.Println("====================")

		fmt.Printf("community : %s\n", checkEmpty(community, isCommunity))
		fmt.Println("====================\n\n")
	}
}
