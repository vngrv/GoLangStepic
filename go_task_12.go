package main
import "fmt"
func main(){
    var number,a,b,c int
    fmt.Scan(&number)
    a = (number -number % 100)/100 
    b = number % 100 / 10
    c = number % 100 % 10
    if a != b && a != c && b != c{
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
