package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	errors "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	g.Go(createHTTPServer(ctx, "hello world", ":17000", helloWorldHandler))
	g.Go(createHTTPServer(ctx, "hello name", ":18000", helloNameHandler))
	g.Go(createHTTPServer(ctx, "echo", ":19000", echoHandler))

	g.Go(func() error {
		select {
		case sig := <-signals:
			fmt.Println("\nReady to exit...")
			cancel()
			return errors.Wrapf(context.Canceled, "handle %s signal ", sig.String())
		case <-ctx.Done():
			return nil
		}
	})

	log.Println("All servers have been started successfully")
	if err := g.Wait(); errors.Is(err, context.Canceled) {
		log.Printf("%v", err)
	} else if err != nil {
		log.Fatalf("error in the server goroutines: %s\n", err)
	}
	log.Println("All servers have been closed successfully")
}

func createHTTPServer(ctx context.Context, name, addr string, handler http.HandlerFunc) func() error {
	return func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", handler)
		server := &http.Server{Addr: addr, Handler: mux}
		errChan := make(chan error, 1)

		go func() {
			<-ctx.Done()
			shutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(shutCtx); err != nil {
				errChan <- fmt.Errorf("error shutting down the %s server: %w", name, err)
			}
			log.Printf("The %s server is closed gracefully\n", name)
			close(errChan)
		}()

		log.Printf("The %s server is listening on port %s\n", name, addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			return fmt.Errorf("error starting the %s server: %w", name, err)
		}
		return <-errChan
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Hello, world!`))
}

func helloNameHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, r.Body)
}
