package main

import (
        //"flag"
        "fmt"
        "os"
        "io/ioutil"
)
//func ReadFile(filename string) ([]byte, error)

// takes a file as input
func main() {

        if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }
    data, err := ioutil.ReadFile(os.Args[1]) //os.Args[2]
    if err != nil {
        fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }
    // data is the file content, you can use it
    fmt.Println("File content is:")
    fmt.Println(string(data))
 }


/*
func OutputFile(outfile) {
        *out,d1 := []byte(out,d1)
        err := os.WriteFile()
 }
*/

