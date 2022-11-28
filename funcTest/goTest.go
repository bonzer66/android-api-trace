// testing functions before adding to main package
package main

/* get flags to work 
type Input struct {
      
      }
      
      */
import (
        //"flag"
        "fmt"
        "os"
        
)

// takes a file as input
func main() {

        if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }
    data, err := os.ReadFile(os.Args[1]) //os.Args[2] <<-- for sub commands
    if err != nil {
        fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }
    fmt.Println(string(data))
 }




/*
func OutputFile(filename string) {
        out,d1 := []byte(out,d1)
        err := os.WriteFile()
 }
*/
