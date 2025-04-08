package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net"
	"net/http"
	"strconv"
	"sudokusolver/db"
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

    db.Init()
    router := http.NewServeMux()
    rand := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))

    router.HandleFunc("/solve/{grid}/", func(w http.ResponseWriter, r *http.Request) {
        input := r.PathValue("grid")
        cachedSolutions := db.FindSolutions(input)
        if cachedSolutions != nil {
            fmt.Fprintf(w, "found %v (cached) solutions:\n", len(cachedSolutions))
            for _, solution := range cachedSolutions {
                fmt.Fprintf(w, "%s\n", solution)
            }
            return
        }
        solutions, err := solver.Solve(input)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        db.StoreSolutions(input, solutions)
        fmt.Fprintf(w, "found %v solutions:\n", len(solutions))
        for _, solution := range solutions {
            fmt.Fprintf(w, "%s\n", solution)
        }
    })
    router.HandleFunc("/valid/{grid}/", func(w http.ResponseWriter, r *http.Request) {
        if solver.IsValid(r.PathValue("grid")) {
            fmt.Fprintf(w, "valid")
        } else {
            fmt.Fprintf(w, "not valid")
        }
    })
    router.HandleFunc("/gen/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/gen/40", http.StatusMovedPermanently)
    })
    router.HandleFunc("/gen/{unknowns}", func(w http.ResponseWriter, r *http.Request) {
        unknowns, err := strconv.ParseUint(r.PathValue("unknowns"), 10, 8)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "generated:\n%v", solver.Generate(rand, uint8(unknowns)))
    })
    router.HandleFunc("/quit/", func(w http.ResponseWriter, r *http.Request) {
        log.Fatal("quitting")
    })
    router.HandleFunc("/db/", func(w http.ResponseWriter, r *http.Request) {
        solutions := db.AllSolutions()
        fmt.Fprintf(w, "total: %v cached inputs\n\n", len(solutions))
        for k, v := range solutions {
            fmt.Fprintf(w, "%s --> %s\n\n", k, v)
        }
    })


    server := http.Server{
        Addr: ":8080",
        Handler: LoggingMiddleware(router),
    }
    fmt.Printf("server listening on %v:%v\n", GetLocalIP(), 8080)
    log.Fatal(server.ListenAndServe())
}

