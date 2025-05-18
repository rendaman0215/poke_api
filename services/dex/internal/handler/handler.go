package handler

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	dexv1 "github.com/rendaman0215/poke_api/schema/gen/go/dex/v1"
	"github.com/rendaman0215/poke_api/schema/gen/go/dex/v1/dexv1connect"
	service "github.com/rendaman0215/poke_api/services/dex/internal/service"
)

type dexService struct{}

func (w *dexService) GetImage(ctx context.Context, req *connect.Request[dexv1.GetImageRequest]) (*connect.Response[dexv1.GetImageResponse], error) {
	slog.Info("GetImage", "req", req.Msg)
	res, err := service.GetImage(req)
	if err != nil {
		slog.Error("GetImage", "err", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res.Header().Set("Weather-Version", "v1")
	return res, nil
}

func Handler() {
	port := ":8080"

	ds := &dexService{}
	mux := http.NewServeMux()
	path, handler := dexv1connect.NewDexImageServiceHandler(ds)
	mux.Handle(path, handler)
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost%s", port),
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

	slog.Info("Starting server", "port", port)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error("HTTP server ListenAndServe", "err", err)
	}

	<-idleConnsClosed // Shutdown処理が完了するまで待機する
}
