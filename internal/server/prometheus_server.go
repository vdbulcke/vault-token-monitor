package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

// startPrometheusServer start the local http server
// exposing prometheus metric on /metrics path
func (v *VaultMonitorServer) startPrometheusServer() error {
	// register prometheus metrics
	PrometheusMetricsRegister()

	// starts
	promListenPort := strconv.Itoa(v.config.PrometheusListeningPort)

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	mux := http.DefaultServeMux
	mux.Handle("/metrics", promhttp.Handler())
	bindAddress := ":" + promListenPort

	// create a new server
	httpServer := http.Server{
		Addr:     bindAddress,             // configure the bind address
		Handler:  mux,                     // set the default handler
		ErrorLog: zap.NewStdLog(v.logger), // set the logger for the server
		// ReadTimeout:  5 * time.Second,                                          // max time to read request from the client
		// WriteTimeout: 10 * time.Second,                                         // max time to write response to the client
		// IdleTimeout:  120 * time.Second,                                        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		v.logger.Info("Starting Server", zap.String("port", bindAddress))

		err := httpServer.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				v.logger.Info("Server is shuting down", zap.Error(err))
				os.Exit(0)
			}
			v.logger.Error("Error starting server", zap.Error(err))
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	v.logger.Info("Got signal", zap.Any("sig", sig))

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(v.ctx, 30*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		v.logger.Error("failure while shutting down server", zap.Error(err))
		return err
	}

	return nil
}
