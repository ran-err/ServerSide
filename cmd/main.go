package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go server!"))
	})
	http.ListenAndServeTLS(":8443", "cert/cert.pem", "cert/key.pem", nil)
}
