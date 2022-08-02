package delivery

import (
	"backend-core/domain"
	"context"
	"github.com/labstack/echo/v4"
)

type assetHttpHandler struct {
	e *echo.Echo
}

func (a assetHttpHandler) GetByUserId(ctx context.Context) error {
	panic("hello")
}

func NewAssetHttpHandler(e *echo.Echo) domain.AssetDelivery {
	return &assetHttpHandler{e}
}
