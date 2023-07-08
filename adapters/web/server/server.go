package adapters_web_server

import (
	negroni "github.com/codegangsta/negroni"
	mux "github.com/gorilla/mux"
	handler "github.com/phelipperibeiro/golang-hexagonal-architecture/adapters/web/handler"
	application "github.com/phelipperibeiro/golang-hexagonal-architecture/application"
	log "log"
	http "net/http"
	os "os"
	time "time"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (webserver Webserver) Serve() {
	router := mux.NewRouter()
	logger := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProductHandlers(router, logger, webserver.Service)
	http.Handle("/", router)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
