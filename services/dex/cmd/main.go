package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rendaman0215/poke_api/schema/gen/go/dex/v1/dexv1connect"
	"github.com/rendaman0215/poke_api/services/dex/internal/handler"
	"github.com/rendaman0215/poke_api/services/dex/internal/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Initialize the logger
	logger.InitLogger()

	port := ":8080"
	addr := fmt.Sprintf("0.0.0.0%s", port)

	ds := &handler.DexService{}
	mux := http.NewServeMux()
	path, handler := dexv1connect.NewDexImageServiceHandler(ds)
	mux.Handle(path, handler)
	server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	idleConnsClosed := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM) // SIGINT, SIGTERM を検知する
		<-c

		slog.Info("Server is shutting down...")
		if err := server.Shutdown(context.Background()); err != nil {
			slog.Error("Server Shutdown", "err", err)
			close(idleConnsClosed)
			return
		}
		slog.Info("Server is shut down")
		close(idleConnsClosed)
	}()

	slog.Info("Starting server", "addr", addr)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error("HTTP server ListenAndServe", "err", err)
	}

	<-idleConnsClosed // Shutdown処理が完了するまで待機する
}
