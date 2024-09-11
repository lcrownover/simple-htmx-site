package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func htmxNavHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `
		 <nav>
		  <a href="/html/">HTML</a> |
		  <a href="/css/">CSS</a> |
		  <a href="/js/">JavaScript</a> |
		  <a href="/python/">Python</a>
		</nav> 
		`)
}

func htmxFooterHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `
		<footer>
		This is my foot
		</footer>
		`)
}

var header string = `
	<head>
		<script src="https://unpkg.com/htmx.org@2.0.2"></script>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
		<meta charset="utf-8">
	</head>
	`

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			<html>
			%s
			<div hx-get="/htmx/nav" hx-trigger="load" hx-swap="outerHTML"></div>
			<div class="container">
			content
			</div>
			<div hx-get="/htmx/footer" hx-trigger="load" hx-swap="outerHTML"></div>
			</html>
			`, header)
	})

	router.HandleFunc("/htmx/nav", htmxNavHandler)
	router.HandleFunc("/htmx/footer", htmxFooterHandler)

	address := "0.0.0.0"
	port := 8080
	connStr := fmt.Sprintf("%s:%s", address, strconv.Itoa(port))

	slog.Info("starting server", "address", address, "port", port)
	err := http.ListenAndServe(connStr, router)
	if err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
