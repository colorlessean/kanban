package main

import(
    "fmt"
    "log"
    "net/http"
)

// Main datatypes for the api and the baord itself
// idea is that tickets are contained within columns contained within boards
// This is to be kept simple at first with complexity growing over time
type Ticket struct {
    title string
    body string
}

type Column struct {
    title string
    tickets []Ticket
}

type Board struct {
    title string
    columns []Column
}

func ticketHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/api/ticket" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported", http.StatusNotImplemented)
        return
    }

    fmt.Println("Ticket Handler")
}

func columnHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/api/column" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported", http.StatusNotImplemented)
        return
    }

    fmt.Println("Column Handler")
}

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
    fileServer := http.FileServer(http.Dir("./web"))
    http.Handle("/", fileServer)
    http.HandleFunc("/api/board", boardHandler)
    http.HandleFunc("/api/column", columnHandler)
    http.HandleFunc("/api/ticket", ticketHandler)

    fmt.Printf("Starting Server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
