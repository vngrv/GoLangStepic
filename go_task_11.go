package main
import "fmt"
func main(){
    var number int
    fmt.Scan(&number)
    if number > 0 {
        fmt.Print("Число положительное")
    } else if number < 0 {
        fmt.Print("Число отрицательное")
    } else {
        fmt.Print("Ноль") 
    }
}
