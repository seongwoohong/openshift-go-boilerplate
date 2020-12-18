package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello Duke OKD! A new Duke T3 will stay here."
	}

	fmt.Fprintln(w, response)
	ip := GetIP(r)
	fmt.Println("Servicing request at " + time.Now().String() + " towards " + ip + " " + r.RemoteAddr + " via " + r.Method)
	for key,val := range r.Header {
		fmt.Printf(key + ": " + val)
	}
	fmt.Println(r.TLS)
}

func listenAndServe(port string) {
	fmt.Printf("serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	select {}
}
