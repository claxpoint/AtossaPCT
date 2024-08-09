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

func socks5Handshake(conn net.Conn) error {
        // Method selection
        conn.Write([]byte{0x05, 0x01, 0x00}) // Version 5, 1 method, no authentication

        // Read server response
        buf := make([]byte, 2)
        _, err := io.ReadFull(conn, buf)
        if err != nil {
                return err
        }
        if buf[0] != 0x05 {
                return errors.New("invalid SOCKS5 version")
        }
        numMethods := int(buf[1])

        // Read supported methods
        methods := make([]byte, numMethods)
        _, err = io.ReadFull(conn, methods)
        if err != nil {
                return err
        }

        // Select a method (e.g., no authentication)
        conn.Write([]byte{0x05, 0x00})

        // ... (handle authentication if required)

        return nil
}

import (
        "log"
        "net"
        // ... other imports
)

func main() {
        // ...
        defer func() {
                if err != nil {
                        log.Printf("Error: %v", err)
                }
        }()
        // ...
}

// Assuming a simple username/password authentication
func handleAuth(conn net.Conn) error {
        // ... read authentication method
        if method == 0x02 { // Username/password
                // ... read username and password
                // ... authenticate
        }
        // ... handle other methods
        return nil
}

func handleConnect(conn net.Conn, req []byte) error {
        // Parse CONNECT request
        // ... extract destination address and port
        // ... establish connection to destination
        // ... send CONNECT reply
        return nil
}

go func() {
        buf := make([]byte, 4096)
        for {
                n, err := conn.Read(buf)
                if err != nil {
                    // Handle error
                    return
                }
                // Forward data to destination
            }
        }()

        // ... similar goroutine for reading from destination and writing to client

