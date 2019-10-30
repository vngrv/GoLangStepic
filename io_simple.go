package main
import "fmt"
func main() {
    var name string
    var age int
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    fmt.Print("Enger your age: ")
    fmt.Scan(&age)
    fmt.Println(name, age)
}
