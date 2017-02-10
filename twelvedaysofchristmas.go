/*
December 19th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
No bonuses
*/

package main

import (
  "fmt"
)

func main(){
  day := [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
  gift := [12]string{"Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds", "Golden Rings", "Geese a Laying",
                  "Swans a Swimming", "Maids a Milking", "Ladies Dancing", "Lords a Leaping", "Pipers Piping", "Drummers Drumming"}

  for x, theDay := range day{//cycle through all the days
    fmt.Println("On the " + theDay + " day of Christmas\nmy true love sent to me:")
    //start in the xth position of gift and work down to 0
    for i := x; i >= 0; i--{
      if i == 0 && x != 0{fmt.Print("and ")}
      fmt.Println(i+1, gift[i])
    }

  }
}
