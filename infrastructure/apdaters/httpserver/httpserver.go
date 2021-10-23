package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func StartHTTPServer(host, port string) (func() error, func()) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: NewRouter(),
	}

	start := func() error {
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				return nil
			}

			return err
		}

		return nil
	}

	stop := func() {
		timeOutCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Second*5,
		)
		defer cancel()
		server.Shutdown(timeOutCtx)
	}

	return start, stop
}
