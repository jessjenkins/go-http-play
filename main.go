package main

import (
	"fmt"
	"github.com/jessjenkins/go-http-play/links"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle)
	slog.Info("starting server")
	http.ListenAndServe("localhost:9090", links.NewMiddleWare("default", mux))
}

func Handle(w http.ResponseWriter, req *http.Request) {
	slog.Info("handling request")
	ctx := req.Context()

	url1 := links.URLBuild(ctx, "url_one")
	url2 := links.URLBuild(ctx, "url_two")

	response := fmt.Sprintf("[%s] [%s]\n", url1, url2)

	w.Write([]byte(response))
}
