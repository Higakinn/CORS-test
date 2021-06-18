package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"

)

// test
func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	w.Write([]byte("backend"))
}
func main() {
    // ルーターのイニシャライズ
    r := mux.NewRouter()

    // ルート(エンドポイント)
    r.HandleFunc("/", test).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", r))
}