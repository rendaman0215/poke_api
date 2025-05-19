package handler

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"

	dexv1 "github.com/rendaman0215/poke_api/schema/gen/go/dex/v1"
	"github.com/rendaman0215/poke_api/services/dex/internal/domain/models"
	service "github.com/rendaman0215/poke_api/services/dex/internal/service"
)

type DexService struct{}

func (w *DexService) GetImage(ctx context.Context, req *connect.Request[dexv1.GetImageRequest]) (*connect.Response[dexv1.GetImageResponse], error) {
	slog.Info("GetImage", "req", req.Msg)
	url, err := service.GetImage(models.PokemonId{
		ID:   int(req.Msg.Id),
		Form: int(req.Msg.Form),
	})

	res := connect.NewResponse(&dexv1.GetImageResponse{
		Url: url,
	})

	if err != nil {
		slog.Error("GetImage", "err", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res.Header().Set("Weather-Version", "v1")
	return res, nil
}
