package main

import (
"encoding/json"
"net/http"
"sync"
)

type Note struct {
ID      int    `json:"id"`
Content string `json:"content"`
}

var (
notes  = []Note{}
mu     sync.Mutex
nextID = 1
)

func main() {
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]interface{}{
"status": "ok",
"notes":  len(notes),
})
})

http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
mu.Lock()
defer mu.Unlock()
if r.Method == http.MethodPost {
var n Note
if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}
n.ID = nextID
nextID++
notes = append(notes, n)
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(n)
return
}
json.NewEncoder(w).Encode(notes)
})

http.ListenAndServe(":8080", nil)
}
