package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type assetHttpHandler struct {
	e  *echo.Echo
	au domain.AssetUseCase
	tu domain.TransactionUseCase
}

// GetByUserId
// @Summary  Get authenticated user asset
// @Tags     asset,user
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /user/asset [get]
func (a assetHttpHandler) GetByUserId(ctx echo.Context) error {
	c := ctx.Request().Context()
	uid := ctx.Get("userID").(int)
	amount, _ := a.au.GetAmountByUserId(c, uint(uid))
	totalProfit, _ := a.tu.GetTotalProfitByUserId(c, uid)
	assetResponse := domain.AssetResponseDTO{
		Amount:      amount,
		TotalProfit: totalProfit,
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data:    assetResponse,
		Message: http.StatusText(http.StatusOK),
	})

}

func NewAssetHttpHandler(e *echo.Echo, au domain.AssetUseCase, tu domain.TransactionUseCase) domain.AssetDelivery {
	handler := &assetHttpHandler{e, au, tu}
	e.GET("/user/asset", handler.GetByUserId, common.AuthMiddleWare())
	return handler
}
