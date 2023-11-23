// A.20 Fungsi variadic 
// A.21 Fungsi Closure - anonymouse function
// A.22 Fungsi sebagai parameter
package main

import (
  "fmt"
  "strings"
)

func main(){
  fmt.Println("===== function 2 =====")
  var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
  var avg = calculate(numbers...)
  var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
  fmt.Println(msg)

  // fungsi closure
  var getMinMax = func(n []int) (int, int){
    var min, max int
    for i, e := range n {
      switch {
      case i == 0:
        max, min = e, e
      case e > max:
        max = e
      case e < min:
        min = e
      }
    }
    return min, max
  }

  var numbers2 = []int{2,3,4,3,4,2,3}
  var min, max = getMinMax(numbers2)
  fmt.Printf("Data : %v\nmin : %v\nmax : %v\n", numbers2, min, max);

  fmt.Println("\n")
  // IIFE (Immediately-Invoked Function Expression)
  var numbers3 = []int{2,3,0,4,3,2,0,4,2,0,3}
  var newNumbers3 = func(min int) []int {
    var result []int
    for _, e := range numbers3 {
      if e < min {
        continue
      }
      result = append(result,e)
    }
    return result
  }(3)
  fmt.Println("Original numbers :", numbers3)
  fmt.Println("Filtered number :", newNumbers3)

  fmt.Println("\n")
  // call closure as return
  var max4 = 3
  var numbers4 = []int{2,3,0,4,3,2,0,4,2,0,3}
  var howMany, getNumbers = findMax(numbers4, max4)
  var theNumbers = getNumbers()

  fmt.Println("numbers\t:", numbers4)
  fmt.Printf("find \t: %d\n\n", max4)

  fmt.Println("found \t:", howMany)    // 9
  fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

  fmt.Println("\n")
  // Fungsi sebagai parameter
  var data5 = []string{"wick", "jason", "ethan"}
  var dataContainsO = filter(data5, func(each string) bool {
    return strings.Contains(each, "o")
  })
  var dataLength5 = filter(data5, func(each string) bool {
    return len(each) == 5
  })

  fmt.Println("data asli \t:", data5)
  // data asli : [wick jason ethan]
  fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
  // filter ada huruf "o" : [json]
  fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)

}

func calculate(numbers ...int) float64 {
  var total int = 0
  for _, number := range numbers {
    total += number
  }

  var avg = float64(total) / float64(len(numbers))
  return avg
}

// Closure sebagai nilai kembalian
func findMax(numbers []int, max int) (int, func() []int){
  var res []int
  for _, e := range numbers {
    if e <= max {
      res = append(res, e)
    }
  }

  var getNumbers = func() []int {
    return res
  }
  return len(res), getNumbers
}

type FilterCallback func(string) bool

// Fungsi sebagai parameter
func filter(data []string, callback FilterCallback) []string {
  var result []string
  for _, each := range data {
    if filtered := callback(each); filtered {
      result = append(result, each)
    }
  }
  return result
}


