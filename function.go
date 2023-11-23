package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
  "strings"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix())) 

func main(){
  fmt.Println("=====Function=====")
  var names = []string{"John", "Wick"}
  printMessage("Haloo", names)

  // random value
  var randomValue int
  for i := 0; i < 3; i++ {
    randomValue = randomWithRange(2, 10)
    fmt.Println("Random number : ", randomValue)
  }

  // calculate
  var diameter float64 = 15
  var area, circumference = calculate2(diameter)
  fmt.Printf("luas lingkaran: %.2f \n", area)
  fmt.Printf("keliling lingkaran: %.2f \n", circumference)
}  

func printMessage(message string, arr []string){
  var nameString = strings.Join(arr, " ")
  fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
  var value = randomizer.Int()%(max-min+1) + min
  return value
}

// function multiple return
func calculate(d float64) (float64, float64) {
  // hitung luas
  var area = math.Pi * math.Pow(d / 2, 2)
  // hitung keliling
  var circumference = math.Pi * d

  // kembalikan 2 nilai
  return area, circumference
}

// function with predefined return value
func calculate2(d float64)(area float64, circumference float64){
  area = math.Pi * math.Pow(d / 2, 2)
  circumference = math.Pi * d
  
  return
}
