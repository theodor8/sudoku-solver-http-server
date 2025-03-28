package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"sudokusolver/solver"
	"time"
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



type wrappedWriter struct {
    http.ResponseWriter
    statusCode int
}
func (w *wrappedWriter) WriteHeader(statusCode int) {
    w.ResponseWriter.WriteHeader(statusCode)
    w.statusCode = statusCode
}
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        wrapped := &wrappedWriter{w, http.StatusOK}
        next.ServeHTTP(wrapped, r)
        log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
    })
}




func main() {

    router := http.NewServeMux()

    rand := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))

    router.HandleFunc("/solve/{grid}/", func(w http.ResponseWriter, r *http.Request) {
        solutions, err := solver.Solve(r.PathValue("grid"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "found %v solutions: %v", len(solutions), solutions)
    })
    router.HandleFunc("/valid/{grid}/", func(w http.ResponseWriter, r *http.Request) {
        if solver.IsValid(r.PathValue("grid")) {
            fmt.Fprintf(w, "valid")
        } else {
            fmt.Fprintf(w, "not valid")
        }
    })
    router.HandleFunc("/gen/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "generated: %v", solver.Generate(rand))
    })
    router.HandleFunc("/quit/", func(w http.ResponseWriter, r *http.Request) {
        log.Fatal("quitting")
    })


    server := http.Server{
        Addr: ":8080",
        Handler: LoggingMiddleware(router),
    }

    fmt.Printf("server listening on %v:%v\n", GetLocalIP(), 8080)

    log.Fatal(server.ListenAndServe())

}

