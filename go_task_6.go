package main
import "fmt"
func main(){
    var value, hours, minutes int
    fmt.Scan(&value)
    hours = (value*2) / 60
    minutes = (value*2) %  60
    fmt.Print("It is ",hours, " hours ", minutes, " minutes.")
}
