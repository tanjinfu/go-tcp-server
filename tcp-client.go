package main

        import "net"
        import "fmt"
        import "bufio"
        import "os"

        func main() {

        // connect to this socket
var address string = "127.0.0.1:31069"
        // fmt.Println("Connected to "+address)
        conn, _ := net.Dial("tcp", address)
        for {
        // read in input from stdin
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')
        // send to socket
        fmt.Fprintf(conn, text + "\n")
        // listen for reply
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Message from server: "+message)
        }
        }