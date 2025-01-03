package main

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"
   "sync"
   "time"
)

var (
   urlStore = make(map[string]string)
   mux      sync.RWMutex
)

type LURL struct {
   url string `json:"original"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
   var lurl LURL
   if err := json.NewDecoder(r.Body).Decode(&lurl); err != nil {
      http.Error(w, err.Error(), http.StatusBadRequest)
      return
   }
   
   short := generateShortURL(lurl.url)
   mux.Lock()
   urlStore[short] = lurl.url
   mux.Unlock()
   
   json.NewEncoder(w).Encode(map[string]string{"short": short})
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
   short := r.URL.Path[1:]
   mux.RLock()
   original, exists := urlStore[short]
   mux.RUnlock()
   
   if !exists {
      http.NotFound(w, r)
      return
   }
   
   http.Redirect(w, r, original, http.StatusFound)
}

func generateShortURL(url string) string {
   return fmt.Sprintf("%d", time.Now().UnixNano())
}

func main() {
   http.HandleFunc("/shorten", shortenHandler)
   http.HandleFunc("/", redirectHandler)
   
   log.Println("Starting server on :8080")
   log.Fatal(http.ListenAndServe(":8080", nil))
}
