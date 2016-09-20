package main

import "net"
import "fmt"
import "bufio"
import "os"
import "time"
import "strconv"

func main() {

    // connect to this socket
    var address string = "127.0.0.1:31069"
    if len(os.Args) > 1 {
        fmt.Println(os.Args)
        address = os.Args[1]
    }
    // var address string = "10.74.113.215:1069"
    // var address string = "10.74.113.217:1069"
    // fmt.Println("Connected to "+address)
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println("error when dial", err)
        os.Exit(1)
    }
    for {
        var text string = "a"
        file := os.Stdin
        fileInfo, err := file.Stat()
        // fmt.Println("input:", fileInfo)
        if err != nil {
            fmt.Println("file.Stat()", err)
            break
        }
        if fileInfo.Size() > 0 {
            reader := bufio.NewReader(os.Stdin)
            fmt.Print("Text to send: ")
            temp, err := reader.ReadString('\n')
            if err != nil {
                fmt.Println("error when reading from console", err)
            }
            text = temp
        }
        fmt.Println("text:", text)
        if text == "q" {
            conn.Close()
            break
        }
        // send to socket
        sendingTime := time.Now().Nanosecond()
        fmt.Fprintf(conn, text + "\n")
        // listen for reply
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }
        receivingTime := time.Now().Nanosecond();
        fmt.Println("Message from server: "+message+ " time(Nanosecond): "+strconv.Itoa(receivingTime-sendingTime))
        time.Sleep(time.Second*2)
    }
}