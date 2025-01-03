package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
)

type LongURL struct {
   Original string `json:"original"`
}

type ShortURL struct {
   Short string `json:"short"`
}

func main() {
   // URL que quieres acortar
   originalURL := LongURL{Original: "https://www.example.com"}
   
   // Convertir a JSON
   jsonData, err := json.Marshal(originalURL)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   // Hacer la solicitud POST
   resp, err := http.Post("http://localhost:8080/shorten", "application/json", bytes.NewBuffer(jsonData))
   if err != nil {
      fmt.Println(err)
      return
   }
   defer resp.Body.Close()
   
   // Leer la respuesta
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   // Imprimir la URL acortada
   var shortURL ShortURL
   json.Unmarshal(body, &shortURL)
   fmt.Println("Shortened URL:", shortURL.Short)
}
