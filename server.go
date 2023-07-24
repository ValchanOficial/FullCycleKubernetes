package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.Handle("/", http.HandlerFunc(Hello))
	http.Handle("/healthz", http.HandlerFunc(Healthz))
	http.Handle("/configmap", http.HandlerFunc(ConfigMap))
	http.Handle("/secret", http.HandlerFunc(Secret))
	http.ListenAndServe(":8080", nil)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Printf("Uptime: %v", time.Since(startedAt))
	w.Write([]byte("ok"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello %s! You are %s years old.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/exemplo/example.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Fprintf(w, "My Family: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s - Pass: %s", user, password)
}
