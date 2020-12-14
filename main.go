package main

import "fmt"
import "strings"
import "bufio"
import "os"

// import "rsc.io/quote"

func getNumber(id int) string {
	if id == 1 {
        fmt.Print("First number: ")
    } else {
        fmt.Print("Second number: ")
    }

    reader := bufio.NewReader(os.Stdin)
    val, _ := reader.ReadString('\n')
    // convert CRLF to LF
    val = strings.Replace(val, "\n", "", -1)
    
    return val
}

func main() {
    fmt.Println("Simple Calculator")
    fmt.Println("-----------------")
    var num1, num2 string

    for {
        num1 = getNumber(1)
        num2 = getNumber(2)
        fmt.Println(num1 + " and " + num2) 
    }
}