package main

import (
	"fmt"
	"net/http"
	"io"
	"time"
	"../domains"

)

const(
	port = ":8085"
)

func main(){
	fmt.Printf("Mock server running on port: %s ...", port)
	mock()

}

func mock() {

	http.HandleFunc("/users/", homePageUsers)
	http.HandleFunc("/countries/", homePageCountries)
	http.HandleFunc("/sites/", homePageSites)

	http.HandleFunc("/users/ping", pong)
	http.HandleFunc("/countries/ping", pong)
	http.HandleFunc("/sites/ping", pong)

	http.ListenAndServe(port, nil)
}

func homePageUsers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, domains.MockMap["users"])
}


func homePageCountries(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, domains.MockMap["countries"])
}


func homePageSites(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, domains.MockMap["sites"])
}

func pong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 500)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Pong")
}

func return500(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 200)
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Error Mockeado")
}