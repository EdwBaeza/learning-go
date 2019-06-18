package main

import (
	"io"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Hola Mundo Desde el servidor en go")
}

func main() {

	http.HandleFunc("/", PostHandler)
	http.ListenAndServe(":8000", nil)

}
