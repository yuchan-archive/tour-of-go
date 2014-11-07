package main

import (
"fmt"
"time"
"os"
"net"
)


func main() {
	fmt.Println("The time is ", time.Now())
	fmt.Println(os.Open("filename"))
	fmt.Println(net.Dial("tcp", "google.com"))
}

