package main

import "net"
import "fmt"
import "bufio"
import "os"
import "time"

func main() {

    // connect to this socket
    var address string = "127.0.0.1:31069"
    if len(os.Args) > 1 {
        fmt.Println(os.Args)
        address = os.Args[1]
    }
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println("error when dial", err)
        os.Exit(1)
    }
    fmt.Println("Connected to " + address)

    for {
        var text string = "a"

        sendingTime := time.Now()

        // send to socket
        fmt.Fprintf(conn, text + "\n")

        // listen for reply
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("error when reading connection", err)
            break
        }
        elapsed := time.Since(sendingTime)
        fmt.Println("Message received:", message,
            ", server:", address,
            ", round trip duration: ", elapsed);
        time.Sleep(time.Second * 2)
    }
}