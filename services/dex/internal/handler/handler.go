package handler

package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	dexv1 "github.com/rendaman0215/poke_api/schema/gen/go/dex/v1"
	"github.com/rendaman0215/poke_api/schema/gen/go/dex/v1/dexv1connect"
)

const port = ":8080"

type dexService struct{}

func (w *dexService) GetImage(ctx context.Context, req *connect.Request[dexv1.GetImageRequest]) (*connect.Response[dexv1.GetImageResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&dexv1.GetImageResponse{
		Url: "https://zukan.pokemon.co.jp/zukan-api/up/images/index/7b705082db2e24dd4ba25166dac84e0a.png",
	})
	res.Header().Set("Weather-Version", "v1")
	return res, nil
}

func serviceStart() {
	ds := &dexService{}
	mux := http.NewServeMux()
	path, handler := dexv1connect.NewDexImageServiceHandler(ds)
	mux.Handle(path, handler)

	logger.Info("Starting server", "port", port)
	http.ListenAndServe(
		fmt.Sprintf("localhost%s", port),
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
