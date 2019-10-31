package main
import "fmt"
func main(){
    var number int
    fmt.Scan(&number)
    if number < 10000 {
       fmt.Sprintf("%v",number)
    }
}
