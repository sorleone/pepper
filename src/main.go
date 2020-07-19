package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/sorleone/pepper/pkg/product"
	"gitlab.com/sorleone/pepper/pkg/receipt"
)

var eol byte = 0x0A

func writeError(res http.ResponseWriter, statusCode int) {
	statusText := http.StatusText(statusCode)
	http.Error(res, statusText, statusCode)
}

func receiptHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeError(res, http.StatusBadRequest)
		return
	}

	var products []product.Product
	if err := json.Unmarshal(body, &products); err != nil {
		writeError(res, http.StatusBadRequest)
		return
	}

	total := receipt.NewStandardReceipt().Add(products...).GetTotal()
	encodedTotal, err := json.MarshalIndent(total, "", "  ")
	if err != nil {
		writeError(res, http.StatusInternalServerError)
		return
	}

	res.Write(append(encodedTotal, eol))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipt", receiptHandler).Methods("POST")

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,

		// Set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("starting server on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Blocks until a signal is received.
	<-signalChannel
	os.Stdout.Write([]byte{eol})

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
