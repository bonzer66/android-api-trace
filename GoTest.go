// testing functions before adding them to main package 


package main
  2 
  3 import (
  4         //"flag"
  5         "fmt"
  6         "os"
  7         "io/ioutil"
  8 )
  9 
 10 
 11 
 12 
 13 //func ReadFile(filename string) ([]byte, error)
 14 
 15 func main() {
 16 
 17         if len(os.Args) < 2 {
 18         fmt.Println("Missing parameter, provide file name!")
 19         return
 20     }
 21     data, err := ioutil.ReadFile(os.Args[1]) //os.Args[2]
 22     if err != nil {
 23         fmt.Println("Can't read file:", os.Args[1])
 24         panic(err)
 25     }   
 26     // data is the file content, you can use it
 27     fmt.Println("File content is:")
 28     fmt.Println(string(data))
 29  }  
 30  
 31  
 32 /*
 33 func OutputFile(outfile) {
 34         *out,d1 := []byte(out,d1)
 35         err := os.WriteFile()
 36  }

