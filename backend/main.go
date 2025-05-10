package main

import (
	"fmt"
	"net"
	"net/http"

	"sudoku-server/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func GetLocalIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    localAddress := conn.LocalAddr().(*net.UDPAddr)
    return localAddress.IP
}

func main() {
    const port = 8081

    var r *chi.Mux = chi.NewRouter()
    handlers.Handler(r)

    fmt.Printf("server listening on %v:%v\n", GetLocalIP(), port)

    err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), r)
    if err != nil {
        log.Error(err)
    }
}

