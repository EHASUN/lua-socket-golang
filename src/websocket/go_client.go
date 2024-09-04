
//Creating a TCP Server in go
package main

import (
    "fmt"
    "net"
)

func main() {
    // Listen for incoming connections on port 8080
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Accept incoming connections and handle them
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Handle the connection in a new goroutine
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // Close the connection when we're done
    defer conn.Close()

    // Read incoming data
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the incoming data
    fmt.Printf("Received: %s", buf)
}


//Creating a Golang socket Client

package main

import (
    "fmt"
    "net"
)

func main() {
    // Connect to the server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Send some data to the server
    _, err = conn.Write([]byte("Hello, server!"))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Close the connection
    conn.Close()
}