package main
import (
	"fmt"
	"net/http"
	"os"
)
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/",
http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Producto 1 - Go</title>
		</head>
		<body>
			<h1>Soy alumno de la UOC</h1>
			<p>Servidor: %s</p>
			<img src="/static/imagen.jpg" width="360">
		</body>
		</html>
		`, hostname)
	})
	http.ListenAndServe(":8080", nil)
}
