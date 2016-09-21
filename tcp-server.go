package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"
import "time"

func main() {

    fmt.Println("Launching server...")
    var port string = "31069"
    // listen on all interfaces
    ln, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Println("error when listening, exit.", err)
        os.Exit(1)
    }
    // run loop forever (or until ctrl-c)
    // accept connection on port
    fmt.Println("Listening on "+ port)
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("error when accepting, exit. ", err)
            os.Exit(1)
        }
        go handleClient(conn);
    }
}

func handleClient(conn net.Conn) {
    for {
        // will listen for message to process ending in newline (\n)
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("error when reading connection. connection:", conn, err)
            fmt.Println("Closing the connection")
            conn.Close()
            break
        }
        // output message received
        fmt.Println("Message Received:", string(message),
            ", client:", conn.RemoteAddr(),
            ", current time:", time.Now())

        // process string received
        newMessage := strings.ToUpper(message)

        // send new string back to client
        conn.Write([]byte(newMessage + "\n"))
    }
}