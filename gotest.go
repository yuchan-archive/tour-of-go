package main

import(
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Address string = ":9090"
)

func mapRoutes() {
	//
	// Map the homepage...
	//
	goweb.MapBefore(func(c context.Context) error {

		// add a custom header
		c.HttpResponseWriter().Header().Set("X-Custom-Header", "Goweb")

		return nil
	})

	goweb.MapAfter(func(c context.Context) error {
		// TODO: log this
		return nil
	})

	goweb.Map("/", func(c context.Context) error {
		return goweb.Respond.With(c, 200, []byte("Welcome to the Goweb example app - see the terminal for instructions."))
	})
}

func main() {
	mapRoutes()
	s := &http.Server{
		Addr:           Address,
		Handler:        goweb.DefaultHttpHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	log.Printf("%s", goweb.DefaultHttpHandler())

	go func() {
		for _ = range c {

			// sig is a ^C, handle it

			// stop the HTTP server
			log.Print("Stopping the server...")
			listener.Close()

			/*
   Tidy up and tear down
*/
			log.Print("Tearing down...")

			// TODO: tidy code up here

			log.Fatal("Finished - bye bye.  ;-)")

		}
	}()

	// begin the server
	log.Fatalf("Error in Serve: %s", s.Serve(listener))
}
