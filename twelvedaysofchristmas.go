/*
December 19th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
Bonuses 1 & 2 Completed
*/

package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  day := [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
  gift := [12]string{}
  fmt.Println("Enter the 12 gifts to receive on Christmas: ")
  for x, _ := range gift{
    fmt.Print("Gift ", x+1, ": ")
    gift[x], _ = reader.ReadString('\n')
  }
  fmt.Println("Entries Accepted.\n\n")
  numString := [12]string{"a", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}


  for x, theDay := range day{//cycle through all the days
    fmt.Println("On the " + theDay + " day of Christmas\nmy true love sent to me:")
    //start in the xth position of gift and work down to 0
    for i := x; i >= 0; i--{
      if i == 0 && x != 0{fmt.Print("and ")}
      fmt.Println(numString[i], gift[i])
    }

  }
}
