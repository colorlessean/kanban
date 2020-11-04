package main

import(
    "fmt"
    "log"
    "net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/api/board" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotImplemented)
        return
    }

    log.Println(w, "Board Handler")
}

func main() {
    // allows the for importing of the static assets
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/api/board", boardHandler)

    fmt.Printf("Starting Server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
