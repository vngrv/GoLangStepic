// Определите является ли билет счастливым. 
// Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

package main
import "fmt"
func main() {
  var number, sumLeft, sumRight int
  fmt.Scan(&number)
  sumRight = number % 1000
  sumLeft = number / 1000
  
  var sumLeftOne, sumLeftTwo, sumLeftThree int
  sumLeftOne = sumLeft / 100
  sumLeftTwo = (sumLeft % 100) / 10
  sumLeftThree = sumLeft % 10
  
  var sumRightOne, sumRightTwo, sumRightThree int
  sumRightOne = sumRight / 100
  sumRightTwo = (sumRight % 100) / 10
  sumRightThree = sumRight % 10
  sumLeft = sumLeftOne + sumLeftTwo + sumLeftThree
  sumRight = sumRightOne + sumRightTwo + sumRightThree

  if(sumLeft == sumRight){
    fmt.Print("YES")
  } else {
    fmt.Print("NO")
  }
}
