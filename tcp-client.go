package main

        import "net"
        import "fmt"
        import "bufio"
        import "time"
        import "strconv"

        func main() {

        // connect to this socket
        var address string = "127.0.0.1:31069"
        // var address string = "10.74.113.215:1069"
        // var address string = "10.74.113.217:1069"
        // fmt.Println("Connected to "+address)
        conn, _ := net.Dial("tcp", address)
        for {

        text := "a"
        // send to socket
        sendingTime := time.Now().Nanosecond()
        fmt.Fprintf(conn, text + "\n")
        // listen for reply
        message, _ := bufio.NewReader(conn).ReadString('\n')
        receivingTime := time.Now().Nanosecond();
        fmt.Println("Message from server: "+message+ " time(Nanosecond): "+strconv.Itoa(receivingTime-sendingTime))
        time.Sleep(time.Second*2)
        }
        }