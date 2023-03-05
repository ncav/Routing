package main

import (
    "log"
    "net"

    "github.com/armon/go-socks5"
    "golang.org/x/net/proxy"
)

func main() {
    // Start a Tor proxy on the default port
    dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
    if err != nil {
        log.Fatal(err)
    }

    // Create a SOCKS5 server that routes traffic through the Tor proxy
    conf := &socks5.Config{
        Dial: func(ctx *socks5.DialContext) (net.Conn, error) {
            return dialer.Dial("tcp", ctx.DestAddr.String())
        },
    }
    server, err := socks5.New(conf)
    if err != nil {
        log.Fatal(err)
    }

    // Start the SOCKS5 server on port 1080
    log.Println("Starting SOCKS5 proxy on port 1080")
    if err := server.ListenAndServe("tcp", "127.0.0.1:1080"); err != nil {
        log.Fatal(err)
    }
}
