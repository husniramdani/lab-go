package main
import "fmt"

type person struct {
  name string
  age int
}

// embeded struct
type student struct {
    grade int
    person
}


func main(){
  var s1 = student{}
  // s1.name = "wick"
  // s1.grade = 2

  // var s2 = student{"ethan", 2}
  // var s3 = student{name: "jason"}

  // fmt.Println("student 1 :", s1.name)
  // fmt.Println("student 2 :", s2.name)
  // fmt.Println("student 3 :", s3.name)

  s1.name = "wick"
  s1.age = 21
  s1.grade = 2

  fmt.Println("name  :", s1.name)
  fmt.Println("age   :", s1.age)
  fmt.Println("age   :", s1.person.age)
  fmt.Println("grade :", s1.grade)

  fmt.Println("\n")
  // kombinasi slice & struct
  var allStudents = []person{
    {name: "Wick", age: 23},
    {name: "Ethan", age: 23},
    {name: "Bourne", age: 22},
  }

  for _, student := range allStudents {
    fmt.Println(student.name, "age is", student.age)
  }
}

// tag property dalam struct
// biasa untuk encode/decode data json
type personTag struct {
  name string `tag1`
  age int `tag2`
}
