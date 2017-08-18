package main

import (
	"flag"
	"mime/multipart"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	var (
		addr      = flag.String("addr", ":8000", "Address to listen for connections")
		staticDir = flag.String("static-dir", "build", "directory to serve static assets from")
	)
	flag.Parse()

	logrus.Printf("Listening on: http://localhost%s/", *addr)
	if err := server(*addr, *staticDir).ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}

func server(addr, staticDir string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		_, header, err := r.FormFile("file")
		if err != nil {
			logrus.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sendToLaserCutter(header, r.FormValue("cut"), r.FormValue("age"))

		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("/", http.FileServer(http.Dir(staticDir)))

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}

func sendToLaserCutter(header *multipart.FileHeader, cut, age string) {
	logrus.Printf("New Order: %s on %s day old %s", header.Filename, age, cut)
}
