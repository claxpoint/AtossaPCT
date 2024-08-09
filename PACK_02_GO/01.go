package main

import (
        "bufio"
        "fmt"
        "net"
        "strconv"
)

func main() {
        conn, err := net.Dial("tcp", "your_socks_proxy:1080")
        if err != nil {
                fmt.Println("Error connecting to SOCKS proxy:", err)
                return
        }
        defer conn.Close()

        // SOCKS5 handshake (simplified)
        conn.Write([]byte{0x05, 0x01, 0x00})
        // ... (handle authentication)

        // Establish a TCP connection
        conn.Write([]byte{0x05, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x08, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x00, 0x00, 0x01, 0x00}) // Connect to google.com
        // ... (handle response)

        // Tunnel data
        go func() {
                buf := make([]byte, 1024)
                for {
                        n, err := conn.Read(buf)
                        if err != nil {
                                fmt.Println("Error reading from SOCKS proxy:", err)
                                return
                        }
                        // Handle data from SOCKS proxy
                }
        }()

        // ... (send data to SOCKS proxy)
}
